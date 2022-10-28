package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetgroupallunitlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetgroupallunitlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetgroupallunitlistLogic {
	return GetgroupallunitlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetgroupallunitlistLogic) Getgroupallunitlist(req types.GetGroupAllUnitListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
