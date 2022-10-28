package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModeldownunittenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModeldownunittenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) ModeldownunittenantLogic {
	return ModeldownunittenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModeldownunittenantLogic) Modeldownunittenant(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
