package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListroleIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListroleIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListroleIdListLogic {
	return ListroleIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListroleIdListLogic) ListroleIdList(req types.ListRoleIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
