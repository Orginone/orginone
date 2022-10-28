package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveDeptJobUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveDeptJobUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveDeptJobUserLogic {
	return &RemoveDeptJobUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 移除内部机构人员
func (l *RemoveDeptJobUserLogic) RemoveDeptJobUser(in *system.RemoveDeptUserJobReq) (*system.CommonRpcRes, error) {
	if in.DeptId != -1 {
		return system.JsonResult(l.svcCtx.SystemStore.AsInnerAgency.UpdateOneID(in.DeptId).
			RemovePersonIDs(in.PersonIds...).Save(l.ctx))
	}
	if in.JobId != -1 {
		return system.JsonResult(l.svcCtx.SystemStore.AsJob.UpdateOneID(in.JobId).
			RemovePersonIDs(in.PersonIds...).Save(l.ctx))
	}
	return &system.CommonRpcRes{}, nil
}
