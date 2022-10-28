package marketappusertemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplateIdByLoginInUseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplateIdByLoginInUseLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTemplateIdByLoginInUseLogic {
	return GetTemplateIdByLoginInUseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplateIdByLoginInUseLogic) GetTemplateIdByLoginInUse(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
