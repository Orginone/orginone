package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetdictallversionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetdictallversionLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetdictallversionLogic {
	return GetdictallversionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetdictallversionLogic) Getdictallversion(req types.DicGetdictallversionReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
