package tenant

import (
	"context"
	"net/http"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttrremoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttrremoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) AttrremoveLogic {
	return AttrremoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttrremoveLogic) Attrremove(req types.AttrRemoveReq) (resp *types.CommonResponse, err error) {
	_, err = l.svcCtx.SystemRpc.TenantAttrRemove(l.ctx, &system.TenantAttrRemoveReq{
		Id: tools.ArrayStrToInt64(strings.Split(req.Id, ",")),
	})
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	return types.Successful(true)
}
