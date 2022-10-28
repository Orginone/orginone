package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModeldownreplenishphoneNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModeldownreplenishphoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) ModeldownreplenishphoneNumberLogic {
	return ModeldownreplenishphoneNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModeldownreplenishphoneNumberLogic) ModeldownreplenishphoneNumber(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
