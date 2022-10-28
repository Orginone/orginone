package jobs

import (
	"context"
	"strconv"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatejobuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatejobuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatejobuserLogic {
	return UpdatejobuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatejobuserLogic) Updatejobuser(req types.UpdateJobUserReq) (resp *types.CommonResponse, err error) {
	sort, err := strconv.ParseInt(req.Sort, 10, 64)
	if err != nil {
		return &types.CommonResponse{}, err
	}
	return types.BoolResult(l.svcCtx.SystemRpc.UpdateJobUser(l.ctx, &system.UpdateJobUserReq{
		JobId:     req.JobId,
		JobName:   req.JobName,
		PersonIds: tools.ArrayStrToInt64(strings.Split(req.UserIds, ",")),
		Sort:      sort,
	}))
}
