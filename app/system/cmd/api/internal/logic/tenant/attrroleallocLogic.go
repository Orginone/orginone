package tenant

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttrroleallocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttrroleallocLogic(ctx context.Context, svcCtx *svc.ServiceContext) AttrroleallocLogic {
	return AttrroleallocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttrroleallocLogic) Attrrolealloc(req types.AttrRoleAllocReq) (resp *types.CommonResponse, err error) {
	attrId, err := strconv.ParseInt(req.AttrId, 10, 64)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	_, err = l.svcCtx.SystemRpc.TenantAttrRoleAlloc(l.ctx, &system.TenantAttrRoleAllocReq{
		AttrId:  attrId,
		RoleIds: tools.ArrayStrToInt64(strings.Split(req.RoleIds, ",")),
	})
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	return types.Successful(true)
}
