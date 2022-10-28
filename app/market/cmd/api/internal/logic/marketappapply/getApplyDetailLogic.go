package marketappapply

import (
	"context"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplyDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetApplyDetailLogic {
	return GetApplyDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplyDetailLogic) GetApplyDetail(req types.GetApplyDetailReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
