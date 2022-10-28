package user

import (
	"context"
	"net/http"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common"
	"orginone/common/rpcs/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTenantuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) TenantuserLogic {
	return TenantuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TenantuserLogic) Tenantuser(req types.TenantUserReq) (resp *types.CommonResponse, err error) {
	rangeState := make([]int64, 0)
	var userId int64 = -1
	switch req.Count {
	case 4:
		break
	case 5:
		rangeState = []int64{0, 2}
		break
	case 6:
		rangeState = []int64{2, 3}
		break
	case 7:
		rangeState = []int64{1, 2, 3}
		break
	default:
		rangeState = []int64{req.Count}
		break
	}
	_, _, tenantCode, err := common.GetTokenInfo(l.ctx)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.PageResult(l.svcCtx.UserRpc.FindUserTenantUser(l.ctx, &user.FindserTenantUserReq{
		Page: &user.PageRequest{
			Offset: (req.Current - 1) * req.Size,
			Limit:  req.Size,
		},
		Id:                      userId,
		UserName:                req.RealName,
		TenantCode:              tenantCode,
		TenantApplyinStateRange: rangeState,
	}))
}
