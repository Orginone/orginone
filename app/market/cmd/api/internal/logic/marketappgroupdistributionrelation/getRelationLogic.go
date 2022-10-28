package marketappgroupdistributionrelation

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRelationLogic {
	return GetRelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRelationLogic) GetRelation(req types.GetRelationReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
