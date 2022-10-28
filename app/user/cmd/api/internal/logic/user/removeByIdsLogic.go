package user

import (
	"context"
	"strings"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common/rpcs/user"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveByIdsLogic {
	return RemoveByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveByIdsLogic) RemoveByIds(req types.RemoveByIdsReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.UserRpc.RemoveUserByIds(l.ctx, &user.RemoveUserByIdsReq{
		Ids: tools.ArrayStrToInt64(strings.Split(req.Ids, ",")),
	}))
}
