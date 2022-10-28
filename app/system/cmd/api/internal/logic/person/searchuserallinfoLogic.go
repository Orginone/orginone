package person

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/models"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchuserallinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchuserallinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchuserallinfoLogic {
	return SearchuserallinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchuserallinfoLogic) Searchuserallinfo(req types.SearchUserAllinfoReq) (resp *types.CommonResponse, err error) {
	rpcres, err := l.svcCtx.SystemRpc.FindTenantPersons(l.ctx, &system.SearchPersonByTenantCodeReq{
		Page: &system.PageRequest{
			Offset: int64(req.Size * (req.Current - 1)),
			Limit:  int64(req.Size),
			Filter: "",
		},
		TenantCode: req.TenantCode,
		IsActivate: 0,
		DeptId:     req.DeptId,
		JobId:      req.JobId,
	})
	if req.Version != "v2" {
		return types.PageResult(rpcres, err)
	}
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	var data *models.PageResponse
	err = json.Unmarshal([]byte(rpcres.Json), &data)
	if data != nil {
		return types.JsonResult(data.Records, err)
	}
	emptys := make([]string, 0)
	return types.JsonResult(emptys, err)
}
