package appjob

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllAppJobPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllAppJobPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllAppJobPageLogic {
	return GetAllAppJobPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllAppJobPageLogic) GetAllAppJobPage(req types.GetAllAppJobPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
