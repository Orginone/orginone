package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketappcomponent"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppComponmentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppComponmentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppComponmentsLogic {
	return &GetAppComponmentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用组件
func (l *GetAppComponmentsLogic) GetAppComponments(in *market.GetComponmentReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppComponent.Query().Where(asmarketappcomponent.IsDeleted(0))
	if !tools.IsNilOrEmpty(in.Page.Filter) {
		query = query.Where(asmarketappcomponent.NameContains(in.Page.Filter))
	}
	switch in.Flag {
	case 1:
		query = query.Where(asmarketappcomponent.AppIDIsNil())
		break
	case 2:
		query = query.Where(asmarketappcomponent.AppIDNotNil())
		break
	}
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if in.Page.Limit > 100 {
		in.Page.Limit = 100
	}
	componments, err := query.Order(schema.Asc(asmarketappcomponent.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return market.PageResult(in.Page, total, componments, err)
}
