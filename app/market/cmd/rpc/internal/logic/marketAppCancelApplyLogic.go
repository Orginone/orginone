package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketapp"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketAppCancelApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarketAppCancelApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketAppCancelApplyLogic {
	return &MarketAppCancelApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用取消部署
func (l *MarketAppCancelApplyLogic) MarketAppCancelApply(in *market.AppCancelApplyReq) (*market.CommonRpcRes, error) {
	//开启事务
	tx, err := l.svcCtx.MarketStore.Tx(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketApp.Update().Where(asmarketapp.IsDeleted(0), asmarketapp.IDIn(in.AppIds...),
		asmarketapp.Status(0)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketApp.Update().Where(asmarketapp.IsDeleted(0), asmarketapp.IDIn(in.AppIds...),
		asmarketapp.Status(3)).SetStatus(1).Save(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketApp.Update().Where(asmarketapp.IsDeleted(0), asmarketapp.IDIn(in.AppIds...),
		asmarketapp.Status(6)).SetStatus(4).Save(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	err = tx.Commit()
	return market.Result("true", err)
}
