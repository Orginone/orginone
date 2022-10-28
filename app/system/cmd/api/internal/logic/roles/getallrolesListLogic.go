package roles

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallrolesListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallrolesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallrolesListLogic {
	return GetallrolesListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallrolesListLogic) GetallrolesList(req types.GetAllRolesListReq) (resp *types.CommonResponse, err error) {
	if tools.IsNilOrEmpty(req.TenantCode) {
		_, _, req.TenantCode, err = common.GetTokenInfo(l.ctx)
		if err != nil {
			return types.Failed(http.StatusInternalServerError, err)
		}
	}
	rpcRes, err := l.svcCtx.SystemRpc.FindAllRolesList(l.ctx, &system.FindAllRolesListReq{
		TenantCode: req.TenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var roleEntitys []*schema.AsRole
	err = json.Unmarshal([]byte(rpcRes.Json), &roleEntitys)
	if err != nil {
		return types.Failed(http.StatusNotExtended, err)
	}
	return types.Successful(roleEntitys)
}
