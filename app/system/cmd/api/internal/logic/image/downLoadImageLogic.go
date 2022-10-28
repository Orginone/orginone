package image

import (
	"context"
	"os"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownLoadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownLoadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) DownLoadImageLogic {
	return DownLoadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownLoadImageLogic) DownLoadImage(req types.ExportImageReq) (resp []byte, err error) {
	return os.ReadFile("./Image/" + req.FileName)
}
