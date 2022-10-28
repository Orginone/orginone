package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisbandallgroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDisbandallgroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) DisbandallgroupLogic {
	return DisbandallgroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisbandallgroupLogic) Disbandallgroup(req types.DisbandAllGroupReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
