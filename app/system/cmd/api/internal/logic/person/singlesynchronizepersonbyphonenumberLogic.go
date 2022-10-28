package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SinglesynchronizepersonbyphonenumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSinglesynchronizepersonbyphonenumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) SinglesynchronizepersonbyphonenumberLogic {
	return SinglesynchronizepersonbyphonenumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SinglesynchronizepersonbyphonenumberLogic) Singlesynchronizepersonbyphonenumber(req types.SingleSynchronizePersonByPhoneNumberReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
