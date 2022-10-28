package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallpersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallpersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallpersonLogic {
	return GetallpersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallpersonLogic) Getallperson(req types.GetAllPersonReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
