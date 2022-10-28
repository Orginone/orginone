package token

import (
	"context"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type BladeauthappinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBladeauthappinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) BladeauthappinfoLogic {
	return BladeauthappinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BladeauthappinfoLogic) Bladeauthappinfo(req types.AppTokenReq) (resp *types.CommonResponse, err error) {
	appId, appName, err := common.GetAppTokenInfo(l.ctx)
	return types.JsonResult(map[string]interface{}{
		"appId":   appId,
		"appName": appName,
	}, nil)
}
