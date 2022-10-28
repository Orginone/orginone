package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGroupUnitByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchGroupUnitByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchGroupUnitByNameLogic {
	return SearchGroupUnitByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchGroupUnitByNameLogic) SearchGroupUnitByName(req types.SearchGroupUnitByNameReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
