package logic

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserTenantUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserTenantUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserTenantUserLogic {
	return &FindUserTenantUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询用户
func (l *FindUserTenantUserLogic) FindUserTenantUser(in *user.FindserTenantUserReq) (*user.CommonRpcRes, error) {
	query := l.svcCtx.UserStore.AsUser.Query().Where(asuser.UserNameContains(in.UserName),
		asuser.TenantCode(in.TenantCode), asuser.IsDeleted(0))
	if in.Id != -1 {
		query.Where(asuser.IDNEQ(in.Id))
	}
	if len(in.TenantApplyinStateRange) > 0 {
		query.Where(asuser.TenantApplyingStateIn(in.TenantApplyinStateRange...))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	userEntitys, err := query.Order(schema.Desc(asuser.FieldCreateTime)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return user.PageResult(in.Page, int64(count), userEntitys, err)
}
