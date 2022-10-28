package marketappapply

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecretByKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSecretByKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSecretByKeyLogic {
	return GetSecretByKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSecretByKeyLogic) GetSecretByKey(req types.GetSecretByKeyReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
