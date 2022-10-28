package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallpersonsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallpersonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallpersonsLogic {
	return GetallpersonsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallpersonsLogic) Getallpersons(req types.GetAllPersonsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
