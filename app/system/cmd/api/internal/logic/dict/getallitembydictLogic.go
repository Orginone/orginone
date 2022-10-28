package dict

import (
	"context"
	"encoding/json"
	"errors"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetallitembydictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetallitembydictLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetallitembydictLogic {
	return GetallitembydictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetallitembydictLogic) Getallitembydict(req types.DicGetallitembydictReq) (resp []*schema.AsDict, err error) {
	dicEntitys := make([]*schema.AsDict, 0)
	if tools.IsNilOrEmpty(req.Filtertext) {
		return dicEntitys, errors.New("filtertext is null error.")
	}
	rpcres, err := l.svcCtx.SystemRpc.FindDicItemsByValue(l.ctx, &system.FindDicItemReq{
		DictValue: req.Filtertext,
		ParentId:  -1,
		Version:   -1})
	if err != nil {
		return dicEntitys, err
	}
	err = json.Unmarshal([]byte(rpcres.Json), &dicEntitys)
	return dicEntitys, err
}
