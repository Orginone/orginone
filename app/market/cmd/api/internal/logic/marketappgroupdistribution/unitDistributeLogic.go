package marketappgroupdistribution

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnitDistributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnitDistributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) UnitDistributeLogic {
	return UnitDistributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnitDistributeLogic) UnitDistribute(req types.UnitDistributeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
