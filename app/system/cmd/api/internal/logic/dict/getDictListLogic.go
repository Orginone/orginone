package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDictListLogic {
	return GetDictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDictListLogic) GetDictList(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
