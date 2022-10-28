package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyinnerAgencymodelv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyinnerAgencymodelv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyinnerAgencymodelv2Logic {
	return AgencyinnerAgencymodelv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyinnerAgencymodelv2Logic) AgencyinnerAgencymodelv2(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
