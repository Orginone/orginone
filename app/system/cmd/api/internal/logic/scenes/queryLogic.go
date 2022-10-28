package scenes

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryLogic {
	return QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req types.QueryReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindTenantById(l.ctx, &system.TenantByIdReq{Id: req.Id})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var tenantEntity schema.AsTenant
	err = json.Unmarshal([]byte(rpcres.Json), &tenantEntity)
	if err != nil {
		return types.Failed(http.StatusNotExtended, err)
	}
	if tenantEntity.Edges.Unit != nil {
		return types.Successful(model.UnitModel{
			AsUnit:     tenantEntity.Edges.Unit,
			TenantType: tenantEntity.TenantType,
			TenantCode: tenantEntity.TenantCode,
		})
	}
	return types.Failed(http.StatusNotFound, nil)
}
