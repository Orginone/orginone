package allgroup

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetusermanagergroupidlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetusermanagergroupidlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetusermanagergroupidlistLogic {
	return GetusermanagergroupidlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetusermanagergroupidlistLogic) Getusermanagergroupidlist(req types.GetUserManagerGroupIdListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
