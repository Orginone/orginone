package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallusersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallusersLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallusersLogic {
	return GetallusersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallusersLogic) Getallusers(req types.GetAllUsersReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
