package logic

import (
	"context"
	"strings"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/asjob"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTenantPersonsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTenantPersonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTenantPersonsLogic {
	return &FindTenantPersonsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取租户下人员信息
func (l *FindTenantPersonsLogic) FindTenantPersons(in *system.SearchPersonByTenantCodeReq) (*system.CommonRpcRes, error) {
	personQuery := l.svcCtx.SystemStore.AsPerson.Query().Where(asperson.IsDeleted(0), asperson.TenantCode(in.TenantCode)).
		Where(asperson.HasUserxWith(asuser.IsDeleted(0)), asperson.RealNameContains(in.Page.Filter))
	//isActivate 0-未激活, 1-激活, 3-全部
	switch in.IsActivate {
	case 3:
		personQuery = personQuery.Where(asperson.StatusNEQ(1))
		break
	case 1, 2:
		personQuery = personQuery.Where(asperson.Status(in.IsActivate + 1)).
			Where(asperson.HasUserxWith(asuser.Status(in.IsActivate + 1)))
		break
	}
	if in.DeptId > 0 {
		if in.HasChild {
			innerAgencys, err := l.svcCtx.SystemStore.AsInnerAgency.Query().Where(asinneragency.TenantCode(in.TenantCode), asinneragency.IsDeleted(0)).All(l.ctx)
			if err != nil {
				return &system.CommonRpcRes{}, err
			}
			ids := GetChildInagencyIds(innerAgencys, in.DeptId, "")
			ids = append(ids, in.DeptId)
			personQuery = personQuery.Where(asperson.HasAgencysWith(asinneragency.IDIn(ids...)))
		} else {
			personQuery = personQuery.Where(asperson.HasAgencysWith(asinneragency.ID(in.DeptId)))
		}
	}
	if in.DeptId == -1 {
		personQuery = personQuery.Where(asperson.Not(asperson.HasAgencysWith()))
	}
	if in.JobId > 0 {
		personQuery = personQuery.Where(asperson.HasJobsWith(asjob.ID(in.JobId)))
	}
	if in.JobId == -1 {
		personQuery = personQuery.Where(asperson.Not(asperson.HasJobs()))
	}
	//排序flag  0-代表升序，1-代表降序
	if in.Flag == 0 {
		personQuery.Order(schema.Asc(asperson.FieldUserCode))
	} else {
		personQuery.Order(schema.Desc(asperson.FieldUserCode))
	}
	total, err := personQuery.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	persons := make([]*model.PersonInfo, 0)
	if in.Page.Offset < 0 {
		in.Page.Offset = 0
	}
	if in.Page.Limit < 1 {
		in.Page.Limit = 100000
	}
	userIds, err := personQuery.Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).Select(asperson.FieldUserID).Int64s(l.ctx)
	if err == nil && len(userIds) > 0 {
		persons, err = l.svcCtx.SystemStore.QueryPersonInfo(l.ctx, userIds...)
	}
	return system.PageResult(in.Page, total, persons, err)
}

func GetChildInagencyIds(array []*schema.AsInnerAgency, parentID int64, filter string) []int64 {
	datas := make([]int64, 0)
	for _, item := range array {
		if item.ParentID == parentID {
			if strings.Contains(item.AgencyName, filter) {
				datas = append(datas, item.ID)
			}
			datas = append(datas, GetChildInagencyIds(array, item.ID, filter)...)
		}
	}
	return datas
}
