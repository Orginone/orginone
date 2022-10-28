package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListroleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListroleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListroleListLogic {
	return ListroleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListroleListLogic) ListroleList(req types.ListRoleListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
