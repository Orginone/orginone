package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetpersonbylikeNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetpersonbylikeNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetpersonbylikeNameLogic {
	return GetpersonbylikeNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetpersonbylikeNameLogic) GetpersonbylikeName(req types.GetPersonByLikeNameReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
