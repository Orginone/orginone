package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessMessageUserIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessMessageUserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcessMessageUserIdsLogic {
	return ProcessMessageUserIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessMessageUserIdsLogic) ProcessMessageUserIds(req types.ProcessMessageUserIdsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
