// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "orginone/app/market/cmd/api/internal/handler/api"
	groupappdistributiondata "orginone/app/market/cmd/api/internal/handler/groupappdistributiondata"
	kubesphere "orginone/app/market/cmd/api/internal/handler/kubesphere"
	marketapp "orginone/app/market/cmd/api/internal/handler/marketapp"
	marketappalert "orginone/app/market/cmd/api/internal/handler/marketappalert"
	marketappapply "orginone/app/market/cmd/api/internal/handler/marketappapply"
	marketappcomponent "orginone/app/market/cmd/api/internal/handler/marketappcomponent"
	marketappcomponenttemplate "orginone/app/market/cmd/api/internal/handler/marketappcomponenttemplate"
	marketappgroupdistribution "orginone/app/market/cmd/api/internal/handler/marketappgroupdistribution"
	marketappgroupdistributionrelation "orginone/app/market/cmd/api/internal/handler/marketappgroupdistributionrelation"
	marketappkeysecret "orginone/app/market/cmd/api/internal/handler/marketappkeysecret"
	marketappnotice "orginone/app/market/cmd/api/internal/handler/marketappnotice"
	marketapppurchase "orginone/app/market/cmd/api/internal/handler/marketapppurchase"
	marketapprole "orginone/app/market/cmd/api/internal/handler/marketapprole"
	marketapproledistributionnew "orginone/app/market/cmd/api/internal/handler/marketapproledistributionnew"
	marketapprolemenu "orginone/app/market/cmd/api/internal/handler/marketapprolemenu"
	marketappusertemplate "orginone/app/market/cmd/api/internal/handler/marketappusertemplate"
	marketmenu "orginone/app/market/cmd/api/internal/handler/marketmenu"
	marketmenuusersort "orginone/app/market/cmd/api/internal/handler/marketmenuusersort"
	marketusedapp "orginone/app/market/cmd/api/internal/handler/marketusedapp"
	portal "orginone/app/market/cmd/api/internal/handler/portal"
	redeploy "orginone/app/market/cmd/api/internal/handler/redeploy"
	"orginone/app/market/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/role/list", // TEST OK
				Handler: api.RolelistHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/api"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getConfig", // TEST OK
				Handler: groupappdistributiondata.GetConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: groupappdistributiondata.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listPage",
				Handler: groupappdistributiondata.ListPageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/groupappdistributiondata"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/getToken",
				Handler: kubesphere.GetTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/kubesphere"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/appCheck", 
				Handler: marketapp.AppCheckHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cancelApply", // TEST OK
				Handler: marketapp.CancelApplyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/deleteApp", // TEST OK
				Handler: marketapp.DeleteAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketapp.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAll", // TEST OK
				Handler: marketapp.GetAllHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAllApp", // TEST OK
				Handler: marketapp.GetAllAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/freezeApp",
				Handler: marketapp.FreezeAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getOnSaleAppNum",
				Handler: marketapp.GetOnSaleAppNumHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getRedeployAppList", // TEST OK
				Handler: marketapp.GetRedeployAppListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getTenantName",
				Handler: marketapp.GetTenantNameHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list", // TEST OK
				Handler: marketapp.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list/check", // TEST OK
				Handler: marketapp.ListcheckHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listAll",
				Handler: marketapp.ListAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/listAppBySort",
				Handler: marketapp.ListAppBySortHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listAppList",
				Handler: marketapp.ListAppListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listByDistributeTenantId",
				Handler: marketapp.ListByDistributeTenantIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listByPurchaseGroupId", // TEST OK
				Handler: marketapp.ListByPurchaseGroupIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listByPurchaseTenantId", // TEST OK
				Handler: marketapp.ListByPurchaseTenantIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/reformApp",
				Handler: marketapp.ReformAppHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketapp.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/setOffSale", // TEST OK
				Handler: marketapp.SetOffSaleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/setOnSale", // TEST OK
				Handler: marketapp.SetOnSaleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketapp.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submitAll",
				Handler: marketapp.SubmitAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submitFirst", // TEST OK
				Handler: marketapp.SubmitFirstHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submitSecond", // TEST OK
				Handler: marketapp.SubmitSecondHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submitThird", // TEST OK
				Handler: marketapp.SubmitThirdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/token/detail",
				Handler: marketapp.TokendetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/token/listAppBySort",
				Handler: marketapp.TokenlistAppBySortHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/token/listWithoutToken",
				Handler: marketapp.TokenlistWithoutTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketapp"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/cancelCheckAlert",
				Handler: marketappalert.CancelCheckAlertHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/checkAlert",
				Handler: marketappalert.CheckAlertHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketappalert.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/myAlertList",
				Handler: marketappalert.MyAlertListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketappalert.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sendAlert",
				Handler: marketappalert.SendAlertHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketappalert.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappalert"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/apply",
				Handler: marketappapply.ApplyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/applyCancel",
				Handler: marketappapply.ApplyCancelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/applyhandle",
				Handler: marketappapply.ApplyhandleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketappapply.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getApplyDetail",
				Handler: marketappapply.GetApplyDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getMyList",
				Handler: marketappapply.GetMyListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getSecretByKey",
				Handler: marketappapply.GetSecretByKeyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketappapply.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketappapply.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketappapply.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/test",
				Handler: marketappapply.TestHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/updateApply",
				Handler: marketappapply.UpdateApplyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappapply"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getComponentListByAppId", // TEST OK
				Handler: marketappcomponent.GetComponentListByAppIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list", // TEST OK
				Handler: marketappcomponent.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketappcomponent.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketappcomponent.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/appComponentGroupByUser",
				Handler: marketappcomponent.AppComponentGroupByUserHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappcomponent"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getComponentByTemplateId",
				Handler: marketappcomponenttemplate.GetComponentByTemplateIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getUserTemplate", // TEST OK
				Handler: marketappcomponenttemplate.GetUserTemplateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list", // TEST OK
				Handler: marketappcomponenttemplate.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketappcomponenttemplate.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/setDefault",
				Handler: marketappcomponenttemplate.SetDefaultHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketappcomponenttemplate.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappcomponenttemplate"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketappgroupdistribution.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/distributeAppTenantList", // TEST OK
				Handler: marketappgroupdistribution.DistributeAppTenantListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributedGroupIdList",
				Handler: marketappgroupdistribution.GetDistributedGroupIdListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getGroupDistributeTenantApp",
				Handler: marketappgroupdistribution.GetGroupDistributeTenantAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getGroupIdList",
				Handler: marketappgroupdistribution.GetGroupIdListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketappgroupdistribution.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketappgroupdistribution.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submitAll",
				Handler: marketappgroupdistribution.SubmitAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/unitCancelDistribute", // TEST OK
				Handler: marketappgroupdistribution.UnitCancelDistributeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/unitDistribute",
				Handler: marketappgroupdistribution.UnitDistributeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappgroupdistribution"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getRelation",
				Handler: marketappgroupdistributionrelation.GetRelationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappgroupdistributionrelation"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketappkeysecret.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getByAppId",
				Handler: marketappkeysecret.GetByAppIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketappkeysecret.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listAll",
				Handler: marketappkeysecret.ListAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketappkeysecret.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketappkeysecret.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappkeysecret"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/cancelReleaseNote", // TEST OK
				Handler: marketappnotice.CancelReleaseNoteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list", // TEST OK
				Handler: marketappnotice.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/myNoticeList",
				Handler: marketappnotice.MyNoticeListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/releaseNote", // TEST OK
				Handler: marketappnotice.ReleaseNoteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove", // TEST OK
				Handler: marketappnotice.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit", // TEST OK
				Handler: marketappnotice.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappnotice"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketapppurchase.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/lists",
				Handler: marketapppurchase.GetlistsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAppPurchaseConfig", // TEST OK
				Handler: marketapppurchase.GetAppPurchaseConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketapppurchase.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listAppByTenantId",
				Handler: marketapppurchase.ListAppByTenantIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listByGroupIdVO", // TEST OK
				Handler: marketapppurchase.ListByGroupIdVOHandler(serverCtx),
			},
			{
				Method:  http.MethodGet, // TEST OK
				Path:    "/listUnitAndPersonalVO",
				Handler: marketapppurchase.ListUnitAndPersonalVOHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listVO",
				Handler: marketapppurchase.ListVOHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/purchaseAppGroupList", // TEST OK
				Handler: marketapppurchase.PurchaseAppGroupListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/purchaseAppTenantList", // TEST OK
				Handler: marketapppurchase.PurchaseAppTenantListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/purchaseByGroups", // TEST OK
				Handler: marketapppurchase.PurchaseByGroupsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketapppurchase.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit", // TEST OK
				Handler: marketapppurchase.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/unitunsubscribe", // TEST OK
				Handler: marketapppurchase.UnitunsubscribeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/unsubscribe", // TEST OK
				Handler: marketapppurchase.UnsubscribeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateGroupUseStatus",
				Handler: marketapppurchase.UpdateGroupUseStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateTenantUseStatus",
				Handler: marketapppurchase.UpdateTenantUseStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/withdrawGroupAuthority",
				Handler: marketapppurchase.WithdrawGroupAuthorityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/withdrawTenantAuthority",
				Handler: marketapppurchase.WithdrawTenantAuthorityHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketapppurchase"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketapprole.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getRoleList",
				Handler: marketapprole.GetRoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getRoleNameList",
				Handler: marketapprole.GetRoleNameListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketapprole.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketapprole.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketapprole.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketapprole"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/apptoken/verify/user",
				Handler: marketapproledistributionnew.ApptokenverifyuserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketapproledistributionnew.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAppDistributedRole", // TEST OK
				Handler: marketapproledistributionnew.GetAppDistributedRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAppIdByUserId", // TEST OK
				Handler: marketapproledistributionnew.GetAppIdByUserIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributeUserNum",
				Handler: marketapproledistributionnew.GetDistributeUserNumHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributedAgency",
				Handler: marketapproledistributionnew.GetDistributedAgencyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributedJob",
				Handler: marketapproledistributionnew.GetDistributedJobHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributedPersonListByAppId",
				Handler: marketapproledistributionnew.GetDistributedPersonListByAppIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributionAgency",
				Handler: marketapproledistributionnew.GetDistributionAgencyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributionJob",
				Handler: marketapproledistributionnew.GetDistributionJobHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributionPerson",
				Handler: marketapproledistributionnew.GetDistributionPersonHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getDistributionUser",
				Handler: marketapproledistributionnew.GetDistributionUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getRoleByUserId",
				Handler: marketapproledistributionnew.GetRoleByUserIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getUserRoleList",
				Handler: marketapproledistributionnew.GetUserRoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/isDistributed",
				Handler: marketapproledistributionnew.IsDistributedHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketapproledistributionnew.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketapproledistributionnew.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketapproledistributionnew.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit/v2", // TEST OK
				Handler: marketapproledistributionnew.Submitv2Handler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketapproledistributionnew"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketapprolemenu.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketapprolemenu.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketapprolemenu.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketapprolemenu.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketapprolemenu"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getDefaultTemplate", // TEST OK
				Handler: marketappusertemplate.GetDefaultTemplateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getTemplateIdByLoginInUse",
				Handler: marketappusertemplate.GetTemplateIdByLoginInUseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getTemplateIdByUserIdInUse",
				Handler: marketappusertemplate.GetTemplateIdByUserIdInUseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getTemplateIdListByLogin",
				Handler: marketappusertemplate.GetTemplateIdListByLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getTemplateIdListByUserId",
				Handler: marketappusertemplate.GetTemplateIdListByUserIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketappusertemplate.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketappusertemplate.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/setDefaultTemplate",
				Handler: marketappusertemplate.SetDefaultTemplateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/setGroupTenantDefaultTemplate",
				Handler: marketappusertemplate.SetGroupTenantDefaultTemplateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketappusertemplate.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketappusertemplate"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketmenu.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAppMenuList",
				Handler: marketmenu.GetAppMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getMenuList",
				Handler: marketmenu.GetMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getMobileMenuList",
				Handler: marketmenu.GetMobileMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getParentMobileMenuList",
				Handler: marketmenu.GetParentMobileMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getSonMobileMenuList",
				Handler: marketmenu.GetSonMobileMenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketmenu.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketmenu.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketmenu.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketmenu"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: marketmenuusersort.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getUserMenuSortList",
				Handler: marketmenuusersort.GetUserMenuSortListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketmenuusersort.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketmenuusersort.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sortUserMenu",
				Handler: marketmenuusersort.SortUserMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sortUserMenuByUnit",
				Handler: marketmenuusersort.SortUserMenuByUnitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: marketmenuusersort.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketmenuusersort"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getUsedApp",
				Handler: marketusedapp.GetUsedAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getUsedAppMenu", // TEST OK
				Handler: marketusedapp.GetUsedAppMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: marketusedapp.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: marketusedapp.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/removeByRelation", // TEST OK
				Handler: marketusedapp.RemoveByRelationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sortUsedApp", // TEST OK
				Handler: marketusedapp.SortUsedAppHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit", //TEST OK
				Handler: marketusedapp.SubmitHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/marketusedapp"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/addUniApp",
				Handler: portal.AddUniAppHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cancelUniApp",
				Handler: portal.CancelUniAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getAllApp",
				Handler: portal.GetAllAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getCommonUseApp",
				Handler: portal.GetCommonUseAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getRemainApp",
				Handler: portal.GetRemainAppHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getUserPortal",
				Handler: portal.GetUserPortalHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: portal.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: portal.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/saveOrUpdateAllPortal",
				Handler: portal.SaveOrUpdateAllPortalHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit",
				Handler: portal.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v2/getCommonUseApp",
				Handler: portal.V2getCommonUseAppHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/portal"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/checkRedeploy", // TEST OK
				Handler: redeploy.CheckRedeployHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/isAddRole",
				Handler: redeploy.IsAddRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/isDeleteRole", // TEST OK
				Handler: redeploy.IsDeleteRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: redeploy.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listAll",
				Handler: redeploy.ListAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/redeploy", // TEST OK
				Handler: redeploy.RedeployHandler(serverCtx), 
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-market/redeploy"),
	)
}
