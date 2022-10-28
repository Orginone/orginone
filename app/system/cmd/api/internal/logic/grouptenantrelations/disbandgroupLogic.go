package grouptenantrelations

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisbandgroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDisbandgroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) DisbandgroupLogic {
	return DisbandgroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisbandgroupLogic) Disbandgroup(req types.DisbandGroupReq) (resp *types.CommonResponse, err error) {
	userId, _, _, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.CommonResult(l.svcCtx.SystemRpc.DisbandGroupById(l.ctx, &system.DisbandGroupReq{
		UserId:        userId,
		GroupId:       req.GroupId,
		SourceGroupId: req.SourceGroupId,
	}))
}
