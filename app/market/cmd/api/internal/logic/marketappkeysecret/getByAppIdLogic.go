package marketappkeysecret

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByAppIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByAppIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetByAppIdLogic {
	return GetByAppIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByAppIdLogic) GetByAppId(req types.GetByAppIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
