package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetinitpersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetinitpersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetinitpersonLogic {
	return GetinitpersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetinitpersonLogic) Getinitperson(req types.GetInitPersonReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
