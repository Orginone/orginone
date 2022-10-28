package logic

import (
	"context"
	"strconv"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketroledistribution"
	"orginone/common/tools/linq"

	linqx "github.com/ahmetb/go-linq/v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppRoleDistribeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppRoleDistribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppRoleDistribeLogic {
	return &GetAppRoleDistribeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用权限分配数据
func (l *GetAppRoleDistribeLogic) GetAppRoleDistribe(in *market.GetAppRoleDistribeReq) (*market.CommonRpcRes, error) {
	appRoleDistribeEntitys, err := l.svcCtx.MarketStore.AsMarketAppRole.Query().Where(asmarketapprole.AppID(in.AppId), asmarketapprole.IsDeleted(0),
		asmarketapprole.HasRoleDistribsWith(asmarketroledistribution.IsDeleted(0))).QueryRoleDistribs().
		Where(asmarketroledistribution.IsDeleted(0), asmarketroledistribution.TenantCode(in.TenantCode)).WithAgencyx().WithJobx().WithUserx().All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	data := make([]*model.AppRoleDistribeRes, 0)
	linq.From(appRoleDistribeEntitys).GroupByT(
		func(e *schema.AsMarketRoleDistribution) int64 {
			return e.RoleID
		}, func(e interface{}) interface{} {
			return e
		}).SelectT(func(i linqx.Group) *model.AppRoleDistribeRes {
		item := new(model.AppRoleDistribeRes)
		item.Agencys = make([]map[string]interface{}, 0)
		item.Users = make([]map[string]interface{}, 0)
		item.Jobs = make([]map[string]interface{}, 0)
		item.RoleId = i.Key.(int64)
		for _, v := range i.Group {
			entity := v.(*schema.AsMarketRoleDistribution)
			if entity.AgencyID > 0 && entity.Edges.Agencyx != nil {
				item.Agencys = append(item.Agencys, map[string]interface{}{
					"agencyId":   strconv.FormatInt(entity.AgencyID, 10),
					"agencyName": entity.Edges.Agencyx.AgencyName,
				})
			}
			if entity.UserID > 0 && entity.Edges.Userx != nil {
				item.Users = append(item.Users, map[string]interface{}{
					"userId":   strconv.FormatInt(entity.UserID, 10),
					"realName": entity.Edges.Userx.UserName,
				})
			}
			if entity.JobID > 0 && entity.Edges.Jobx != nil {
				item.Jobs = append(item.Jobs, map[string]interface{}{
					"jobId":   strconv.FormatInt(entity.JobID, 10),
					"jobName": entity.Edges.Jobx.JobName,
				})
			}
		}
		return item
	}).ToSlice(&data)
	return market.JsonResult(data, err)
}
