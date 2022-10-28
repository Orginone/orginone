package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2updateunitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2updateunitLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2updateunitLogic {
	return V2updateunitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2updateunitLogic) V2updateunit(req types.V2UpdateUnitReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
