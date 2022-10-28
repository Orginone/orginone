package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletepersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletepersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeletepersonLogic {
	return DeletepersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletepersonLogic) Deleteperson(req types.DeletePersonReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
