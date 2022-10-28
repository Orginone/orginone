package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportPhoneNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportPhoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportPhoneNumberLogic {
	return ImportPhoneNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportPhoneNumberLogic) ImportPhoneNumber(req types.ImportPhoneNumberReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
