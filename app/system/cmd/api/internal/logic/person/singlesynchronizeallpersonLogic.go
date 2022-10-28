package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SinglesynchronizeallpersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSinglesynchronizeallpersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) SinglesynchronizeallpersonLogic {
	return SinglesynchronizeallpersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SinglesynchronizeallpersonLogic) Singlesynchronizeallperson(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
