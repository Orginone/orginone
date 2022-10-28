package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupAllUnitSocialCreditCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupAllUnitSocialCreditCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetGroupAllUnitSocialCreditCodeLogic {
	return GetGroupAllUnitSocialCreditCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupAllUnitSocialCreditCodeLogic) GetGroupAllUnitSocialCreditCode(req types.GetGroupAllUnitSocialCreditCodeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
