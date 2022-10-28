package person

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/app/system/model"
	"orginone/common"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetuserinfoLogic {
	return GetuserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserinfoLogic) Getuserinfo(req types.GetUserInfo) (resp *types.CommonResponse, err error) {
	// 查询人员相关
	userId, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusForbidden, err)
	}
	rpcres, err := l.svcCtx.SystemRpc.FindPersonByUserId(l.ctx, &system.PersonByUserReq{
		UserId:     userId,
		TenantCode: tenantCode,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var persons []*model.PersonInfo
	err = json.Unmarshal([]byte(rpcres.Json), &persons)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	if len(persons) > 0 {
		return types.JsonResult(persons[0], err)
	}
	return types.Failed(http.StatusNotFound, err)
}
