package scenes

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatetenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatetenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreatetenantLogic {
	return CreatetenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatetenantLogic) Createtenant(req types.CreateTenantReq) (resp *types.CommonResponse, err error) {
	// 查询人员相关
	userId, _, _, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	rpcres, err := l.svcCtx.SystemRpc.CreateTenant(l.ctx, &system.CreateTenantReq{
		TenantName:    req.TenantName,
		SocialCreCode: req.SocialCreCode,
		UserId:        userId,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(rpcres.Json)
}
