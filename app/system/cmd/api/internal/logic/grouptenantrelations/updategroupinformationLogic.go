package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdategroupinformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdategroupinformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdategroupinformationLogic {
	return UpdategroupinformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdategroupinformationLogic) Updategroupinformation(req types.UpdateGroupInformationReq) (resp *types.CommonResponse, err error) {
	return types.CommonResult(l.svcCtx.SystemRpc.UpdateGroupInfo(l.ctx, &system.UpdateGroupInfoReq{
		GroupName:        req.GroupName,
		TenantCode:       req.TenantCode,
		SonGroupId:       req.SonGroupId,
		ParentGroupId:    req.ParentGroupId,
		GroupDescription: req.GroupDescription,
		SocialCreditCode: req.SocialCreditCode,
	}))
}
