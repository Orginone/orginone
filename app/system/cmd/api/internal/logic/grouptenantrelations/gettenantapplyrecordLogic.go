package grouptenantrelations

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GettenantapplyrecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGettenantapplyrecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) GettenantapplyrecordLogic {
	return GettenantapplyrecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GettenantapplyrecordLogic) Gettenantapplyrecord(req types.GetTenantApplyRecordReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	statusMap := map[int64][]int64{
		105: {101, 102, 103, 104},
		107: {102, 103, 104},
		101: {101},
		102: {102},
		103: {103},
		104: {104},
	}
	rpcres, err := l.svcCtx.SystemRpc.FindAllGroupByTentantCode(l.ctx, &system.FindAllGroupByTenantCodeReq{
		Page: &system.PageRequest{
			Offset: req.Size * (req.Current - 1),
			Limit:  req.Size,
			Filter: req.LikeName,
		},
		RelationTypeRange: []int64{2},
		StatusRange:       statusMap[req.Flag],
		TenantCode:        tenantCode,
		Type:              -1,
	})
	return types.PageResult(rpcres, err)
}
