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

type GetgroupbyidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetgroupbyidLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetgroupbyidLogic {
	return GetgroupbyidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetgroupbyidLogic) Getgroupbyid(req types.GetGroupByIdReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.GetAllgroupInfoById(l.ctx, &system.PrimaryKeyReq{Id: req.GroupId})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	modelGroup := new(model.AllGroupRecord)
	err = json.Unmarshal([]byte(rpcres.Json), &modelGroup)
	return types.JsonResult(modelGroup, err)
}
