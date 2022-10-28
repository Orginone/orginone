package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetmanagergroupidlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetmanagergroupidlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetmanagergroupidlistLogic {
	return GetmanagergroupidlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetmanagergroupidlistLogic) Getmanagergroupidlist(req types.GetManagerGroupIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
