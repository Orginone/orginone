package marketappnotice

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MyNoticeListLogic {
	return MyNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyNoticeListLogic) MyNoticeList(req types.MyNoticeListReq) (resp *types.CommonResponse, err error) {
	return types.JsonResult(models.PageResponse{
		Current:     int64(req.Current),
		Size:        int64(req.Size),
		Records:     make([]interface{}, 0),
		Total:       0,
		SearchCount: false,
		Pages:       1,
	}, nil)
}
