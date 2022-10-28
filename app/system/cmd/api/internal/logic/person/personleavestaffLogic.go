package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PersonleavestaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersonleavestaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) PersonleavestaffLogic {
	return PersonleavestaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersonleavestaffLogic) Personleavestaff(req types.PersonLeaveStaffReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
