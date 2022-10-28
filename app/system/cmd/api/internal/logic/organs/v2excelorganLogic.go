package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2excelorganLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2excelorganLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2excelorganLogic {
	return V2excelorganLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2excelorganLogic) V2excelorgan(req types.V2ExcelOrganReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
