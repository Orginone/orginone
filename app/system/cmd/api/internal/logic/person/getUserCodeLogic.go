package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserCodeLogic {
	return GetUserCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCodeLogic) GetUserCode(req types.GetUserCodeReq) (resp *types.CommonResponse, err error) {
	rpcRes, err := l.svcCtx.SystemRpc.GetNewPersonUserCode(l.ctx, &system.TenantCodeReq{TenantCode: req.TenantCode})
	return types.JsonResult(rpcRes.Json, err)
}
