package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallusableuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallusableuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallusableuserLogic {
	return GetallusableuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallusableuserLogic) Getallusableuser(req types.GetAllUsableUserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
