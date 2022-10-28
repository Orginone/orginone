package redeploy

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsAddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) IsAddRoleLogic {
	return IsAddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsAddRoleLogic) IsAddRole(req types.IsAddRoleReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
