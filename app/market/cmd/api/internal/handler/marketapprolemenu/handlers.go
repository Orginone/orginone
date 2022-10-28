package handler

import (
	"net/http"

	"orginone/app/market/cmd/api/internal/logic/marketapprolemenu"
	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailReq7
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketapprolemenu.NewDetailLogic(r.Context(), ctx)
		resp, err := l.Detail(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq11
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketapprolemenu.NewListLogic(r.Context(), ctx)
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
		var req types.RemoveReq11
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketapprolemenu.NewRemoveLogic(r.Context(), ctx)
		resp, err := l.Remove(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func SubmitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitReq11
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketapprolemenu.NewSubmitLogic(r.Context(), ctx)
		resp, err := l.Submit(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
