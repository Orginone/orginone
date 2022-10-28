package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asallgroup"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGroupLogic {
	return &QueryGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询集团
func (l *QueryGroupLogic) QueryGroup(in *system.QueryGroupReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0))
	if !tools.IsNilOrEmpty(in.GroupName) {
		query = query.Where(asallgroup.GroupNameContains(in.GroupName))
	}
	if in.GroupType > 0 {
		query = query.Where(asallgroup.Type(in.GroupType))
	}
	groupEntitys, err := query.All(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	groupModels := make([]*model.AllGroupRecord, 0)
	for _, g := range groupEntitys {
		group := new(model.AllGroupRecord)
		group.AsAllGroup = g
		groupModels = append(groupModels, group)
	}
	return system.JsonResult(groupModels, err)
}
