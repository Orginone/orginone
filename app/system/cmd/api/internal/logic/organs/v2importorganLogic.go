package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2importorganLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2importorganLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2importorganLogic {
	return V2importorganLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2importorganLogic) V2importorgan(req types.V2ImportOrganReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
