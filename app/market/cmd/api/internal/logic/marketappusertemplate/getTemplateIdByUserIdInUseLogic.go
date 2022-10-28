package marketappusertemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplateIdByUserIdInUseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplateIdByUserIdInUseLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTemplateIdByUserIdInUseLogic {
	return GetTemplateIdByUserIdInUseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplateIdByUserIdInUseLogic) GetTemplateIdByUserIdInUse(req types.GetTemplateIdByUserIdInUseReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
