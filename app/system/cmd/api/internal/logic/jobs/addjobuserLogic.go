package jobs

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddjobuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddjobuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddjobuserLogic {
	return AddjobuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddjobuserLogic) Addjobuser(req types.AddJobUserReq) (resp *types.CommonResponse, err error) {
	if req.JobName == "未分配人员" {
		return &types.CommonResponse{}, errors.New("不可创建名为未分配人员的岗位!")
	}
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.BoolResult(l.svcCtx.SystemRpc.AddJobUser(l.ctx, &system.AddJobUserReq{
		TenantCode: tenantCode,
		JobName:    req.JobName,
		PersonIds:  tools.ArrayStrToInt64(strings.Split(req.UserIds, ",")),
	}))
}
