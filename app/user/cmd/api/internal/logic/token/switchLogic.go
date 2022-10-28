package token

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/user"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwitchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SwitchLogic {
	return SwitchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchLogic) Switch(req types.SwitchReq) (resp *types.CommonResponse, err error) {
	if tools.IsNilOrEmpty(req.TenantCode) {
		_, _, code, err := common.GetTokenInfo(l.ctx)
		if err != nil {
			return types.Failed(http.StatusForbidden, err)
		}
		req.TenantCode = code
	}
	rpcres, err := l.svcCtx.UserRpc.SwitchTenantByCode(l.ctx, &user.SwitchTenantReq{
		Account:    req.Account,
		TenantCode: req.TenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var userEntity *schema.AsUser
	err = json.Unmarshal([]byte(rpcres.Json), &userEntity)
	if err != nil || userEntity == nil {
		return types.Failed(http.StatusNotFound, errors.New("switch tenant error."))
	}
	res, err := common.GenTokenRes(
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
		userEntity.ID, userEntity.PhoneNumber, userEntity.TenantCode)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	res["account"] = userEntity.PhoneNumber
	res["userName"] = userEntity.UserName
	roleId := make([]string, 0)
	for _, v := range userEntity.Edges.Roles {
		roleId = append(roleId, strconv.Itoa(int(v.ID)))
	}
	res["roleId"] = strings.Join(roleId, ",")
	return types.Successful(res)
}
