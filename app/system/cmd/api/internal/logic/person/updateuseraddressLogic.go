package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateuseraddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateuseraddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateuseraddressLogic {
	return UpdateuseraddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateuseraddressLogic) Updateuseraddress(req types.UpdateUserAddressReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
