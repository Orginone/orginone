package acttodo

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallmysendacttodoBySearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallmysendacttodoBySearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallmysendacttodoBySearchLogic {
	return GetallmysendacttodoBySearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallmysendacttodoBySearchLogic) GetallmysendacttodoBySearch(req types.GetallmysendacttodoBySearchReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
