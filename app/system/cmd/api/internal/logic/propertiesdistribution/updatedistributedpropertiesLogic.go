package propertiesdistribution

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatedistributedpropertiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatedistributedpropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatedistributedpropertiesLogic {
	return UpdatedistributedpropertiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatedistributedpropertiesLogic) Updatedistributedproperties(req types.UpdateDistributedPropertiesReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
