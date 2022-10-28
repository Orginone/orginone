package roles

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallrolesPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallrolesPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallrolesPageLogic {
	return GetallrolesPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallrolesPageLogic) GetallrolesPage(req types.GetAllRolesPageReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
