package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActversionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActversionLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActversionLogic {
	return ActversionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActversionLogic) Actversion(req types.DicActversionReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
