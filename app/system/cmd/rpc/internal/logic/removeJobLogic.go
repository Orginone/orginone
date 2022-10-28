package logic

import (
	"context"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asjob"
	"orginone/common/schema/asworkingdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveJobLogic {
	return &RemoveJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除岗位
func (l *RemoveJobLogic) RemoveJob(in *system.RemoveJobReq) (*system.CommonRpcRes, error) {
	isExist, err := l.svcCtx.SystemStore.AsJob.Query().
		Where(asjob.ID(in.JobId), asjob.IsDeleted(0)).
		QueryPersons().QueryUserx().QueryWorkingDatas().
		Where(asworkingdata.Type(3), asworkingdata.IsDeleted(0)).Exist(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if isExist {
		return &system.CommonRpcRes{}, errors.New("存在在途业务人员，无法删除!")
	}
	return system.JsonResult(l.svcCtx.SystemStore.AsJob.UpdateOneID(in.JobId).
		ClearPersons().SetIsDeleted(1).Save(l.ctx))
}
