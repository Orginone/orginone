package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SimpletenantgrouptreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSimpletenantgrouptreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) SimpletenantgrouptreeLogic {
	return SimpletenantgrouptreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SimpletenantgrouptreeLogic) Simpletenantgrouptree(req types.SimpleTenantGroupTreeReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
