package marketappusertemplate

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetDefaultTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetDefaultTemplateLogic {
	return SetDefaultTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetDefaultTemplateLogic) SetDefaultTemplate(req types.SetDefaultTemplateReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
