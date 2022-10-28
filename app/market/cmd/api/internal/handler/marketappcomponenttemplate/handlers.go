package handler

import (
	"net/http"

	"orginone/app/market/cmd/api/internal/logic/marketappcomponenttemplate"
	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetComponentByTemplateIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetComponentByTemplateIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappcomponenttemplate.NewGetComponentByTemplateIdLogic(r.Context(), ctx)
		resp, err := l.GetComponentByTemplateId(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func GetUserTemplateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappcomponenttemplate.NewGetUserTemplateLogic(r.Context(), ctx)
		resp, err := l.GetUserTemplate(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq5
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappcomponenttemplate.NewListLogic(r.Context(), ctx)
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
		var req types.RemoveReq4
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappcomponenttemplate.NewRemoveLogic(r.Context(), ctx)
		resp, err := l.Remove(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SetDefaultHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetDefaultReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappcomponenttemplate.NewSetDefaultLogic(r.Context(), ctx)
		resp, err := l.SetDefault(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SubmitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitReq4
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappcomponenttemplate.NewSubmitLogic(r.Context(), ctx)
		resp, err := l.Submit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
