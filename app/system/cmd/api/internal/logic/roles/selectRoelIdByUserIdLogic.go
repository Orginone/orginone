package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectRoelIdByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectRoelIdByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SelectRoelIdByUserIdLogic {
	return SelectRoelIdByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectRoelIdByUserIdLogic) SelectRoelIdByUserId(req types.SelectRoelIdByUserIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
