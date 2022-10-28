package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetalluserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetalluserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetalluserLogic {
	return GetalluserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetalluserLogic) Getalluser(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
