package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryteninfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryteninfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryteninfoLogic {
	return QueryteninfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryteninfoLogic) Queryteninfo(req types.QueryTenInfoReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
