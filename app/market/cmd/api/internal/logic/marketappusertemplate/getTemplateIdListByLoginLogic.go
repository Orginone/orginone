package marketappusertemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplateIdListByLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplateIdListByLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTemplateIdListByLoginLogic {
	return GetTemplateIdListByLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplateIdListByLoginLogic) GetTemplateIdListByLogin(req types.Nil) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
