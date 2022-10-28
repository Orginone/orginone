package image

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/tools/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadImageLogic {
	return UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(req io.Reader, ext string) (resp *types.CommonResponse, err error) {
	fileName := fmt.Sprintf("%d%s", snowflake.SnowGen.NextVal(), ext)
	buff, err := ioutil.ReadAll(req)
	if err != nil {
		return
	}
	err = os.WriteFile(fmt.Sprintf("./Image/%s", fileName), buff, 0200)
	if err != nil {
		return
	}
	return &types.CommonResponse{
		Data: fmt.Sprintf("/dev-api/asset-system/image/downloadimage?fileName=%s", fileName),
		Msg:  "上传成功!",
	}, err
}
