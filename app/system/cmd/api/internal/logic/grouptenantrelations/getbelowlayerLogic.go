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

type GetbelowlayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetbelowlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetbelowlayerLogic {
	return GetbelowlayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetbelowlayerLogic) Getbelowlayer(req types.GetBelowLayerReq) (resp *types.CommonResponse, err error) {
	if req.SourceGroupId <= 0 {
		req.SourceGroupId = req.GroupId
	}
	rpcres, err := l.svcCtx.SystemRpc.FindGroupLayerById(l.ctx, &system.GroupLayerByIdReq{GroupId: req.GroupId, SourceGroupId: req.SourceGroupId, WithUnit: true})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var infos []*model.GroupUnitInfo
	json.Unmarshal([]byte(rpcres.Json), &infos)
	return types.JsonResult(infos, err)
}
