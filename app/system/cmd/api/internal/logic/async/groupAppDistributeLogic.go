package async

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAppDistributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupAppDistributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GroupAppDistributeLogic {
	return GroupAppDistributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupAppDistributeLogic) GroupAppDistribute(req types.GroupAppDistributeReq) (resp *types.CommonResponse, err error) {
	userId, _, _, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	_, err = l.svcCtx.SystemRpc.GroupAppDistribute(l.ctx, &system.GroupAppDistributeReq{
		UserId:       userId,
		AppId:        req.AppId,
		GroupId:      req.GroupId,
		DistribeJson: string(req.JsonBody),
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(tools.GenTenantCode())
}
