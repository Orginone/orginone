package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterbydictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterbydictLogic(ctx context.Context, svcCtx *svc.ServiceContext) FilterbydictLogic {
	return FilterbydictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterbydictLogic) Filterbydict(req types.DicFilterbydictReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
