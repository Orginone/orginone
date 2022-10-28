package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdategroupinformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdategroupinformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdategroupinformationLogic {
	return UpdategroupinformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdategroupinformationLogic) Updategroupinformation(req types.UpdateGroupInformationReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
