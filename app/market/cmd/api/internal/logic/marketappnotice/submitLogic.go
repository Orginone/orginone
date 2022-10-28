package marketappnotice

import (
	"context"
	"net/http"
	"strconv"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitLogic {
	return SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req types.SubmitReq7) (resp *types.CommonResponse, err error) {
	id := -1
	if !tools.IsNilOrEmpty(req.Id) {
		id, err = strconv.Atoi(req.Id)
		if err != nil {
			return types.Failed(http.StatusBadRequest, err)
		}
	}
	return types.JsonResult(l.svcCtx.MarketRpc.UpdateMarketAppNotice(l.ctx, &market.UpdateMarketAppNoticeReq{
		Id:                  int64(id),
		NoticeTitle:         req.NoticeTitle,
		NoticeUnitIds:       req.NoticeUnitIds,
		NoticeRoleIds:       req.NoticeRoleIds,
		NoticeContent:       req.NoticeContent,
		GroupOrUnit:         req.GroupOrUnit,
		NoticeReleaseStatus: req.NoticeReleaseStatus,
		NoticeReleaseUnitId: req.NoticeReleaseUnitId,
	}))
}
