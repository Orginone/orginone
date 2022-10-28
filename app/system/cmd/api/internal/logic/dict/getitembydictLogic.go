package dict

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetitembydictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetitembydictLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetitembydictLogic {
	return GetitembydictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetitembydictLogic) Getitembydict(req types.DicGetitembydictReq) (resp *types.CommonResponse, err error) {
	dicEntitys := make([]*schema.AsDict, 0)
	rpcres, err := l.svcCtx.SystemRpc.FindDicItemsByValue(l.ctx, &system.FindDicItemReq{
		DictValue: "",
		ParentId:  req.Parentid,
		Version:   req.Version})
	if err != nil {
		return &types.CommonResponse{}, err
	}
	err = json.Unmarshal([]byte(rpcres.Json), &dicEntitys)
	return types.JsonResult(dicEntitys, err)
}
