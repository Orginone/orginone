package logic

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema"
	"orginone/common/schema/asuser"
	"orginone/common/schema/predicate"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserListPageLogic {
	return &FindUserListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询用户
func (l *FindUserListPageLogic) FindUserListPage(in *user.FindUserListPageReq) (*user.CommonRpcRes, error) {
	preQuery := make([]predicate.AsUser, 0)
	if in.Id != -1 {
		preQuery = append(preQuery, asuser.ID(in.Id))
	} else {
		preQuery = append(preQuery,
			asuser.UserNameContains(in.UserName),
			asuser.PhoneNumberContains(in.PhoneNumber),
			asuser.TenantCode(in.TenantCode),
			asuser.IsDeleted(in.IsDeleted))
	}
	query := l.svcCtx.UserStore.AsUser.Query().Where(preQuery...)
	count, err := query.Count(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	userEntitys, err := query.Order(schema.Asc(asuser.FieldID)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return user.PageResult(in.Page, int64(count), userEntitys, err)
}
