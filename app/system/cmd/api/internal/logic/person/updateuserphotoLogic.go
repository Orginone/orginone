package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateuserphotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateuserphotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateuserphotoLogic {
	return UpdateuserphotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateuserphotoLogic) Updateuserphoto(req types.UpdateUserPhotoReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
