// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"orginone/app/user/cmd/rpc/internal/logic"
	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 用户注册
func (s *UserServer) UserReReg(ctx context.Context, in *user.UserReRegReq) (*user.CommonRpcRes, error) {
	l := logic.NewUserReRegLogic(ctx, s.svcCtx)
	return l.UserReReg(in)
}

// 修改密码
func (s *UserServer) UpdatePasswd(ctx context.Context, in *user.UpdatePasswdReq) (*user.CommonRpcRes, error) {
	l := logic.NewUpdatePasswdLogic(ctx, s.svcCtx)
	return l.UpdatePasswd(in)
}

// 切换主单位
func (s *UserServer) SwitchTenantByCode(ctx context.Context, in *user.SwitchTenantReq) (*user.CommonRpcRes, error) {
	l := logic.NewSwitchTenantByCodeLogic(ctx, s.svcCtx)
	return l.SwitchTenantByCode(in)
}

// 查找用户
func (s *UserServer) FindUserByAccount(ctx context.Context, in *user.UserByAccountReq) (*user.CommonRpcRes, error) {
	l := logic.NewFindUserByAccountLogic(ctx, s.svcCtx)
	return l.FindUserByAccount(in)
}

// 加入租户
func (s *UserServer) JoinTeantByCode(ctx context.Context, in *user.TenantCodeReq) (*user.CommonRpcRes, error) {
	l := logic.NewJoinTeantByCodeLogic(ctx, s.svcCtx)
	return l.JoinTeantByCode(in)
}

// 分页查询用户
func (s *UserServer) FindUserListPage(ctx context.Context, in *user.FindUserListPageReq) (*user.CommonRpcRes, error) {
	l := logic.NewFindUserListPageLogic(ctx, s.svcCtx)
	return l.FindUserListPage(in)
}

// 分页查询用户
func (s *UserServer) FindUserTenantUser(ctx context.Context, in *user.FindserTenantUserReq) (*user.CommonRpcRes, error) {
	l := logic.NewFindUserTenantUserLogic(ctx, s.svcCtx)
	return l.FindUserTenantUser(in)
}

// 删除用户
func (s *UserServer) DeleteUserInfo(ctx context.Context, in *user.DeleteUserInfoReq) (*user.CommonRpcRes, error) {
	l := logic.NewDeleteUserInfoLogic(ctx, s.svcCtx)
	return l.DeleteUserInfo(in)
}

// 审核租户
func (s *UserServer) AuditUser(ctx context.Context, in *user.AuditUserReq) (*user.CommonRpcRes, error) {
	l := logic.NewAuditUserLogic(ctx, s.svcCtx)
	return l.AuditUser(in)
}

// 取消加入租户申请
func (s *UserServer) CancelJoinTenantRequest(ctx context.Context, in *user.CancelJoinTenantReq) (*user.CommonRpcRes, error) {
	l := logic.NewCancelJoinTenantRequestLogic(ctx, s.svcCtx)
	return l.CancelJoinTenantRequest(in)
}

// 更新用户信息
func (s *UserServer) UpdateUserInfo(ctx context.Context, in *user.CommonRpcRes) (*user.CommonRpcRes, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

// 重置用户密码
func (s *UserServer) ResetPassWord(ctx context.Context, in *user.CommonRpcRes) (*user.CommonRpcRes, error) {
	l := logic.NewResetPassWordLogic(ctx, s.svcCtx)
	return l.ResetPassWord(in)
}

// 删除用户
func (s *UserServer) RemoveUserByIds(ctx context.Context, in *user.RemoveUserByIdsReq) (*user.CommonRpcRes, error) {
	l := logic.NewRemoveUserByIdsLogic(ctx, s.svcCtx)
	return l.RemoveUserByIds(in)
}

// 获取应用信息
func (s *UserServer) GetMarketAppInfo(ctx context.Context, in *user.GetMarketAppInfoReq) (*user.CommonRpcRes, error) {
	l := logic.NewGetMarketAppInfoLogic(ctx, s.svcCtx)
	return l.GetMarketAppInfo(in)
}