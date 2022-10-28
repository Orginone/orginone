package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateadmintenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateadmintenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateadmintenantLogic {
	return UpdateadmintenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateadmintenantLogic) Updateadmintenant(req types.UpdateAdminTenantReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
