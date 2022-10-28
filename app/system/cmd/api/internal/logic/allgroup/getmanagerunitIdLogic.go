package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetmanagerunitIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetmanagerunitIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetmanagerunitIdLogic {
	return GetmanagerunitIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetmanagerunitIdLogic) GetmanagerunitId(req types.GetManagerUnitIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
