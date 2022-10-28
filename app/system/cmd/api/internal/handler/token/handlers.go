package handler

import (
	"net/http"

	"orginone/app/system/cmd/api/internal/logic/token"
	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegistrationretrieveidCardHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegistrationRetrieveIdCardReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := token.NewRegistrationretrieveidCardLogic(r.Context(), ctx)
		resp, err := l.RegistrationretrieveidCard(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}


