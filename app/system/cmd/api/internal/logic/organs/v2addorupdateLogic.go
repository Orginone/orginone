package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2addorupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2addorupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2addorupdateLogic {
	return V2addorupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2addorupdateLogic) V2addorupdate(req types.V2AddOrUpdateReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
