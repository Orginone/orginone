package marketappusertemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplateIdListByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplateIdListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTemplateIdListByUserIdLogic {
	return GetTemplateIdListByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplateIdListByUserIdLogic) GetTemplateIdListByUserId(req types.GetTemplateIdListByUserIdReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
