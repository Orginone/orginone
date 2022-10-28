package token

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/user"
	"orginone/common/schema"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type BladeauthtokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBladeauthtokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) BladeauthtokenLogic {
	return BladeauthtokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BladeauthtokenLogic) Bladeauthtoken(req types.UserLoginReq) (resp *types.CommonResponse, err error) {
	plainPwd, err := tools.AesPwdDecrypt(req.Password, l.svcCtx.Config.AesSecretKey)
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	rpcres, err := l.svcCtx.UserRpc.FindUserByAccount(l.ctx, &user.UserByAccountReq{
		Id:       -1,
		Account:  req.Account,
		Password: plainPwd,
	})
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	var userEntitys []*schema.AsUser
	err = json.Unmarshal([]byte(rpcres.Json), &userEntitys)
	if err != nil {
		return types.Failed(http.StatusNotFound, err)
	}
	//获取之前切换的单位
	userEntity := linq.From(userEntitys).FirstTOrDefault(func(u *schema.AsUser) bool { return u.IsMaster == 1 }, userEntitys[0]).(*schema.AsUser)
	res, err := common.GenTokenRes(
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
		userEntity.ID, userEntity.PhoneNumber, userEntity.TenantCode)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	res["account"] = userEntity.PhoneNumber
	res["userName"] = userEntity.UserName
	return types.Successful(res)
}
