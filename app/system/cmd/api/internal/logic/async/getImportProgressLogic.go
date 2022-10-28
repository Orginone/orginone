package async

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImportProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetImportProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetImportProgressLogic {
	return GetImportProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetImportProgressLogic) GetImportProgress(req types.GetImportProgressReq) (resp *types.CommonResponse, err error) {
	type ProgressResult struct {
		Msg    string `json:"msg"`
		Status bool   `json:"status"`
		Value  int64  `json:"value"`
	}
	return types.JsonResult(&ProgressResult{
		Msg:    "成功",
		Status: true,
		Value:  200,
	}, nil)
}
