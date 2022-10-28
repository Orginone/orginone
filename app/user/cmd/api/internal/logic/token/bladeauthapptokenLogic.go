package token

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/user"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type BladeauthapptokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBladeauthapptokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) BladeauthapptokenLogic {
	return BladeauthapptokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BladeauthapptokenLogic) Bladeauthapptoken(req types.AppTokenReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.UserRpc.GetMarketAppInfo(l.ctx, &user.GetMarketAppInfoReq{
		Appkey:    req.Appkey,
		Appsecret: req.Appsecret,
	})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	var entitys []*schema.AsMarketApp
	err = json.Unmarshal([]byte(rpcres.Json), &entitys)
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	if len(entitys) <= 0 {
		return types.Failed(http.StatusNotFound, errors.New("应用服务异常"))
	}

	return types.JsonResult(common.GenAppTokenRes(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire,
		entitys[0].ID, entitys[0].AppName, req.Appkey, req.Appsecret))
}
