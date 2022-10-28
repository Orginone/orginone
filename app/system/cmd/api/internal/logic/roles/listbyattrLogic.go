package roles

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListbyattrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListbyattrLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListbyattrLogic {
	return ListbyattrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListbyattrLogic) Listbyattr(req types.ListByAttrReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.ListRoleByAttr(l.ctx, &system.ListRoleByAttrReq{
		AttrId: req.AttrId,
	})
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	var roles *[]schema.AsRole
	err = json.Unmarshal([]byte(rpcres.Json), &roles)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.JsonResult(roles, err)
}
