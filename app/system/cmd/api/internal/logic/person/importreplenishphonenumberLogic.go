package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportreplenishphonenumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportreplenishphonenumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportreplenishphonenumberLogic {
	return ImportreplenishphonenumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportreplenishphonenumberLogic) Importreplenishphonenumber(req types.ImportReplenishPhoneNumberReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
