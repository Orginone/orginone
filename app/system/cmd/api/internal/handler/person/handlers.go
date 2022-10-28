package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"orginone/app/system/cmd/api/internal/logic/person"
	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ActivateuserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.ActivateUserReq
		// if err := httpx.Parse(r, &req); err != nil {
		// 	httpx.OkJson(w, types.BadRequest(err))
		// 	return
		// }
		req := make([]string, 0)
		bytes, err := io.ReadAll(r.Body)
		json.Unmarshal(bytes, &req)
		l := person.NewActivateuserLogic(r.Context(), ctx)
		resp, err := l.Activateuser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewAddLogic(r.Context(), ctx)
		resp, err := l.Add(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func AddPersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewAddPersonLogic(r.Context(), ctx)
		resp, err := l.AddPerson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func BanuserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.BanUserReq
		// if err := httpx.Parse(r, &req); err != nil {
		// 	httpx.OkJson(w, types.BadRequest(err))
		// 	return
		// }
		req := make([]string, 0)
		bytes, err := io.ReadAll(r.Body)
		json.Unmarshal(bytes, &req)
		l := person.NewBanuserLogic(r.Context(), ctx)
		resp, err := l.Banuser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func DeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewDeleteLogic(r.Context(), ctx)
		resp, err := l.Delete(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ExportExcelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExportExcelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewExportExcelLogic(r.Context(), ctx)
		resp, err := l.ExportExcel(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetactivedaymonthHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetActiveDayMonthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetactivedaymonthLogic(r.Context(), ctx)
		resp, err := l.Getactivedaymonth(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetallconspersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllConsPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetallconspersonLogic(r.Context(), ctx)
		resp, err := l.Getallconsperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func Getallconspersonv2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllConsPersonV2Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetallconspersonv2Logic(r.Context(), ctx)
		resp, err := l.Getallconspersonv2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetallpersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetallpersonLogic(r.Context(), ctx)
		resp, err := l.Getallperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetallpersonindeptHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPersonInDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetallpersonindeptLogic(r.Context(), ctx)
		resp, err := l.Getallpersonindept(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetallpersonintenantHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPersonInTenantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetallpersonintenantLogic(r.Context(), ctx)
		resp, err := l.Getallpersonintenant(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetallpersonsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPersonsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetallpersonsLogic(r.Context(), ctx)
		resp, err := l.Getallpersons(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetapprolepersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAppRolePersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetapprolepersonLogic(r.Context(), ctx)
		resp, err := l.Getapproleperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetinitpersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInitPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetinitpersonLogic(r.Context(), ctx)
		resp, err := l.Getinitperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetpersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetpersonLogic(r.Context(), ctx)
		resp, err := l.Getperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetpersonbyidHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPersonByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetpersonbyidLogic(r.Context(), ctx)
		resp, err := l.Getpersonbyid(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetpersonbylikeNameHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPersonByLikeNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetpersonbylikeNameLogic(r.Context(), ctx)
		resp, err := l.GetpersonbylikeName(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetpersonbypersonidHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPersonByPersonIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetpersonbypersonidLogic(r.Context(), ctx)
		resp, err := l.Getpersonbypersonid(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetpersontenantCodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetpersontenantCodeLogic(r.Context(), ctx)
		resp, err := l.GetpersontenantCode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GettenantintopersonsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTenantIntoPersonsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGettenantintopersonsLogic(r.Context(), ctx)
		resp, err := l.Gettenantintopersons(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetuserdataHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetuserdataLogic(r.Context(), ctx)
		resp, err := l.Getuserdata(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func Getuserdata2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetuserdata2Logic(r.Context(), ctx)
		resp, err := l.Getuserdata2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetuserinfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetuserinfoLogic(r.Context(), ctx)
		resp, err := l.Getuserinfo(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetuserlogintimeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetuserlogintimeLogic(r.Context(), ctx)
		resp, err := l.Getuserlogintime(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetAllPersonInTenantAHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPersonInTenantReq1
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetAllPersonInTenantALogic(r.Context(), ctx)
		resp, err := l.GetAllPersonInTenantA(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetAllPersonInTenantSimpleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPersonInTenantSimpleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetAllPersonInTenantSimpleLogic(r.Context(), ctx)
		resp, err := l.GetAllPersonInTenantSimple(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetMasterUnitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMasterUnitReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetMasterUnitLogic(r.Context(), ctx)
		resp, err := l.GetMasterUnit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetOrgEmployeeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrgEmployeeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetOrgEmployeeLogic(r.Context(), ctx)
		resp, err := l.GetOrgEmployee(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetOrgsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrgsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetOrgsLogic(r.Context(), ctx)
		resp, err := l.GetOrgs(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetPersonListByIdsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPersonListByIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetPersonListByIdsLogic(r.Context(), ctx)
		resp, err := l.GetPersonListByIds(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetUserCodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetUserCodeLogic(r.Context(), ctx)
		resp, err := l.GetUserCode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetUserInfoByMobileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoByMobileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewGetUserInfoByMobileLogic(r.Context(), ctx)
		resp, err := l.GetUserInfoByMobile(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportadminHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImportadminLogic(r.Context(), ctx)
		resp, err := l.Importadmin(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportreplenishadminHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportReplenishAdminReq1
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImportreplenishadminLogic(r.Context(), ctx)
		resp, err := l.Importreplenishadmin(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportreplenishphonenumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportReplenishPhoneNumberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImportreplenishphonenumberLogic(r.Context(), ctx)
		resp, err := l.Importreplenishphonenumber(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImporttenantpersonuserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportTenantPersonUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImporttenantpersonuserLogic(r.Context(), ctx)
		resp, err := l.Importtenantpersonuser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportunittenantHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportUnitTenantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImportunittenantLogic(r.Context(), ctx)
		resp, err := l.Importunittenant(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportExcelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportExcelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImportExcelLogic(r.Context(), ctx)
		resp, err := l.ImportExcel(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportExcelv2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportExcelV2Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImportExcelv2Logic(r.Context(), ctx)
		resp, err := l.ImportExcelv2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportExcel2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportExcel2Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewImportExcel2Logic(r.Context(), ctx)
		resp, err := l.ImportExcel2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ListPageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListPageReq3
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewListPageLogic(r.Context(), ctx)
		resp, err := l.ListPage(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func MergeUnitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MergeUnitReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewMergeUnitLogic(r.Context(), ctx)
		resp, err := l.MergeUnit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ModeldownHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewModeldownLogic(r.Context(), ctx)
		resp, err := l.Modeldown(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ModeldownpersonuserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewModeldownpersonuserLogic(r.Context(), ctx)
		resp, err := l.Modeldownpersonuser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ModeldownreplenishadminHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewModeldownreplenishadminLogic(r.Context(), ctx)
		resp, err := l.Modeldownreplenishadmin(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ModeldownreplenishphoneNumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewModeldownreplenishphoneNumberLogic(r.Context(), ctx)
		resp, err := l.ModeldownreplenishphoneNumber(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ModeldownunittenantHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewModeldownunittenantLogic(r.Context(), ctx)
		resp, err := l.Modeldownunittenant(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func PersonjoinstaffHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PersonJoinStaffReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewPersonjoinstaffLogic(r.Context(), ctx)
		resp, err := l.Personjoinstaff(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func PersonleavestaffHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PersonLeaveStaffReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewPersonleavestaffLogic(r.Context(), ctx)
		resp, err := l.Personleavestaff(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ProcessMessageUserIdsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProcessMessageUserIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewProcessMessageUserIdsLogic(r.Context(), ctx)
		resp, err := l.ProcessMessageUserIds(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RemoveHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveReq2
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewRemoveLogic(r.Context(), ctx)
		resp, err := l.Remove(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RemoveuserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewRemoveuserLogic(r.Context(), ctx)
		resp, err := l.Removeuser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ReplenishphoneNumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplenishPhoneNumberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewReplenishphoneNumberLogic(r.Context(), ctx)
		resp, err := l.ReplenishphoneNumber(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ReplenishUserCodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplenishUserCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewReplenishUserCodeLogic(r.Context(), ctx)
		resp, err := l.ReplenishUserCode(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SearchdeptpersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchDeptPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSearchdeptpersonLogic(r.Context(), ctx)
		resp, err := l.Searchdeptperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SearchuserallinfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchUserAllinfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSearchuserallinfoLogic(r.Context(), ctx)
		resp, err := l.Searchuserallinfo(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SearchuserinfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSearchuserinfoLogic(r.Context(), ctx)
		resp, err := l.Searchuserinfo(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SearchPersonListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchPersonListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSearchPersonListLogic(r.Context(), ctx)
		resp, err := l.SearchPersonList(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SetMasterUnitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetMasterUnitReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSetMasterUnitLogic(r.Context(), ctx)
		resp, err := l.SetMasterUnit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SinglelistHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SingleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSinglelistLogic(r.Context(), ctx)
		resp, err := l.Singlelist(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SinglelistpageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SingleListPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSinglelistpageLogic(r.Context(), ctx)
		resp, err := l.Singlelistpage(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SinglesynchronizeallpersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSinglesynchronizeallpersonLogic(r.Context(), ctx)
		resp, err := l.Singlesynchronizeallperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SinglesynchronizepersonbyphonenumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SingleSynchronizePersonByPhoneNumberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSinglesynchronizepersonbyphonenumberLogic(r.Context(), ctx)
		resp, err := l.Singlesynchronizepersonbyphonenumber(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SubmitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *schema.AsPerson
		tools.UnmarshaJsonIOData(r.Body, &req)

		l := person.NewSubmitLogic(r.Context(), ctx)
		resp, err := l.Submit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SynchronizeTenantUserNameHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SynchronizeTenantUserNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewSynchronizeTenantUserNameLogic(r.Context(), ctx)
		resp, err := l.SynchronizeTenantUserName(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func TenantpersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TenantPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewTenantpersonLogic(r.Context(), ctx)
		resp, err := l.Tenantperson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateUserGetImgHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserGetImgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateUserGetImgLogic(r.Context(), ctx)
		resp, err := l.UpdateUserGetImg(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func TokenisExistPhoneNumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TokenIsExistPhoneNumberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewTokenisExistPhoneNumberLogic(r.Context(), ctx)
		resp, err := l.TokenisExistPhoneNumber(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateLogic(r.Context(), ctx)
		resp, err := l.Update(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateuserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateuserLogic(r.Context(), ctx)
		resp, err := l.Updateuser(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateuseraddressHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateuseraddressLogic(r.Context(), ctx)
		resp, err := l.Updateuseraddress(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateuseremailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateuseremailLogic(r.Context(), ctx)
		resp, err := l.Updateuseremail(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateuserphoneNumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserPhoneNumberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateuserphoneNumberLogic(r.Context(), ctx)
		resp, err := l.UpdateuserphoneNumber(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateuserphotoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserPhotoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateuserphotoLogic(r.Context(), ctx)
		resp, err := l.Updateuserphoto(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdateuserrealNameHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserRealNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdateuserrealNameLogic(r.Context(), ctx)
		resp, err := l.UpdateuserrealName(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func Updatev1Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReq1
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdatev1Logic(r.Context(), ctx)
		resp, err := l.Updatev1(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func Updatev2Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateV2Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := person.NewUpdatev2Logic(r.Context(), ctx)
		resp, err := l.Updatev2(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
