package propertiesdistribution

import (
	"context"
	"encoding/json"
	"net/http"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetpropertiesunitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetpropertiesunitLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetpropertiesunitLogic {
	return GetpropertiesunitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetpropertiesunitLogic) Getpropertiesunit(req types.GetPropertiesUnitReq) (resp *types.CommonResponse, err error) {
	rpcRes, err := l.svcCtx.SystemRpc.FindPropertiesUnitList(l.ctx, &system.FindPropertiesUnitListReq{
		PropertiesId: req.PropertiesId,
	})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	data := make([]map[string]interface{}, 0)
	err = json.Unmarshal([]byte(rpcRes.Json), &data)
	return types.JsonResult(data, err)
}
