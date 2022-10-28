package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2listPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2listPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2listPageLogic {
	return V2listPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2listPageLogic) V2listPage(req types.V2ListPageReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindOrgansV2ListPage(l.ctx, &system.FindOrgansV2ListPageReq{
		Page: &system.PageRequest{
			Limit:       req.Size,
			Offset:      (req.Current - 1) * req.Size,
			SearchCount: true,
		},
		UnitName:         req.UnitName,
		SocialCreditCode: req.SocialCreditCode,
		TenantId:         req.TenantId,
		IsDeleted:        req.IsDeleted,
		Id:               req.Id,
	}))
}
