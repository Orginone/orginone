package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetnewitembydictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetnewitembydictLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetnewitembydictLogic {
	return GetnewitembydictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetnewitembydictLogic) Getnewitembydict(req types.DicGetnewitembydictReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
