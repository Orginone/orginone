package scenes

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerytenantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerytenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) QuerytenantLogic {
	return QuerytenantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerytenantLogic) Querytenant(req types.QueryTenantReq) (resp *types.CommonResponse, err error) {
	if !tools.IsNilOrEmpty(req.PhoneNumber) {
		_, account, _, err := common.GetTokenInfo(l.ctx)
		if err != nil {
			return types.Failed(http.StatusForbidden, err)
		}
		req.PhoneNumber = account
	}
	return types.PageResult(l.svcCtx.SystemRpc.FindTenantInfoByAccount(l.ctx, &system.AccountTenantReq{
		Account: req.PhoneNumber,
		Count:   req.Count,
		Page: &system.PageRequest{
			Filter: req.TenantName,
			Limit:  req.Size,
			Offset: (req.Current - 1) * req.Size,
		},
	}))
}
