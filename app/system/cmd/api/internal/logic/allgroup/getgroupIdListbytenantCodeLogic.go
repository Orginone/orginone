package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetgroupIdListbytenantCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetgroupIdListbytenantCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetgroupIdListbytenantCodeLogic {
	return GetgroupIdListbytenantCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetgroupIdListbytenantCodeLogic) GetgroupIdListbytenantCode(req types.GetGroupIdListByTenantCodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
