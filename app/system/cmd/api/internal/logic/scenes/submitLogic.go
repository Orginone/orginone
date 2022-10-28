package scenes

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SubmitLogic {
	return SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req *schema.AsTenant) (resp *types.CommonResponse, err error) {
	bytes, err := json.Marshal(req)
	return types.JsonResult(l.svcCtx.SystemRpc.UpdateScenesInfo(l.ctx, &system.CommonRpcRes{Json: string(bytes)}))
}
