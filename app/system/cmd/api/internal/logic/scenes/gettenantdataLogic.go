package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GettenantdataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGettenantdataLogic(ctx context.Context, svcCtx *svc.ServiceContext) GettenantdataLogic {
	return GettenantdataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GettenantdataLogic) Gettenantdata(req types.GetTenantDataReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
