package jobs

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req types.RemoveReq1) (resp *types.CommonResponse, err error) {
	for _, id := range tools.StrArrayToNumArray(strings.Split(req.Ids, ",")) {
		_, err := l.svcCtx.SystemRpc.RemoveJob(l.ctx, &system.RemoveJobReq{
			JobId: id,
		})
		if err != nil {
			return types.Failed(http.StatusInternalServerError, errors.New("删除中止!"+err.Error()))
		}
	}
	return types.BoolResult(true, nil)
}
