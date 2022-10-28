package logic

import (
	"context"
	"encoding/json"
	"strings"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/tools"
	"orginone/common/tools/date"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitMarketAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitMarketAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitMarketAppLogic {
	return &SubmitMarketAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 提交应用市场应用注册
func (l *SubmitMarketAppLogic) SubmitMarketApp(in *market.CommonRpcReq) (*market.CommonRpcRes, error) {
	in.Json = strings.ReplaceAll(in.Json, `"status":""`, `"status":0`)
	appEntity := new(schema.AsMarketApp)
	err := json.Unmarshal([]byte(in.Json), appEntity)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	creater := l.svcCtx.MarketStore.AsMarketApp.Create()
	tools.SchemaUpdateAny(creater.Mutation(), appEntity, "id", "status")
	creater.SetStatus(0).SetApplyTime(date.Now())
	_, err = creater.Save(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	return market.Result("true", err)
}
