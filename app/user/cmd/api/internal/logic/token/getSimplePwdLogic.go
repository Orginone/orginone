package token

import (
	"context"
	"net/http"
	"orginone/common/tools"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSimplePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSimplePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSimplePwdLogic {
	return GetSimplePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSimplePwdLogic) GetSimplePwd(req types.GetSimplePwdReq) (resp *types.CommonResponse, err error) {
	pwd, err := tools.AesPwdDecrypt(req.Password, l.svcCtx.Config.AesSecretKey)
	if err != nil {
		return types.Failed(http.StatusExpectationFailed, err)
	}
	return types.Successful(tools.CheckPwd(pwd))
}
