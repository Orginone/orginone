package marketappcomponenttemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComponentByTemplateIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetComponentByTemplateIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetComponentByTemplateIdLogic {
	return GetComponentByTemplateIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComponentByTemplateIdLogic) GetComponentByTemplateId(req types.GetComponentByTemplateIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
