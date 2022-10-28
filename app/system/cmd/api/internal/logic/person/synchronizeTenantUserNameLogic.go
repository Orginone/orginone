package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SynchronizeTenantUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSynchronizeTenantUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) SynchronizeTenantUserNameLogic {
	return SynchronizeTenantUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SynchronizeTenantUserNameLogic) SynchronizeTenantUserName(req types.SynchronizeTenantUserNameReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
