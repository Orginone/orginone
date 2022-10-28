package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GettenantintopersonsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGettenantintopersonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GettenantintopersonsLogic {
	return GettenantintopersonsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GettenantintopersonsLogic) Gettenantintopersons(req types.GetTenantIntoPersonsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
