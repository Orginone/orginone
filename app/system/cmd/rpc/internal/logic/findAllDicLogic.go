package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asdict"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllDicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllDicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllDicLogic {
	return &FindAllDicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 字典分页查询
func (l *FindAllDicLogic) FindAllDic(in *system.FindAllDicReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsDict.Query().Where(asdict.IsDeleted(0), asdict.DictKey(in.DictKey))
	if tools.IsNilOrEmpty(in.Code) {
		query = query.Where(asdict.CodeContains(in.Code))
	}
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	dictEntitys, err := query.Order(schema.Asc(asdict.FieldSort)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, total, dictEntitys, err)
}
