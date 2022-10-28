package marketappapply

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateApplyLogic {
	return UpdateApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateApplyLogic) UpdateApply(req types.UpdateApplyReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
