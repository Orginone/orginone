package scenes

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerybynameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerybynameLogic(ctx context.Context, svcCtx *svc.ServiceContext) QuerybynameLogic {
	return QuerybynameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerybynameLogic) Querybyname(req types.QueryByNameReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
