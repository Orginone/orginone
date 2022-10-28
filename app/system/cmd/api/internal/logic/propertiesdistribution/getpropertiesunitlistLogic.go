package propertiesdistribution

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetpropertiesunitlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetpropertiesunitlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetpropertiesunitlistLogic {
	return GetpropertiesunitlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetpropertiesunitlistLogic) Getpropertiesunitlist(req types.GetPropertiesUnitListReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindPropertiesUnitList(l.ctx, &system.FindPropertiesUnitListReq{
		PropertiesId: req.PropertiesId,
		Page: &system.PageRequest{
			Offset:      (req.Current - 1) * req.Size,
			Limit:       req.Size,
			SearchCount: true,
		},
	}))
}
