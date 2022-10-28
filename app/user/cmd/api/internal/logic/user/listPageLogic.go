package user

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListPageLogic {
	return ListPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPageLogic) ListPage(req types.ListPageReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return &types.CommonResponse{}, err
	}
	return types.PageResult(l.svcCtx.UserRpc.FindUserListPage(l.ctx, &user.FindUserListPageReq{
		Page: &user.PageRequest{
			Limit:       int64(req.Size),
			Offset:      int64((req.Current - 1) * req.Size),
			SearchCount: true,
		},
		PhoneNumber: req.PhoneNumber,
		UserName:    req.UserName,
		TenantCode:  tenantCode,
		IsDeleted:   req.IsDeleted,
		Id:          req.Id,
	}))
}
