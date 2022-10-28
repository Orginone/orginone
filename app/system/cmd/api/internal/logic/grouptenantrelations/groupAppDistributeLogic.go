package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

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

func (l *GroupAppDistributeLogic) GroupAppDistribute(req types.GroupAppDistributeReq1) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
