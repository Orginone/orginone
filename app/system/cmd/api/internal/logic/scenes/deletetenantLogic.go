package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletetenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletetenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeletetenantLogic {
	return DeletetenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletetenantLogic) Deletetenant(req types.DeleteTenantReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
