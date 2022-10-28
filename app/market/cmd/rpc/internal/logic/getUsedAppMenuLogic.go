package logic

import (
	"context"
	"orginone/common/schema"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/asmarketusedapp"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsedAppMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsedAppMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsedAppMenuLogic {
	return &GetUsedAppMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取常用应用菜单
func (l *GetUsedAppMenuLogic) GetUsedAppMenu(in *market.GetUsedAppMenuReq) (*market.CommonRpcRes, error) {
	appIds, err := l.svcCtx.MarketStore.AsMarketUsedApp.Query().Where(asmarketusedapp.IsDeleted(0), asmarketusedapp.UserID(in.UserId)).
		Order(schema.Asc(asmarketusedapp.FieldSort)).Select(asmarketusedapp.FieldAppID).Int64s(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	appMenuEntitys, err := l.svcCtx.MarketStore.AsMarketMenu.Query().Where(asmarketmenu.IsDeleted(0), asmarketmenu.AppIDIn(appIds...),
		asmarketmenu.Or(asmarketmenu.ParentIDIsNil(), asmarketmenu.ParentID(0))).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	appMenus := make([]*schema.AsMarketMenu, 0)
	for _, appId := range appIds {
		for _, menu := range appMenuEntitys {
			if appId == menu.AppID {
				appMenus = append(appMenus, menu)
			}
		}
	}
	return market.JsonResult(appMenus, err)
}
