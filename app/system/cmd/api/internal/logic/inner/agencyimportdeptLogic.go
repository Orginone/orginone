package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyimportdeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyimportdeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyimportdeptLogic {
	return AgencyimportdeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyimportdeptLogic) Agencyimportdept(req types.AgencyImportDeptReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
