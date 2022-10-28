package handler

import (
	"net/http"

	"orginone/app/system/cmd/api/internal/logic/inner"
	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AgencyadddeptbyappidHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyAddDeptByAppidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyadddeptbyappidLogic(r.Context(), ctx)
		resp, err := l.Agencyadddeptbyappid(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyaddinneruserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyAddInnerUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyaddinneruserLogic(r.Context(), ctx)
		resp, err := l.Agencyaddinneruser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyaddnodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyAddNodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyaddnodeLogic(r.Context(), ctx)
		resp, err := l.Agencyaddnode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyaddnodesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyAddNodesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyaddnodesLogic(r.Context(), ctx)
		resp, err := l.Agencyaddnodes(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyaddUsersToAgencyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyAddUsersToAgencyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyaddUsersToAgencyLogic(r.Context(), ctx)
		resp, err := l.AgencyaddUsersToAgency(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencydeletejobdeptusersHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyDeleteJobDeptUsersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencydeletejobdeptusersLogic(r.Context(), ctx)
		resp, err := l.Agencydeletejobdeptusers(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencydeletenodesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyDeleteNodesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencydeletenodesLogic(r.Context(), ctx)
		resp, err := l.Agencydeletenodes(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencydeleteUsersToAgencyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyDeleteUsersToAgencyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencydeleteUsersToAgencyLogic(r.Context(), ctx)
		resp, err := l.AgencydeleteUsersToAgency(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencydeletedinnernodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyDeletedInnerNodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencydeletedinnernodeLogic(r.Context(), ctx)
		resp, err := l.Agencydeletedinnernode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencydistributedeptpersonsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyDistributeDeptPersonsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencydistributedeptpersonsLogic(r.Context(), ctx)
		resp, err := l.Agencydistributedeptpersons(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyexportinnerAgencyv2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyExportInnerAgencyV2Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyexportinnerAgencyv2Logic(r.Context(), ctx)
		resp, err := l.AgencyexportinnerAgencyv2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyexportInnerAgencyDataHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyExportInnerAgencyDataReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		}
		if err := tools.UnmarshaJsonIOData(r.Body, &req.InnerAgencyList); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		l := inner.NewAgencyexportInnerAgencyDataLogic(r.Context(), ctx)
		resp, err := l.AgencyexportInnerAgencyData(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			w.Write(resp)
		}
	}
}

func AgencygetallpNodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetAllPNodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetallpNodeLogic(r.Context(), ctx)
		resp, err := l.AgencygetallpNode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetdeptallpersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetDeptAllPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetdeptallpersonLogic(r.Context(), ctx)
		resp, err := l.Agencygetdeptallperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetjobdeptHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetJobDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetjobdeptLogic(r.Context(), ctx)
		resp, err := l.Agencygetjobdept(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetjobsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetJobsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetjobsLogic(r.Context(), ctx)
		resp, err := l.Agencygetjobs(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetjobspageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetJobsPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetjobspageLogic(r.Context(), ctx)
		resp, err := l.Agencygetjobspage(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetnodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetNodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetnodeLogic(r.Context(), ctx)
		resp, err := l.Agencygetnode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetAgencyListByIdsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetAgencyListByIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetAgencyListByIdsLogic(r.Context(), ctx)
		resp, err := l.AgencygetAgencyListByIds(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetDeptCodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetDeptCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetDeptCodeLogic(r.Context(), ctx)
		resp, err := l.AgencygetDeptCode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetInnerAgencyByTenantCodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetInnerAgencyByTenantCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetInnerAgencyByTenantCodeLogic(r.Context(), ctx)
		resp, err := l.AgencygetInnerAgencyByTenantCode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencygetOrgDepartmentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyGetOrgDepartmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencygetOrgDepartmentLogic(r.Context(), ctx)
		resp, err := l.AgencygetOrgDepartment(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyimportdeptHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyImportDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyimportdeptLogic(r.Context(), ctx)
		resp, err := l.Agencyimportdept(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyimportinnerAgencyv2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyImportInnerAgencyV2Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyimportinnerAgencyv2Logic(r.Context(), ctx)
		resp, err := l.AgencyimportinnerAgencyv2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyinnertreeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyInnerTreeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyinnertreeLogic(r.Context(), ctx)
		resp, err := l.Agencyinnertree(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyinnerAgencymodelv2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyinnerAgencymodelv2Logic(r.Context(), ctx)
		resp, err := l.AgencyinnerAgencymodelv2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencylistinnerHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyListInnerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencylistinnerLogic(r.Context(), ctx)
		resp, err := l.Agencylistinner(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func Agencylistv2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyListV2Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencylistv2Logic(r.Context(), ctx)
		resp, err := l.Agencylistv2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func Agencylistv3Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyListV3Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencylistv3Logic(r.Context(), ctx)
		resp, err := l.Agencylistv3(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencymodeldowndeptHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencymodeldowndeptLogic(r.Context(), ctx)
		resp, err := l.Agencymodeldowndept(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencysearchDeptTreeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencySearchDeptTreeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencysearchDeptTreeLogic(r.Context(), ctx)
		resp, err := l.AgencysearchDeptTree(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyselectusersbyagencyIdOrJobIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencySelectUsersByAgencyIdOrJobIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyselectusersbyagencyIdOrJobIdLogic(r.Context(), ctx)
		resp, err := l.AgencyselectusersbyagencyIdOrJobId(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyselectAgencyIdByUserIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencySelectAgencyIdByUserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyselectAgencyIdByUserIdLogic(r.Context(), ctx)
		resp, err := l.AgencyselectAgencyIdByUserId(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyselectUserIdsByAgencyIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencySelectUserIdsByAgencyIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyselectUserIdsByAgencyIdLogic(r.Context(), ctx)
		resp, err := l.AgencyselectUserIdsByAgencyId(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyselectUsersByAgencyIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencySelectUsersByAgencyIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyselectUsersByAgencyIdLogic(r.Context(), ctx)
		resp, err := l.AgencyselectUsersByAgencyId(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyupdateinnernodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyUpdateInnerNodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyupdateinnernodeLogic(r.Context(), ctx)
		resp, err := l.Agencyupdateinnernode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyupdateinneruserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyUpdateInnerUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyupdateinneruserLogic(r.Context(), ctx)
		resp, err := l.Agencyupdateinneruser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AgencyupdateUsersToAgencyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AgencyUpdateUsersToAgencyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := inner.NewAgencyupdateUsersToAgencyLogic(r.Context(), ctx)
		resp, err := l.AgencyupdateUsersToAgency(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
