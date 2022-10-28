package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetdeptjobroleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetdeptjobroleLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetdeptjobroleLogic {
	return GetdeptjobroleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetdeptjobroleLogic) Getdeptjobrole(req types.GetDeptJobRoleReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
