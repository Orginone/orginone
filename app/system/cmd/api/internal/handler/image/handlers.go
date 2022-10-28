package handler

import (
	"errors"
	"io"
	"net/http"
	"orginone/app/system/cmd/api/internal/logic/image"
	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownLoadImageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExportImageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := image.NewDownLoadImageLogic(r.Context(), ctx)
		resp, err := l.DownLoadImage(req)
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			w.Write(resp)
		}
	}
}

func UploadImageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, head, err := r.FormFile("icon")
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}
		if file == nil {
			httpx.OkJson(w, types.BadRequest(errors.New("file is not upload.")))
			return
		}
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
			return
		}

		l := image.NewUploadImageLogic(r.Context(), ctx)
		resp, err := l.UploadImage(file.(io.Reader), path.Ext(head.Filename))
		if err != nil {
			httpx.OkJson(w, types.BadRequest(err))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
