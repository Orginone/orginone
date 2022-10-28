package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeldictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeldictLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeldictLogic {
	return DeldictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeldictLogic) Deldict(req types.DicDeldictReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
