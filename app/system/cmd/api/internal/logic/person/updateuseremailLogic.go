package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateuseremailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateuseremailLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateuseremailLogic {
	return UpdateuseremailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateuseremailLogic) Updateuseremail(req types.UpdateUserEmailReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
