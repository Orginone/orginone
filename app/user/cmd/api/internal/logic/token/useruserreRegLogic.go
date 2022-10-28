package token

import (
	"context"
	"net/http"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common/rpcs/user"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseruserreRegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseruserreRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) UseruserreRegLogic {
	return UseruserreRegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseruserreRegLogic) UseruserreReg(req types.UserOptReq) (resp *types.CommonResponse, err error) {
	plainPwd, err := tools.AesPwdDecrypt(req.Pwd, l.svcCtx.Config.AesSecretKey)
	if err != nil {
		return types.BadRequest(err), nil
	}
	_, err = l.svcCtx.UserRpc.UserReReg(l.ctx, &user.UserReRegReq{
		PhoneNumber: req.PhoneNumber,
		Pwd:         plainPwd,
		RealName:    req.RealName,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(true)
}
