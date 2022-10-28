package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetpersonbyidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetpersonbyidLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetpersonbyidLogic {
	return GetpersonbyidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetpersonbyidLogic) Getpersonbyid(req types.GetPersonByIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
