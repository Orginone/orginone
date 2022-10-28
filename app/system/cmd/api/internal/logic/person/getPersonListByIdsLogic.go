package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonListByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPersonListByIdsLogic {
	return GetPersonListByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonListByIdsLogic) GetPersonListByIds(req types.GetPersonListByIdsReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
