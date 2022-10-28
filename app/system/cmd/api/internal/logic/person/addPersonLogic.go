package person

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddPersonLogic {
	return AddPersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPersonLogic) AddPerson(req types.AddPersonReq) (resp *types.CommonResponse, err error) {
	if len(req.Phone) != 11 {
		return types.Failed(http.StatusBadRequest, errors.New("手机号长度不是11位!"))
	}
	roleId, err := strconv.ParseInt(req.Role, 10, 64)
	if err != nil {
		return types.Failed(http.StatusBadRequest, errors.New("角色ID无效!"))
	}
	deptIds := make([]int64, 0)
	if len(req.Depart) > 0 {
		for _, d := range req.Depart {
			deptId, err := strconv.ParseInt(d, 10, 64)
			if err == nil {
				deptIds = append(deptIds, deptId)
			}
		}
	}
	return types.JsonResult(l.svcCtx.SystemRpc.AddPerson(l.ctx, &system.AddPersonReq{
		Role:       roleId,
		TenantCode: req.TenantCode,
		Phone:      req.Phone,
		Name:       req.Name,
		UserCode:   req.UserCode,
		Depart:     deptIds,
	}))
}
