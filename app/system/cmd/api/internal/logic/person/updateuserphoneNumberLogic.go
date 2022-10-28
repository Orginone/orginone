package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateuserphoneNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateuserphoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateuserphoneNumberLogic {
	return UpdateuserphoneNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateuserphoneNumberLogic) UpdateuserphoneNumber(req types.UpdateUserPhoneNumberReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
