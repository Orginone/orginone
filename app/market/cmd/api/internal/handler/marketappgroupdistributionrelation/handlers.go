package handler

import (
	"net/http"

	"orginone/app/market/cmd/api/internal/logic/marketappgroupdistributionrelation"
	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRelationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRelationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := marketappgroupdistributionrelation.NewGetRelationLogic(r.Context(), ctx)
		resp, err := l.GetRelation(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
