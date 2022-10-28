package marketappgroupdistribution

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetGroupIdListLogic {
	return GetGroupIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupIdListLogic) GetGroupIdList(req types.GetGroupIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
