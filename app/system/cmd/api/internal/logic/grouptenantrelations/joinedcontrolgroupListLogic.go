package grouptenantrelations

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinedcontrolgroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinedcontrolgroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) JoinedcontrolgroupListLogic {
	return JoinedcontrolgroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinedcontrolgroupListLogic) JoinedcontrolgroupList(req types.JoinedControlGroupListReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindTenantById(l.ctx, &system.TenantByIdReq{Id: req.TenantId})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var asTenantEntity *schema.AsTenant
	err = json.Unmarshal([]byte(rpcres.Json), &asTenantEntity)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	if asTenantEntity == nil {
		return types.Failed(http.StatusInternalServerError, errors.New("tenant is not found."))
	}
	rpcres, err = l.svcCtx.SystemRpc.FindAllGroupByTentantCode(l.ctx, &system.FindAllGroupByTenantCodeReq{
		Page: &system.PageRequest{
			Offset: req.Size * (req.Current - 1),
			Limit:  req.Size,
			Filter: req.GroupName,
		},
		RelationTypeRange: []int64{1, 2},
		StatusRange:       []int64{102, 106},
		TenantCode:        asTenantEntity.TenantCode,
		Type:              req.Type,
	})
	return types.PageResult(rpcres, err)
}
