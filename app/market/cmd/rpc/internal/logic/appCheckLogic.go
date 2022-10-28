package logic

import (
	"context"
	"strings"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketappcomponent"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppCheckLogic {
	return &AppCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过或拒绝应用 status 0.注册审核中,1.注册通过,2.注册拒绝,3.部署审核中,4.部署通过,5.部署拒绝,6.发布审核中,7.发布通过,8.发布拒绝
func (l *AppCheckLogic) AppCheck(in *market.AppCheckReq) (*market.CommonRpcRes, error) {
	var subStatus int64 = 1
	if in.Status {
		subStatus = 0
	}
	var err error = nil
	for _, appId := range in.AppIds {
		appEntity, err := l.svcCtx.MarketStore.AsMarketApp.FindOne(l.ctx, appId)
		if err != nil {
			continue
		}
		newStatus := subStatus
		switch appEntity.Status {
		case 0:
			newStatus += 1
			if in.Status {
				err = l.svcCtx.MarketStore.AsMarketAppKeySecret.Create().SetAppID(appId).
					SetAppSecret(strings.ToUpper(tools.GenAppKeyOrSecret(128))).
					SetAppKey(tools.GenAppKeyOrSecret(50)).Exec(l.ctx)
			}
		case 3:
			newStatus += 4
		case 6:
			newStatus += 7
			if in.Status {
				err = l.svcCtx.MarketStore.AsMarketAppComponent.Update().Where(asmarketappcomponent.IsDeleted(0),
					asmarketappcomponent.AppID(appId)).SetStatus(-1).Exec(l.ctx)
			}
		default:
			continue
		}
		err = appEntity.Update().SetStatus(newStatus).Exec(l.ctx)
	}
	return market.Result("true", err)
}
