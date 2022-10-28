package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictionaryLogic {
	return DictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryLogic) Dictionary(req types.DicDictionaryReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
