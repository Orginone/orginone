package allgroup

import (
	"context"
	"net/http"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreategroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreategroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreategroupLogic {
	return CreategroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreategroupLogic) Creategroup(req types.CreateGroupReq) (resp *types.CommonResponse, err error) {
	userId, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	if tools.IsNilOrEmpty(req.TenantCode) {
		req.TenantCode = tenantCode
	}
	shapes := tools.ArrayStrToInt64(strings.Split(req.Shape, ","))
	return types.CommonResult(l.svcCtx.SystemRpc.CreateGroup(l.ctx, &system.CreateGroupReq{
		UserId:           userId,
		Shape:            shapes,
		GroupId:          req.GroupId,
		GroupCode:        req.GroupCode,
		GroupName:        req.GroupName,
		TenantCode:       req.TenantCode,
		GroupDescription: req.GroupDescription,
	}))
}
