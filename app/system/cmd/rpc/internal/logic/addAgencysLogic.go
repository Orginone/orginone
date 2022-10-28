package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asinneragency"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAgencysLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAgencysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAgencysLogic {
	return &AddAgencysLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 部门导入
func (l *AddAgencysLogic) AddAgencys(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	l.ctx = context.Background()
	agencys := make([]*model.AgencyData, 0)
	err := json.Unmarshal([]byte(in.Json), &agencys)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	parentNames := make([]string, 0)
	agencyNames := make([]string, 0)
	for _, v := range agencys {
		parentNames = append(parentNames, v.ParentAgencyName)
		agencyNames = append(agencyNames, v.AgencyName)
	}
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	parentData, err := tx.AsInnerAgency.Query().
		Where(
			asinneragency.IsDeleted(0),
			asinneragency.AgencyNameIn(parentNames...),
			asinneragency.TenantCode(agencys[0].TenantCode)).
		All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if len(parentData) < 0 {
		return &system.CommonRpcRes{}, err
	}
	existAgencyNames, err := tx.AsInnerAgency.Query().
		Where(
			asinneragency.AgencyNameIn(agencyNames...),
			asinneragency.TenantCode(agencys[0].TenantCode),
			asinneragency.IsDeleted(0),
		).Select(asinneragency.FieldAgencyName).Strings(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	count, err := tx.AsInnerAgency.Query().
		Where(
			asinneragency.TenantCode(agencys[0].TenantCode),
			asinneragency.IsDeleted(0)).Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	prefixCode := parentData[0].AgencyCode[0:3]
	createAgencyCreates := make([]*schema.AsInnerAgencyCreate, 0)
	for _, v := range agencys {
		if tools.ArrIndexOf(existAgencyNames, v.AgencyName) < 0 {
			for _, data := range parentData {
				if data.AgencyName == v.ParentAgencyName {
					count++
					createAgencyCreates = append(
						createAgencyCreates,
						l.svcCtx.SystemStore.AsInnerAgency.Create().
							SetTenantCode(v.TenantCode).
							SetAgencyCode(fmt.Sprintf("%s%04d", prefixCode, count)).
							SetParentID(data.ID).
							SetAgencyName(v.AgencyName))
					break
				}
			}
		}
	}
	_, err = tx.AsInnerAgency.CreateBulk(createAgencyCreates...).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &system.CommonRpcRes{}, err
	}
	return system.JsonResult(true, tx.Commit())
}
