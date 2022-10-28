package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketappcomponenttemplate"
	"orginone/common/schema/asmarketappusertemplate"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTemplatesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTemplatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTemplatesLogic {
	return &GetUserTemplatesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户应用组件
func (l *GetUserTemplatesLogic) GetUserTemplates(in *market.GetUserTemplateReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppUserTemplate.Query().Where(asmarketappusertemplate.IsDeleted(0), asmarketappusertemplate.UserID(in.UserId)).QueryTemplatex().
		Where(asmarketappcomponenttemplate.IsDeleted(0), asmarketappcomponenttemplate.Status(1))
	if !tools.IsNilOrEmpty(in.Page.Filter) {
		query = query.Where(asmarketappcomponenttemplate.NameContains(in.Page.Filter))
	}
	total, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if total == 0 {
		query = l.svcCtx.MarketStore.AsMarketAppComponentTemplate.Query().Where(asmarketappcomponenttemplate.IsDeleted(0), asmarketappcomponenttemplate.Status(-1))
		total, err = query.Count(l.ctx)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
	}
	if in.Page.Limit > 100 {
		in.Page.Limit = 100
	}
	actEntitys, err := query.Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return market.PageResult(in.Page, total, actEntitys, err)
}
