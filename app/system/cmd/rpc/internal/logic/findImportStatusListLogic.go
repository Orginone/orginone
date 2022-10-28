package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asinputdata"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindImportStatusListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindImportStatusListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindImportStatusListLogic {
	return &FindImportStatusListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取导入状态详情
func (l *FindImportStatusListLogic) FindImportStatusList(in *system.FindImportStatusListReq) (*system.CommonRpcRes, error) {
	query := l.svcCtx.SystemStore.AsInputData.Query().Where(asinputdata.FileNameContains(in.FileName), asinputdata.IsDeleted(0))
	if !tools.IsNilOrEmpty(in.TableName) {
		query.Where(asinputdata.TableName(in.TableName))
	}
	if !tools.IsNilOrEmpty(in.TenantCode) {
		query.Where(asinputdata.TenantCode(in.TenantCode))
	}
	if in.Type != -1 {
		query.Where(asinputdata.Type(in.Type))
	}
	if in.Status != -1 {
		query.Where(asinputdata.Status(in.Status))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	inputDataEntitys, err := query.Order(schema.Desc(asinputdata.FieldCreateTime)).Offset(int(in.Page.Offset)).Limit(int(in.Page.Limit)).All(l.ctx)
	return system.PageResult(in.Page, int64(count), inputDataEntitys, err)
}
