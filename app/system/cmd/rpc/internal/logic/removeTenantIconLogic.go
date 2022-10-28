package logic

import (
	"context"
	"strings"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema/astenanticon"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveTenantIconLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveTenantIconLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveTenantIconLogic {
	return &RemoveTenantIconLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 租户图片批量导入
func (l *RemoveTenantIconLogic) RemoveTenantIcon(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	pre := astenanticon.IDIn(tools.ArrayStrToInt64(strings.Split(in.Json, ","))...)
	urls, err := l.svcCtx.SystemStore.AsTenantIcon.Query().Where(pre).Select(astenanticon.FieldIcon).Strings(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	_, err = l.svcCtx.SystemStore.AsTenantIcon.Update().Where(pre).SetIsDeleted(1).Save(l.ctx)
	return system.JsonResult(urls, err)
}
