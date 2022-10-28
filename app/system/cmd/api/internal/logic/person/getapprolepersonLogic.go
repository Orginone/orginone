package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetapprolepersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetapprolepersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetapprolepersonLogic {
	return GetapprolepersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetapprolepersonLogic) Getapproleperson(req types.GetAppRolePersonReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
