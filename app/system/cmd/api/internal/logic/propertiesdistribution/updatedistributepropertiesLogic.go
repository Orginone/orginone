package propertiesdistribution

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatedistributepropertiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatedistributepropertiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatedistributepropertiesLogic {
	return UpdatedistributepropertiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatedistributepropertiesLogic) Updatedistributeproperties(req types.UpdateDistributePropertiesReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
