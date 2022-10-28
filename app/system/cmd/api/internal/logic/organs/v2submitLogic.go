package organs

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2submitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2submitLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2submitLogic {
	return V2submitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2submitLogic) V2submit(req *schema.AsUnit) (resp *types.CommonResponse, err error) {
	bytes, err := json.Marshal(req)
	return types.JsonResult(l.svcCtx.SystemRpc.UpdateOrgansInfo(l.ctx, &system.CommonRpcRes{Json: string(bytes)}))
}
