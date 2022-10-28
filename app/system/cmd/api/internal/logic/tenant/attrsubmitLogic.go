package tenant

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttrsubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttrsubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) AttrsubmitLogic {
	return AttrsubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttrsubmitLogic) Attrsubmit(req []byte) (resp *types.CommonResponse, err error) {
	return types.JsonResult(l.svcCtx.SystemRpc.TenantAttrSubmit(l.ctx, &system.CommonRpcRes{Json: string(req)}))
}
