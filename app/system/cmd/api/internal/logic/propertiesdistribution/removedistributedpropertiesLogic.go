package propertiesdistribution

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovedistributedpropertiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovedistributedpropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemovedistributedpropertiesLogic {
	return RemovedistributedpropertiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovedistributedpropertiesLogic) Removedistributedproperties(req types.RemoveDistributedPropertiesReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
