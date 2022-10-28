package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetmanagerunitidlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetmanagerunitidlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetmanagerunitidlistLogic {
	return GetmanagerunitidlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetmanagerunitidlistLogic) Getmanagerunitidlist(req types.GetManagerUnitIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
