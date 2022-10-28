package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Agencylistv3Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencylistv3Logic(ctx context.Context, svcCtx *svc.ServiceContext) Agencylistv3Logic {
	return Agencylistv3Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Agencylistv3Logic) Agencylistv3(req types.AgencyListV3Req) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
