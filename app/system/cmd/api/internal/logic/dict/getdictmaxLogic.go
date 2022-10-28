package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetdictmaxLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetdictmaxLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetdictmaxLogic {
	return GetdictmaxLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetdictmaxLogic) Getdictmax(req types.DicGetdictmaxReq) (resp *types.CommonResponse, err error) {
	return types.Successful(1)
}
