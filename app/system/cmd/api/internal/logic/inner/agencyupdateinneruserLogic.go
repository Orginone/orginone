package inner

import (
	"context"
	"strconv"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgencyupdateinneruserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgencyupdateinneruserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AgencyupdateinneruserLogic {
	return AgencyupdateinneruserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgencyupdateinneruserLogic) Agencyupdateinneruser(req types.AgencyUpdateInnerUserReq) (resp *types.CommonResponse, err error) {
	deptId, err := strconv.ParseInt(req.DeptId, 10, 64)
	if err != nil {
		return types.BadRequest(err), err
	}
	var personIds []int64
	for _, i := range req.PersonIds {
		personId, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return types.BadRequest(err), err
		}
		personIds = append(personIds, personId)
	}
	return types.JsonResult(l.svcCtx.SystemRpc.UpdateAgencyInnerUser(l.ctx, &system.UpdateAgencyInnerUserReq{
		AgencyCode: req.AgencyCode,
		DeptId:     deptId,
		DeptName:   req.DeptName,
		PId:        req.PId,
		PersonIds:  personIds,
	}))
}
