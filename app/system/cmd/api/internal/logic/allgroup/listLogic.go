package allgroup

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.QueryGroup(l.ctx, &system.QueryGroupReq{
		GroupName: req.GroupName,
		GroupType: req.Type,
	})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	groupModels := make([]*model.AllGroupRecord, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &groupModels)
	return types.JsonResult(groupModels, err)
}
