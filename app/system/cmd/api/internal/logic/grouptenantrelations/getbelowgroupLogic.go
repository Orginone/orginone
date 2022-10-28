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

type GetbelowgroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetbelowgroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetbelowgroupLogic {
	return GetbelowgroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetbelowgroupLogic) Getbelowgroup(req types.GetBelowGroupReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindGroupLayerById(l.ctx, &system.GroupLayerByIdReq{GroupId: req.GroupId, SourceGroupId: req.GroupId, WithUnit: false})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var infos []*model.GroupUnitInfo
	json.Unmarshal([]byte(rpcres.Json), &infos)
	return types.JsonResult(infos, err)
}
