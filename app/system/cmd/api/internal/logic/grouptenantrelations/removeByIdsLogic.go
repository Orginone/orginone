package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveByIdsLogic {
	return RemoveByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveByIdsLogic) RemoveByIds(req types.RemoveByIdsReq2) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
