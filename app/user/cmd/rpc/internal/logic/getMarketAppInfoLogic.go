package logic

import (
	"context"
	"orginone/common/schema/asmarketappkeysecret"

	"orginone/app/user/cmd/rpc/internal/svc"
	"orginone/common/rpcs/user"
	"orginone/common/schema/asmarketapp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMarketAppInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMarketAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMarketAppInfoLogic {
	return &GetMarketAppInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取应用信息
func (l *GetMarketAppInfoLogic) GetMarketAppInfo(in *user.GetMarketAppInfoReq) (*user.CommonRpcRes, error) {
	return user.JsonResult(l.svcCtx.UserStore.AsMarketApp.Query().Where(asmarketapp.IsDeleted(0), asmarketapp.HasAppKeysWith(asmarketappkeysecret.AppKey(in.Appkey), asmarketappkeysecret.AppSecret(in.Appsecret))).All(l.ctx))
}
