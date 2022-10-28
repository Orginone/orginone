package grouptenantrelations

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetimportstatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetimportstatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetimportstatusLogic {
	return GetimportstatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetimportstatusLogic) Getimportstatus(req types.GetImportStatusReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindImportStatusList(l.ctx, &system.FindImportStatusListReq{
		Page: &system.PageRequest{
			Limit:       req.Size,
			Offset:      (req.Current - 1) * req.Size,
			SearchCount: true,
		},
		TenantCode: req.TenantCode,
		Status:     req.Status,
		FileName:   req.FileName,
		TableName:  req.TableName,
		Type:       req.Type,
	}))
}
