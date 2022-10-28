package dict

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetalldictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetalldictLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetalldictLogic {
	return GetalldictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetalldictLogic) Getalldict(req types.DicGetalldictReq) (resp *types.CommonResponse, err error) {
	return types.PageResult(l.svcCtx.SystemRpc.FindAllDic(l.ctx, &system.FindAllDicReq{
		Code:    req.Code,
		DictKey: req.DictKey,
		Page: &system.PageRequest{
			Offset: ((req.Current - 1) * req.Size),
			Limit:  req.Size,
		},
	}))
}
