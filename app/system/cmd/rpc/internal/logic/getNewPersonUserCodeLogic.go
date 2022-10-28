package logic

import (
	"context"
	"fmt"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewPersonUserCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNewPersonUserCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewPersonUserCodeLogic {
	return &GetNewPersonUserCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询全新人员编号
func (l *GetNewPersonUserCodeLogic) GetNewPersonUserCode(in *system.TenantCodeReq) (*system.CommonRpcRes, error) {
	code, index, err := l.svcCtx.SystemStore.GetNewUserCodeInfo(l.ctx, in.TenantCode)
	return system.Result(fmt.Sprintf("%s%06d", code, index), err)
}
