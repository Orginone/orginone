package grouptenantrelations

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

type GrouptreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrouptreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrouptreeLogic {
	return GrouptreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrouptreeLogic) Grouptree(req types.GroupTreeReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.GetGroupTree(l.ctx, &system.GroupTreeReq{
		GroupId:       req.GroupId,
		SourceGroupId: req.GroupId,
	})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	groupTrees := make([]*model.GroupUnitInfo, 0)
	err = json.Unmarshal([]byte(rpcres.Json), &groupTrees)
	return types.JsonResult(groupTrees, err)
}
