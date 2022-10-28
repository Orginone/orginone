package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectJobIdByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectJobIdByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SelectJobIdByUserIdLogic {
	return SelectJobIdByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectJobIdByUserIdLogic) SelectJobIdByUserId(req types.SelectJobIdByUserIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
