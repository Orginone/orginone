package user

import (
	"context"
	"encoding/json"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common/rpcs/user"
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

func (l *SubmitLogic) Submit(req *schema.AsUser) (resp *types.CommonResponse, err error) {
	bytes, err := json.Marshal(req)
	return types.JsonResult(l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &user.CommonRpcRes{Json: string(bytes)}))
}
