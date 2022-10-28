package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetactivedaymonthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetactivedaymonthLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetactivedaymonthLogic {
	return GetactivedaymonthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetactivedaymonthLogic) Getactivedaymonth(req types.GetActiveDayMonthReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
