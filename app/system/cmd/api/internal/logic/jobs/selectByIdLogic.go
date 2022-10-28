package jobs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SelectByIdLogic {
	return SelectByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectByIdLogic) SelectById(req types.SelectByIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
