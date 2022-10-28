package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectuserIdByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectuserIdByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SelectuserIdByRoleIdLogic {
	return SelectuserIdByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectuserIdByRoleIdLogic) SelectuserIdByRoleId(req types.SelectuserIdByRoleIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
