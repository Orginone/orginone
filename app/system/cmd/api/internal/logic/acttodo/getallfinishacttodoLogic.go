package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallfinishacttodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallfinishacttodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallfinishacttodoLogic {
	return GetallfinishacttodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallfinishacttodoLogic) Getallfinishacttodo(req types.GetallfinishacttodoReq) (resp *types.CommonResponse, err error) {
	return types.JsonResult(models.PageResponse{
		Current:     int64(req.Current),
		Size:        int64(req.Size),
		Records:     make([]interface{}, 0),
		Total:       0,
		SearchCount: false,
		Pages:       1,
	}, nil)
}
