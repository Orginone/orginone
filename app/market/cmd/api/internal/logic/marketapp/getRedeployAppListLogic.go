package marketapp

import (
	"context"
	"errors"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedeployAppListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRedeployAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRedeployAppListLogic {
	return GetRedeployAppListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRedeployAppListLogic) GetRedeployAppList(req types.GetRedeployAppListReq) (resp *types.CommonResponse, err error) {
	status := make([]int64, 0)
	switch req.Flag {
	case 1:
		status = append(status, 0, 1, 2)
		break
	case 2:
		status = append(status, 0)
		break
	case 3:
		status = append(status, 1, 2)
		break
	default:
		return types.Failed(http.StatusBadRequest, errors.New("flag error."))
	}
	return types.PageResult(l.svcCtx.MarketRpc.GetRedeployAppList(l.ctx, &market.GetRedeployReq{
		Status:  status,
		AppName: req.AppName,
		Page: &market.PageRequest{
			Offset: ((req.Current - 1) * req.Size),
			Limit:  req.Size,
		},
	}))
}
