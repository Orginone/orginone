package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketusedapp"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitMarketUsedAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitMarketUsedAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitMarketUsedAppLogic {
	return &SubmitMarketUsedAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增或修改
func (l *SubmitMarketUsedAppLogic) SubmitMarketUsedApp(in *market.SubmitMarketUsedAppReq) (*market.CommonRpcRes, error) {
	useAppEntitys, err := l.svcCtx.MarketStore.AsMarketUsedApp.Query().Where(asmarketusedapp.UserID(in.UserId),
		asmarketusedapp.AppIDNEQ(in.AppId), asmarketusedapp.IDNEQ(in.Id)).Order(schema.Asc(asmarketusedapp.FieldSort)).All(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	for index, use := range useAppEntitys {
		_, err = use.Update().SetSort(int64(index + 1)).Save(l.ctx)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
	}
	in.Sort = int64(len(useAppEntitys) + 1)
	count, err := l.svcCtx.MarketStore.AsMarketUsedApp.Update().Where(asmarketusedapp.IsDeleted(0), asmarketusedapp.Or(
		asmarketusedapp.ID(in.Id),
		asmarketusedapp.And(asmarketusedapp.UserID(in.UserId), asmarketusedapp.AppID(in.AppId)))).SetSort(in.Sort).Save(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if count < 1 {
		_, err = l.svcCtx.MarketStore.AsMarketUsedApp.Create().SetAppID(in.AppId).SetUserID(in.UserId).SetSort(in.Sort).Save(l.ctx)
	}
	return market.Result("true", err)
}
