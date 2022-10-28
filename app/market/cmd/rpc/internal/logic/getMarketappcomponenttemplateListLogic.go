package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketappcomponenttemplate"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMarketappcomponenttemplateListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMarketappcomponenttemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMarketappcomponenttemplateListLogic {
	return &GetMarketappcomponenttemplateListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询Portal模板
func (l *GetMarketappcomponenttemplateListLogic) GetMarketappcomponenttemplateList(in *market.MarketappcomponenttemplateListReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppComponentTemplate.Query().Where(asmarketappcomponenttemplate.IsDeleted(0),
		asmarketappcomponenttemplate.NameContains(in.Name), asmarketappcomponenttemplate.Status(in.Status))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	componmentEntitys, err := query.Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return market.PageResult(in.Page, total, componmentEntitys, err)
}
