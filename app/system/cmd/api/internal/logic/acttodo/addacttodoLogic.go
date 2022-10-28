package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddacttodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddacttodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddacttodoLogic {
	return AddacttodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddacttodoLogic) Addacttodo(req types.AddacttodoReq) (resp *types.CommonResponse, err error) {
	return
}
