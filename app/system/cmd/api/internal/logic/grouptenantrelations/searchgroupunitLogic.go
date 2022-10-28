package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchgroupunitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchgroupunitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchgroupunitLogic {
	return SearchgroupunitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchgroupunitLogic) Searchgroupunit(req types.SearchGroupUnitReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
