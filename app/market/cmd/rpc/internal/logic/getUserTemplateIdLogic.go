package logic

import (
	"context"
	"fmt"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketappusertemplate"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTemplateIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTemplateIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTemplateIdLogic {
	return &GetUserTemplateIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户应用组件ID
func (l *GetUserTemplateIdLogic) GetUserTemplateId(in *market.GetUserTemplateIdReq) (*market.CommonRpcRes, error) {
	templateEntitys, err := l.svcCtx.MarketStore.AsMarketAppUserTemplate.Query().Where(asmarketappusertemplate.IsDeleted(0), asmarketappusertemplate.UserID(in.UserId)).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if len(templateEntitys) < 1 {
		return market.Result("", err)
	}
	return market.Result(fmt.Sprintf("%d", templateEntitys[0].ID), err)
}
