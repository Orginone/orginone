package logic

import (
	"context"
	"encoding/json"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//更新用户信息
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.CommonRpcRes) (*user.CommonRpcRes, error) {
	var userEntity *schema.AsUser
	err := json.Unmarshal([]byte(in.Json), &userEntity)
	if err != nil {
		return &user.CommonRpcRes{}, err
	}
	local, err := l.svcCtx.UserStore.AsUser.Get(l.ctx, userEntity.ID)
	updater := l.svcCtx.UserStore.AsUser.UpdateOne(local)
	tools.SchemaUpdateAny(updater.Mutation(), userEntity, "id")
	return user.JsonResult(updater.Save(l.ctx))
}
