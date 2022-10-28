package tenant

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttrgetdefaultroleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttrgetdefaultroleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) AttrgetdefaultroleIdLogic {
	return AttrgetdefaultroleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttrgetdefaultroleIdLogic) AttrgetdefaultroleId(req types.AttrGetDefaultRoleIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
