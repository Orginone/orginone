package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetbelowunitlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetbelowunitlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetbelowunitlistLogic {
	return GetbelowunitlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetbelowunitlistLogic) Getbelowunitlist(req types.GetBelowUnitListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
