package marketapproledistributionnew

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApptokenverifyuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApptokenverifyuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) ApptokenverifyuserLogic {
	return ApptokenverifyuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApptokenverifyuserLogic) Apptokenverifyuser(req types.ApptokenVerifyUserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
