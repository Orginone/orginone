package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectJobsByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectJobsByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SelectJobsByUserIdLogic {
	return SelectJobsByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectJobsByUserIdLogic) SelectJobsByUserId(req types.SelectJobsByUserIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
