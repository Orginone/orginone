package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/asdict"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindDicItemsByValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindDicItemsByValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindDicItemsByValueLogic {
	return &FindDicItemsByValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 系统字典(blade-system)
func (l *FindDicItemsByValueLogic) FindDicItemsByValue(in *system.FindDicItemReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsDict.Query().Where(asdict.IsDeleted(0))
	if !tools.IsNilOrEmpty(in.DictValue) {
		query = query.Where(asdict.DictValue(in.DictValue))
	}
	query = query.QueryChildrens()
	if in.ParentId > -1 {
		query = query.Where(asdict.ParentID(in.ParentId))
	}
	if in.Version > -1 {
		query = query.Where(asdict.Version(in.Version))
	}
	return system.JsonResult(query.All(l.ctx))
}
