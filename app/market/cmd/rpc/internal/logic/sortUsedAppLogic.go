package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketusedapp"

	"github.com/zeromicro/go-zero/core/logx"
)

type SortUsedAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSortUsedAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SortUsedAppLogic {
	return &SortUsedAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 常用应用排序
func (l *SortUsedAppLogic) SortUsedApp(in *market.SortUsedAppReq) (*market.CommonRpcRes, error) {
	for sort, appId := range in.AppIdList {
		_, err := l.svcCtx.MarketStore.AsMarketUsedApp.Update().Where(asmarketusedapp.IsDeleted(0),
			asmarketusedapp.AppID(appId), asmarketusedapp.UserID(in.UserId)).SetSort(int64(sort + 1)).Save(l.ctx)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
	}
	return market.Result("true", nil)
}
