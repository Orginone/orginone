package user

import (
	"context"
	"errors"
	"net/http"

	"orginone/app/user/cmd/api/internal/svc"
	"orginone/app/user/cmd/api/internal/types"
	"orginone/common/rpcs/user"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatepasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatepasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatepasswordLogic {
	return UpdatepasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatepasswordLogic) Updatepassword(req types.UpdatePasswordReq) (resp *types.CommonResponse, err error) {
	if req.NewPassword1 != req.NewPassword {
		return types.Failed(http.StatusBadRequest, errors.New("两次密码输入不一致."))
	}
	if tools.CheckPwd(req.NewPassword) {
		return types.Failed(http.StatusBadRequest, errors.New("新密码不符合规范.(必须包含:大写字母、小写字母、数字)"))
	}
	return types.CommonResult(l.svcCtx.UserRpc.UpdatePasswd(l.ctx, &user.UpdatePasswdReq{
		UserId: req.UserId,
		NewPwd: req.NewPassword,
		OldPwd: req.OldPassword,
	}))
}
