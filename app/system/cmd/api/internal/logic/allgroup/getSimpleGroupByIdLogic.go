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

type GetSimpleGroupByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSimpleGroupByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSimpleGroupByIdLogic {
	return GetSimpleGroupByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSimpleGroupByIdLogic) GetSimpleGroupById(req types.GetSimpleGroupByIdReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FingGroupInfoById(l.ctx, &system.PrimaryKeyReq{Id: req.GroupId})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	var groupTenant *model.GroupUnitInfo
	err = json.Unmarshal([]byte(rpcres.Json), &groupTenant)
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	return types.JsonResult(groupTenant, err)
}
