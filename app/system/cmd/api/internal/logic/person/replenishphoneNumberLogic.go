package person

import (
	"context"
	"errors"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplenishphoneNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplenishphoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReplenishphoneNumberLogic {
	return ReplenishphoneNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplenishphoneNumberLogic) ReplenishphoneNumber(req types.ReplenishPhoneNumberReq) (resp *types.CommonResponse, err error) {
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusBadRequest, errors.New("请输入正确的电话号码!"))
	}
	return types.JsonResult(l.svcCtx.SystemRpc.UpdatePersonPhoneNumber(l.ctx, &system.UpdatePersonPhoneNumberReq{
		PhoneNumber: req.PhoneNumber,
		TenantCode:  tenantCode,
		PersonId:    req.PersonId,
	}))
}
