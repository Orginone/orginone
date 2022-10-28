package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddacthistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddacthistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddacthistoryLogic {
	return AddacthistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddacthistoryLogic) Addacthistory(req types.AddacthistoryReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
