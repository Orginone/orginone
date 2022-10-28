package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateacttodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateacttodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateacttodoLogic {
	return UpdateacttodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateacttodoLogic) Updateacttodo(req types.UpdateacttodoReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
