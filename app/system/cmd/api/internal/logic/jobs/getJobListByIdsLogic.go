package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJobListByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJobListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetJobListByIdsLogic {
	return GetJobListByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJobListByIdsLogic) GetJobListByIds(req types.GetJobListByIdsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
