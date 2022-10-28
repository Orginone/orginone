package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDefaultGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDefaultGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetDefaultGroupLogic {
	return GetDefaultGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDefaultGroupLogic) GetDefaultGroup(req []int64) (resp *types.CommonResponse, err error) {
	return types.JsonResult(tools.ArrIsExist(req, 1354721804653875201) ||
		tools.ArrIsExist(req, 1389828001400287233) ||
		tools.ArrIsExist(req, 1389828227129339905), nil)
}
