package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asproperties"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPropertiesListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPropertiesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPropertiesListLogic {
	return &FindPropertiesListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询性质
func (l *FindPropertiesListLogic) FindPropertiesList(in *system.FindPropertiesListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsProperties.Query().Where(asproperties.IsDeleted(0), asproperties.PropertiesNameContains(in.Page.Filter), asproperties.GroupID(in.GroupId))
	total, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	propertiesEntitys, err := query.Order(schema.Desc(asproperties.FieldCreateTime)).Limit(int(in.Page.Limit)).Offset(int(in.Page.Offset)).All(l.ctx)
	return system.PageResult(in.Page, total, propertiesEntitys, err)
}
