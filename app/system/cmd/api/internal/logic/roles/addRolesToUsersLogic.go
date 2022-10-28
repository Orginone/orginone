package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRolesToUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRolesToUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddRolesToUsersLogic {
	return AddRolesToUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRolesToUsersLogic) AddRolesToUsers(req types.AddRolesToUsersReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
