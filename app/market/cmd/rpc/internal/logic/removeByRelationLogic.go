package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketusedapp"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveByRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveByRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveByRelationLogic {
	return &RemoveByRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除
func (l *RemoveByRelationLogic) RemoveByRelation(in *market.RemoveByRelationReq) (*market.CommonRpcRes, error) {
	_, err := l.svcCtx.MarketStore.AsMarketUsedApp.Delete().Where(asmarketusedapp.IsDeleted(0),
		asmarketusedapp.AppID(in.AppId), asmarketusedapp.UserID(in.UserId)).Exec(l.ctx)
	return market.Result("true", err)
}
