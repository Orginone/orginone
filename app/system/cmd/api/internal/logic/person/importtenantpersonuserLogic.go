package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImporttenantpersonuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImporttenantpersonuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImporttenantpersonuserLogic {
	return ImporttenantpersonuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImporttenantpersonuserLogic) Importtenantpersonuser(req types.ImportTenantPersonUserReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
