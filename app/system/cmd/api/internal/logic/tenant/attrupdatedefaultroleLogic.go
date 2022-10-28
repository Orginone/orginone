package tenant

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttrupdatedefaultroleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttrupdatedefaultroleLogic(ctx context.Context, svcCtx *svc.ServiceContext) AttrupdatedefaultroleLogic {
	return AttrupdatedefaultroleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttrupdatedefaultroleLogic) Attrupdatedefaultrole(req types.AttrUpdateDefaultRoleReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
