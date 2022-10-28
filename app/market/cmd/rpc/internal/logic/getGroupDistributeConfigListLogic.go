package logic

import (
	"context"
	"encoding/json"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/models"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asappgroupdistributiondata"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupDistributeConfigListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupDistributeConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupDistributeConfigListLogic {
	return &GetGroupDistributeConfigListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取集团分发配置
func (l *GetGroupDistributeConfigListLogic) GetGroupDistributeConfigList(in *market.GroupAppByIdReq) (*market.CommonRpcRes, error) {
	configModels := make([]*models.DistributionConfigModel, 0)
	configJson, err := l.svcCtx.MarketStore.AsAppGroupDistributionData.Query().Where(asappgroupdistributiondata.AppID(in.AppId),
		asappgroupdistributiondata.GroupID(in.GroupId), asappgroupdistributiondata.IsDeleted(0)).Limit(1).
		Select(asappgroupdistributiondata.FieldConfig).Strings(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if len(configJson) < 1 {
		for _, v := range []int64{0, 1} {
			configModels = append(configModels, &models.DistributionConfigModel{
				DistributeType: "",
				Contain:        v,
				ConfigList:     make([]*models.DistributionConfigItem, 0),
			})
		}
	} else {
		err = json.Unmarshal([]byte(configJson[0]), &configModels)
	}
	return market.JsonResult(configModels, err)
}
