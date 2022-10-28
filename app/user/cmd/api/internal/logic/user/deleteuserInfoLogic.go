package user

import (
	"context"
	"net/http"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/user"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteuserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteuserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteuserInfoLogic {
	return DeleteuserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteuserInfoLogic) DeleteuserInfo(req []string) (resp *types.CommonResponse, err error) {
	userId, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusRequestTimeout, err)
	}
	deleteUserIds := tools.StrArrayToNumArray(req)
	index := tools.ArrIndexOf(deleteUserIds, userId)
	if index != -1 {
		deleteUserIds = append(deleteUserIds[:index], deleteUserIds[index+1:]...)
	}
	return types.CommonResult(l.svcCtx.UserRpc.DeleteUserInfo(l.ctx, &user.DeleteUserInfoReq{
		UserIds:    deleteUserIds,
		TenantCode: tenantCode,
	}))
}
