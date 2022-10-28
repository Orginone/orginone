package groupappdistributiondata

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/models"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetConfigLogic {
	return GetConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigLogic) GetConfig(req types.GetConfigReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.MarketRpc.GetGroupDistributeConfigList(l.ctx, &market.GroupAppByIdReq{
		GroupId: req.GroupId,
		AppId:   req.AppId,
	})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	configModels := make([]*models.DistributionConfigModel, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &configModels)
	return types.JsonResult(configModels, err)
}
