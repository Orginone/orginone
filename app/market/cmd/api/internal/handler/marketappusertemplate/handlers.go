package handler

import (
	"net/http"

	"orginone/app/market/cmd/api/internal/logic/marketappusertemplate"
	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDefaultTemplateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDefaultTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewGetDefaultTemplateLogic(r.Context(), ctx)
		resp, err := l.GetDefaultTemplate(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetTemplateIdByLoginInUseHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewGetTemplateIdByLoginInUseLogic(r.Context(), ctx)
		resp, err := l.GetTemplateIdByLoginInUse(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetTemplateIdByUserIdInUseHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTemplateIdByUserIdInUseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewGetTemplateIdByUserIdInUseLogic(r.Context(), ctx)
		resp, err := l.GetTemplateIdByUserIdInUse(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetTemplateIdListByLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewGetTemplateIdListByLoginLogic(r.Context(), ctx)
		resp, err := l.GetTemplateIdListByLogin(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetTemplateIdListByUserIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTemplateIdListByUserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewGetTemplateIdListByUserIdLogic(r.Context(), ctx)
		resp, err := l.GetTemplateIdListByUserId(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq12
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewListLogic(r.Context(), ctx)
		resp, err := l.List(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func RemoveHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveReq12
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewRemoveLogic(r.Context(), ctx)
		resp, err := l.Remove(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SetDefaultTemplateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetDefaultTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewSetDefaultTemplateLogic(r.Context(), ctx)
		resp, err := l.SetDefaultTemplate(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SetGroupTenantDefaultTemplateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetGroupTenantDefaultTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewSetGroupTenantDefaultTemplateLogic(r.Context(), ctx)
		resp, err := l.SetGroupTenantDefaultTemplate(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SubmitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitReq12
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappusertemplate.NewSubmitLogic(r.Context(), ctx)
		resp, err := l.Submit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
