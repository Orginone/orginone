// Code generated by goctl. DO NOT EDIT!
// Source: market.proto

package market

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Market interface {
		// 获取单位应用
		FintTenantPurchase(ctx context.Context, in *TenantPurchaseReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取集团应用
		FintGroupPurchase(ctx context.Context, in *GroupPurchaseReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用市场应用分页
		FindMarketApp(ctx context.Context, in *MarketAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 提交应用市场应用注册
		SubmitMarketApp(ctx context.Context, in *CommonRpcReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取重新发布应用列表
		GetRedeployAppList(ctx context.Context, in *GetRedeployReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用详情
		GetMarketAppInfo(ctx context.Context, in *PrimaryKeyReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 集团租户应用购买列表
		GetGroupDistributeAppTenantList(ctx context.Context, in *GroupPurchaseAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取集团分发配置
		GetGroupDistributeConfigList(ctx context.Context, in *GroupAppByIdReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 单位应用退订
		UnitAppUnsubscribe(ctx context.Context, in *UnitAppUnsubscribeReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 集团应用退订
		GroupAppUnsubscribe(ctx context.Context, in *GroupAppUnsubscribeReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用已购租户
		GetAppDisTenantList(ctx context.Context, in *AppLinkReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用已购集团
		GetAppDisGroupList(ctx context.Context, in *AppLinkReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取通知列表
		FindMarkAppNoticeList(ctx context.Context, in *FindMarkAppNoticeListReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 根据ID列表移除应用通知
		RemoveMarkAppNoticeList(ctx context.Context, in *MarkAppNoticeIdsReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 根据ID列表获取应用通知
		FindMarkAppNoticeByIds(ctx context.Context, in *MarkAppNoticeIdsReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 单位获取应用
		UnitAppPurchas(ctx context.Context, in *UnitAppPurchasReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 集团获取应用
		GroupAppPurchas(ctx context.Context, in *GroupAppPurchasReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用对当前租户的分发情况
		GetAppPurchasConfig(ctx context.Context, in *AppPurchasConfigReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 应用上下架
		MarketAppOnOrOffSale(ctx context.Context, in *AppOnOrOffSaleReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 应用删除
		MarketAppDelete(ctx context.Context, in *DeleteAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 应用部署
		MarketAppDeploy(ctx context.Context, in *AppDeployReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 应用取消部署
		MarketAppCancelApply(ctx context.Context, in *AppCancelApplyReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 应用发布
		MarketAppPublish(ctx context.Context, in *CommonRpcReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用组件
		GetAppComponmentList(ctx context.Context, in *AppLinkReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 应用权限分配
		AppRoleDistribe(ctx context.Context, in *AppRoleDistribeRpcReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用权限分配数据
		GetAppRoleDistribe(ctx context.Context, in *GetAppRoleDistribeReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取用户使用的应用Ids
		GetUserAppIds(ctx context.Context, in *GetUserAppIdsReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用组件
		GetAppComponments(ctx context.Context, in *GetComponmentReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取用户应用组件
		GetUserTemplates(ctx context.Context, in *GetUserTemplateReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取用户应用组件ID
		GetUserTemplateId(ctx context.Context, in *GetUserTemplateIdReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 新增或修改通知
		UpdateMarketAppNotice(ctx context.Context, in *UpdateMarketAppNoticeReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 更新应用通知状态
		UpdateMarkAppNoticesStatus(ctx context.Context, in *UpdateMarkAppNoticesStatusReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 查询Portal模板
		GetMarketappcomponenttemplateList(ctx context.Context, in *MarketappcomponenttemplateListReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 集团取消分发
		MarketAppUnitCancelDistribute(ctx context.Context, in *MarketAppUnitCancelDistributeReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取常用应用菜单
		GetUsedAppMenu(ctx context.Context, in *GetUsedAppMenuReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 根据应用token获取应用角色信息
		ApiRoleList(ctx context.Context, in *ApiRoleLitReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 删除
		RemoveByRelation(ctx context.Context, in *RemoveByRelationReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 常用应用排序
		SortUsedApp(ctx context.Context, in *SortUsedAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 新增或修改
		SubmitMarketUsedApp(ctx context.Context, in *SubmitMarketUsedAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 获取应用审核列表分页
		FindMarketAppCheckList(ctx context.Context, in *FindMarketAppCheckListReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 通过或拒绝应用
		CheckRedeploy(ctx context.Context, in *CheckRedeployReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 通过或拒绝应用
		AppCheck(ctx context.Context, in *AppCheckReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
		// 查询app角色是否可以删除(发布时)
		CheckRoleCanDelete(ctx context.Context, in *CheckRoleReq, opts ...grpc.CallOption) (*CommonRpcRes, error)
	}

	defaultMarket struct {
		cli zrpc.Client
	}
)

func NewMarket(cli zrpc.Client) Market {
	return &defaultMarket{
		cli: cli,
	}
}

// 获取单位应用
func (m *defaultMarket) FintTenantPurchase(ctx context.Context, in *TenantPurchaseReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.FintTenantPurchase(ctx, in, opts...)
}

// 获取集团应用
func (m *defaultMarket) FintGroupPurchase(ctx context.Context, in *GroupPurchaseReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.FintGroupPurchase(ctx, in, opts...)
}

// 获取应用市场应用分页
func (m *defaultMarket) FindMarketApp(ctx context.Context, in *MarketAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.FindMarketApp(ctx, in, opts...)
}

// 提交应用市场应用注册
func (m *defaultMarket) SubmitMarketApp(ctx context.Context, in *CommonRpcReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.SubmitMarketApp(ctx, in, opts...)
}

// 获取重新发布应用列表
func (m *defaultMarket) GetRedeployAppList(ctx context.Context, in *GetRedeployReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetRedeployAppList(ctx, in, opts...)
}

// 获取应用详情
func (m *defaultMarket) GetMarketAppInfo(ctx context.Context, in *PrimaryKeyReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetMarketAppInfo(ctx, in, opts...)
}

// 集团租户应用购买列表
func (m *defaultMarket) GetGroupDistributeAppTenantList(ctx context.Context, in *GroupPurchaseAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetGroupDistributeAppTenantList(ctx, in, opts...)
}

// 获取集团分发配置
func (m *defaultMarket) GetGroupDistributeConfigList(ctx context.Context, in *GroupAppByIdReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetGroupDistributeConfigList(ctx, in, opts...)
}

// 单位应用退订
func (m *defaultMarket) UnitAppUnsubscribe(ctx context.Context, in *UnitAppUnsubscribeReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.UnitAppUnsubscribe(ctx, in, opts...)
}

// 集团应用退订
func (m *defaultMarket) GroupAppUnsubscribe(ctx context.Context, in *GroupAppUnsubscribeReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GroupAppUnsubscribe(ctx, in, opts...)
}

// 获取应用已购租户
func (m *defaultMarket) GetAppDisTenantList(ctx context.Context, in *AppLinkReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetAppDisTenantList(ctx, in, opts...)
}

// 获取应用已购集团
func (m *defaultMarket) GetAppDisGroupList(ctx context.Context, in *AppLinkReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetAppDisGroupList(ctx, in, opts...)
}

// 获取通知列表
func (m *defaultMarket) FindMarkAppNoticeList(ctx context.Context, in *FindMarkAppNoticeListReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.FindMarkAppNoticeList(ctx, in, opts...)
}

// 根据ID列表移除应用通知
func (m *defaultMarket) RemoveMarkAppNoticeList(ctx context.Context, in *MarkAppNoticeIdsReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.RemoveMarkAppNoticeList(ctx, in, opts...)
}

// 根据ID列表获取应用通知
func (m *defaultMarket) FindMarkAppNoticeByIds(ctx context.Context, in *MarkAppNoticeIdsReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.FindMarkAppNoticeByIds(ctx, in, opts...)
}

// 单位获取应用
func (m *defaultMarket) UnitAppPurchas(ctx context.Context, in *UnitAppPurchasReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.UnitAppPurchas(ctx, in, opts...)
}

// 集团获取应用
func (m *defaultMarket) GroupAppPurchas(ctx context.Context, in *GroupAppPurchasReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GroupAppPurchas(ctx, in, opts...)
}

// 获取应用对当前租户的分发情况
func (m *defaultMarket) GetAppPurchasConfig(ctx context.Context, in *AppPurchasConfigReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetAppPurchasConfig(ctx, in, opts...)
}

// 应用上下架
func (m *defaultMarket) MarketAppOnOrOffSale(ctx context.Context, in *AppOnOrOffSaleReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.MarketAppOnOrOffSale(ctx, in, opts...)
}

// 应用删除
func (m *defaultMarket) MarketAppDelete(ctx context.Context, in *DeleteAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.MarketAppDelete(ctx, in, opts...)
}

// 应用部署
func (m *defaultMarket) MarketAppDeploy(ctx context.Context, in *AppDeployReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.MarketAppDeploy(ctx, in, opts...)
}

// 应用取消部署
func (m *defaultMarket) MarketAppCancelApply(ctx context.Context, in *AppCancelApplyReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.MarketAppCancelApply(ctx, in, opts...)
}

// 应用发布
func (m *defaultMarket) MarketAppPublish(ctx context.Context, in *CommonRpcReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.MarketAppPublish(ctx, in, opts...)
}

// 获取应用组件
func (m *defaultMarket) GetAppComponmentList(ctx context.Context, in *AppLinkReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetAppComponmentList(ctx, in, opts...)
}

// 应用权限分配
func (m *defaultMarket) AppRoleDistribe(ctx context.Context, in *AppRoleDistribeRpcReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.AppRoleDistribe(ctx, in, opts...)
}

// 获取应用权限分配数据
func (m *defaultMarket) GetAppRoleDistribe(ctx context.Context, in *GetAppRoleDistribeReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetAppRoleDistribe(ctx, in, opts...)
}

// 获取用户使用的应用Ids
func (m *defaultMarket) GetUserAppIds(ctx context.Context, in *GetUserAppIdsReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetUserAppIds(ctx, in, opts...)
}

// 获取应用组件
func (m *defaultMarket) GetAppComponments(ctx context.Context, in *GetComponmentReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetAppComponments(ctx, in, opts...)
}

// 获取用户应用组件
func (m *defaultMarket) GetUserTemplates(ctx context.Context, in *GetUserTemplateReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetUserTemplates(ctx, in, opts...)
}

// 获取用户应用组件ID
func (m *defaultMarket) GetUserTemplateId(ctx context.Context, in *GetUserTemplateIdReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetUserTemplateId(ctx, in, opts...)
}

// 新增或修改通知
func (m *defaultMarket) UpdateMarketAppNotice(ctx context.Context, in *UpdateMarketAppNoticeReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.UpdateMarketAppNotice(ctx, in, opts...)
}

// 更新应用通知状态
func (m *defaultMarket) UpdateMarkAppNoticesStatus(ctx context.Context, in *UpdateMarkAppNoticesStatusReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.UpdateMarkAppNoticesStatus(ctx, in, opts...)
}

// 查询Portal模板
func (m *defaultMarket) GetMarketappcomponenttemplateList(ctx context.Context, in *MarketappcomponenttemplateListReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetMarketappcomponenttemplateList(ctx, in, opts...)
}

// 集团取消分发
func (m *defaultMarket) MarketAppUnitCancelDistribute(ctx context.Context, in *MarketAppUnitCancelDistributeReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.MarketAppUnitCancelDistribute(ctx, in, opts...)
}

// 获取常用应用菜单
func (m *defaultMarket) GetUsedAppMenu(ctx context.Context, in *GetUsedAppMenuReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.GetUsedAppMenu(ctx, in, opts...)
}

// 根据应用token获取应用角色信息
func (m *defaultMarket) ApiRoleList(ctx context.Context, in *ApiRoleLitReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.ApiRoleList(ctx, in, opts...)
}

// 删除
func (m *defaultMarket) RemoveByRelation(ctx context.Context, in *RemoveByRelationReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.RemoveByRelation(ctx, in, opts...)
}

// 常用应用排序
func (m *defaultMarket) SortUsedApp(ctx context.Context, in *SortUsedAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.SortUsedApp(ctx, in, opts...)
}

// 新增或修改
func (m *defaultMarket) SubmitMarketUsedApp(ctx context.Context, in *SubmitMarketUsedAppReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.SubmitMarketUsedApp(ctx, in, opts...)
}

// 获取应用审核列表分页
func (m *defaultMarket) FindMarketAppCheckList(ctx context.Context, in *FindMarketAppCheckListReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.FindMarketAppCheckList(ctx, in, opts...)
}

// 通过或拒绝应用
func (m *defaultMarket) CheckRedeploy(ctx context.Context, in *CheckRedeployReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.CheckRedeploy(ctx, in, opts...)
}

// 通过或拒绝应用
func (m *defaultMarket) AppCheck(ctx context.Context, in *AppCheckReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.AppCheck(ctx, in, opts...)
}

// 查询app角色是否可以删除(发布时)
func (m *defaultMarket) CheckRoleCanDelete(ctx context.Context, in *CheckRoleReq, opts ...grpc.CallOption) (*CommonRpcRes, error) {
	client := NewMarketClient(m.cli.Conn())
	return client.CheckRoleCanDelete(ctx, in, opts...)
}
