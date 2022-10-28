package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupDefaultJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupDefaultJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateGroupDefaultJobLogic {
	return CreateGroupDefaultJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupDefaultJobLogic) CreateGroupDefaultJob(req types.CreateGroupDefaultJobReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
