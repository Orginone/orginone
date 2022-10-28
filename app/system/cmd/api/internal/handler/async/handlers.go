package handler

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"orginone/app/system/cmd/api/internal/logic/async"
	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExportGroupUnitDataHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := make([]string, 0)
		if err := tools.UnmarshaJsonIOData(r.Body, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		}

		l := async.NewExportGroupUnitDataLogic(r.Context(), ctx)
		resp, err := l.ExportGroupUnitData(tools.ArrayStrToInt64(req))
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			w.Write(resp)
			// httpx.OkJson(w, resp)
		}
	}
}

func ExportPersonDataHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExportPersonDataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		l := async.NewExportPersonDataLogic(r.Context(), ctx)
		resp, err := l.ExportPersonData(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			w.Write(resp)
		}
	}
}

func GetImportProgressHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetImportProgressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewGetImportProgressLogic(r.Context(), ctx)
		resp, err := l.GetImportProgress(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GroupAppDistributeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupAppDistributeReq
		// if err := httpx.Parse(r, &req); err != nil {
		// 	httpx.OkJson(w, types.BadRequest(err))
		// 	return
		// }
		err := r.ParseForm()
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		appId := r.Form.Get("appId")
		req.AppId, err = strconv.ParseInt(string(appId), 10, 64)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		groupId := r.Form.Get("groupId")
		req.GroupId, err = strconv.ParseInt(string(groupId), 10, 64)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.JsonBody, err = ioutil.ReadAll(r.Body)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		l := async.NewGroupAppDistributeLogic(r.Context(), ctx)
		resp, err := l.GroupAppDistribute(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportDeptHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		req.File = file.(io.Reader)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.FileName = header.Filename
		l := async.NewImportDeptLogic(r.Context(), ctx)
		resp, err := l.ImportDept(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportGroupRelationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportGroupRelationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		req.File = file.(io.Reader)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.FileName = header.Filename
		l := async.NewImportGroupRelationLogic(r.Context(), ctx)
		resp, err := l.ImportGroupRelation(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportJobHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportJobReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		req.File = file.(io.Reader)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.FileName = header.Filename
		l := async.NewImportJobLogic(r.Context(), ctx)
		resp, err := l.ImportJob(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportMasterUnitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportMasterUnitReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		req.File = file.(io.Reader)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.FileName = header.Filename
		l := async.NewImportMasterUnitLogic(r.Context(), ctx)
		resp, err := l.ImportMasterUnit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportNoAdminUnitTenantHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportNoAdminUnitTenantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewImportNoAdminUnitTenantLogic(r.Context(), ctx)
		resp, err := l.ImportNoAdminUnitTenant(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportPersonHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportPersonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		req.File = file.(io.Reader)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.FileName = header.Filename
		l := async.NewImportPersonLogic(r.Context(), ctx)
		resp, err := l.ImportPerson(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportPhoneNumberHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportPhoneNumberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewImportPhoneNumberLogic(r.Context(), ctx)
		resp, err := l.ImportPhoneNumber(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportReplenishAdminHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportReplenishAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewImportReplenishAdminLogic(r.Context(), ctx)
		resp, err := l.ImportReplenishAdmin(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportTenantHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportTenantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		req.File = file.(io.Reader)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.FileName = header.Filename

		l := async.NewImportTenantLogic(r.Context(), ctx)
		resp, err := l.ImportTenant(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ImportTenantRelationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportTenantRelationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		req.File = file.(io.Reader)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		req.FileName = header.Filename

		l := async.NewImportTenantRelationLogic(r.Context(), ctx)
		resp, err := l.ImportTenantRelation(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func InfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewInfoLogic(r.Context(), ctx)
		resp, err := l.Info(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RealGroupMergeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RealGroupMergeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewRealGroupMergeLogic(r.Context(), ctx)
		resp, err := l.RealGroupMerge(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RemoveCacheHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveCacheReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewRemoveCacheLogic(r.Context(), ctx)
		resp, err := l.RemoveCache(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RemoveInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewRemoveInfoLogic(r.Context(), ctx)
		resp, err := l.RemoveInfo(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RemoveProgressHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveProgressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewRemoveProgressLogic(r.Context(), ctx)
		resp, err := l.RemoveProgress(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UpdategroupinformationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGroupInformationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := async.NewUpdategroupinformationLogic(r.Context(), ctx)
		resp, err := l.Updategroupinformation(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func DownloadImportTemplateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadImportTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := async.NewDownloadImportTemplateLogic(r.Context(), ctx)
		resp, err := l.DownloadImportTemplate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			w.Write(resp)
		}
	}
}
