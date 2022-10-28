package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserGetImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserGetImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserGetImgLogic {
	return UpdateUserGetImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserGetImgLogic) UpdateUserGetImg(req types.UpdateUserGetImgReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
