package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/baseinfoadministrativeareaall"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAdministrativeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAdministrativeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAdministrativeListLogic {
	return &FindAdministrativeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取行政区域
func (l *FindAdministrativeListLogic) FindAdministrativeList(in *system.PageRequest) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.Baseinfoadministrativeareaall.Query().Where(baseinfoadministrativeareaall.IsDeleted(0))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	adminEntitys, err := query.Order(schema.Asc(baseinfoadministrativeareaall.FieldID)).Offset(int(in.Offset)).Limit(int(in.Limit)).All(l.ctx)
	return system.PageResult(in, total, adminEntitys, err)
}
