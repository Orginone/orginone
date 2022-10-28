package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonjoinstaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersonjoinstaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) PersonjoinstaffLogic {
	return PersonjoinstaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersonjoinstaffLogic) Personjoinstaff(req types.PersonJoinStaffReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
