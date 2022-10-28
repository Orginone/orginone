package logic

import (
	"context"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asrole"
	"orginone/common/schema/asuser"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoleIdListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoleIdListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoleIdListByUserIdLogic {
	return &FindRoleIdListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页角色Id
func (l *FindRoleIdListByUserIdLogic) FindRoleIdListByUserId(in *system.FindRoleListByUserIdReq) (*system.CommonRpcRes, error) {
	array, err := l.svcCtx.SystemStore.AsUser.Query().
		Where(asuser.ID(in.UserId), asuser.IsDeleted(0)).
		QueryRoles().Where(asrole.IsDeleted(0)).
		Order(schema.Asc(asrole.FieldSort)).
		Select(asrole.FieldID).Ints(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	newArray := make([]int, 0)
	for _, id := range array {
		if tools.ArrIndexOf(newArray, id) < 0 {
			newArray = append(newArray, id)
		}
	}
	length := len(newArray)
	data := make([]int, 0)
	if length > int(in.Page.Offset) {
		if length > int(in.Page.Offset)+int(in.Page.Limit) {
			data = newArray[in.Page.Offset : in.Page.Offset+in.Page.Limit]
		} else {
			data = newArray[in.Page.Offset:]
		}
	}
	return system.PageResult(in.Page, int64(length), data, err)
}
