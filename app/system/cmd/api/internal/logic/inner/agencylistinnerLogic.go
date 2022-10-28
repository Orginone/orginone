package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencylistinnerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencylistinnerLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencylistinnerLogic {
	return AgencylistinnerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencylistinnerLogic) Agencylistinner(req types.AgencyListInnerReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
