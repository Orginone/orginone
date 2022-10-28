package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatefathergroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatefathergroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatefathergroupLogic {
	return UpdatefathergroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatefathergroupLogic) Updatefathergroup(req types.UpdateFatherGroupReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
