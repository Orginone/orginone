package logic

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePersonInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonInfoLogic {
	return &UpdatePersonInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新人员信息
func (l *UpdatePersonInfoLogic) UpdatePersonInfo(in *system.CommonRpcRes) (*system.CommonRpcRes, error) {
	var personEntity *schema.AsPerson
	err := json.Unmarshal([]byte(in.Json), &personEntity)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	local, err := l.svcCtx.SystemStore.AsPerson.Get(l.ctx, personEntity.ID)
	updater := l.svcCtx.SystemStore.AsPerson.UpdateOne(local)
	tools.SchemaUpdateAny(updater.Mutation(), personEntity, "id")
	return system.JsonResult(updater.Save(l.ctx))
}
