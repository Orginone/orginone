package acttodo

import (
	"context"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnitActTodoNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnitActTodoNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUnitActTodoNumberLogic {
	return GetUnitActTodoNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUnitActTodoNumberLogic) GetUnitActTodoNumber(req types.GetUnitActTodoNumberReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	return types.JsonResult([]string{tenantCode}, err)
}
