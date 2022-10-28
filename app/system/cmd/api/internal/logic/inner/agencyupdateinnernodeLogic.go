package inner

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyupdateinnernodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyupdateinnernodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyupdateinnernodeLogic {
	return AgencyupdateinnernodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyupdateinnernodeLogic) Agencyupdateinnernode(req types.AgencyUpdateInnerNodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
