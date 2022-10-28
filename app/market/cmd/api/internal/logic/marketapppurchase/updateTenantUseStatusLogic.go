package marketapppurchase

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTenantUseStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTenantUseStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateTenantUseStatusLogic {
	return UpdateTenantUseStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTenantUseStatusLogic) UpdateTenantUseStatus(req types.UpdateTenantUseStatusReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
