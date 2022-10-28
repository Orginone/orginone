package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpGroupIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUpGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUpGroupIdLogic {
	return GetUpGroupIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUpGroupIdLogic) GetUpGroupId(req types.GetUpGroupIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
