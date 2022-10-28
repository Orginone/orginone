package logic

import (
	"context"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketappnotice"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMarkAppNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMarkAppNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMarkAppNoticeListLogic {
	return &FindMarkAppNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用详情
func (l *FindMarkAppNoticeListLogic) FindMarkAppNoticeList(in *market.FindMarkAppNoticeListReq) (*market.CommonRpcRes, error) {
	query := l.svcCtx.MarketStore.AsMarketAppNotice.Query().Where(asmarketappnotice.NoticeTitleContains(in.NoticeTitle),
		asmarketappnotice.GroupOrUnit(in.GroupOrUnit), asmarketappnotice.IsDeleted(0))
	count, err := query.Count(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	data, err := query.Order(schema.Asc(asmarketappnotice.FieldCreateTime)).
		Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return market.PageResult(in.Page, int64(count), data, err)
}
