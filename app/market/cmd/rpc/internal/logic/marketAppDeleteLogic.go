package logic

import (
	"context"
	"errors"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappalert"
	"orginone/common/schema/asmarketappcomponent"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/asmarketappkeysecret"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/asmarketroledistribution"
	"orginone/common/schema/asmarketusedapp"
	"orginone/common/schema/asrole"
	"orginone/common/schema/asuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketAppDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarketAppDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketAppDeleteLogic {
	return &MarketAppDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用删除
func (l *MarketAppDeleteLogic) MarketAppDelete(in *market.DeleteAppReq) (*market.CommonRpcRes, error) {
	removeAdmin, err := l.CheckRemoveAdmin(in)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	exist, err := l.svcCtx.MarketStore.AsMarketApp.Query().Where(asmarketapp.IsDeleted(0), asmarketapp.ID(in.AppId)).Exist(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	if !exist {
		removeAdmin = true
	}
	in.RemoveAdmin = removeAdmin
	return l.DeleteMarketApp(in)
}

// 判断应用是否存在订阅分发
func (l *MarketAppDeleteLogic) DeleteMarketApp(in *market.DeleteAppReq) (*market.CommonRpcRes, error) {
	//开启事务
	tx, err := l.svcCtx.MarketStore.Tx(l.ctx)
	if err != nil {
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketMenu.Update().Where(asmarketmenu.IsDeleted(0), asmarketmenu.AppID(in.AppId)).ClearRoleMenus().SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppRole.Update().Where(asmarketapprole.IsDeleted(0), asmarketapprole.AppID(in.AppId)).ClearRoleMenus().SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketUsedApp.Update().Where(asmarketusedapp.IsDeleted(0), asmarketusedapp.AppID(in.AppId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppKeySecret.Update().Where(asmarketappkeysecret.IsDeleted(0), asmarketappkeysecret.AppID(in.AppId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppComponent.Update().Where(asmarketappcomponent.IsDeleted(0), asmarketappcomponent.AppID(in.AppId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketAppAlert.Update().Where(asmarketappalert.IsDeleted(0), asmarketappalert.AlertReleaseAppID(in.AppId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	if in.RemoveAdmin {
		_, err = tx.AsMarketAppPurchase.Update().Where(asmarketapppurchase.IsDeleted(0), asmarketapppurchase.AppID(in.AppId)).SetIsDeleted(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
		_, err = tx.AsMarketAppGroupDistribution.Update().Where(asmarketappgroupdistribution.IsDeleted(0), asmarketappgroupdistribution.AppID(in.AppId)).SetIsDeleted(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
		_, err = tx.AsMarketRoleDistribution.Update().Where(asmarketroledistribution.IsDeleted(0), asmarketroledistribution.HasRolexWith(asmarketapprole.AppID(in.AppId))).SetIsDeleted(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
	}
	_, err = tx.AsMarketApp.Update().Where(asmarketapp.IsDeleted(0), asmarketapp.ID(in.AppId)).SetIsDeleted(1).Save(l.ctx)
	err = tx.Commit()
	return &market.CommonRpcRes{Json: "delete app is success."}, err
}

// 判断应用是否存在订阅分发
func (l *MarketAppDeleteLogic) CheckRemoveAdmin(in *market.DeleteAppReq) (bool, error) {
	distributeCount, err := l.svcCtx.MarketStore.AsMarketAppPurchase.Query().Where(asmarketapppurchase.AppID(in.AppId), asmarketapppurchase.IsDeleted(0)).Count(l.ctx)
	if err != nil {
		return false, err
	}
	if distributeCount == 0 {
		distributeCount, err = l.svcCtx.MarketStore.AsMarketAppGroupDistribution.Query().Where(asmarketappgroupdistribution.IsDeleted(0),
			asmarketappgroupdistribution.AppID(in.AppId)).Count(l.ctx)
		if err != nil {
			return false, err
		}
		if distributeCount == 0 {
			distributeCount, err = l.svcCtx.MarketStore.AsMarketRoleDistribution.Query().Where(asmarketroledistribution.IsDeleted(0),
				asmarketroledistribution.HasRolexWith(asmarketapprole.AppID(in.AppId), asmarketapprole.IsDeleted(0))).Count(l.ctx)
		}
	}
	if distributeCount > 0 {
		if in.RemoveAdmin {
			admin, err := l.svcCtx.MarketStore.AsUser.Query().Where(asuser.IsDeleted(0), asuser.HasRolesWith(asrole.ID(1))).Exist(l.ctx)
			if err != nil {
				return false, err
			}
			if !admin {
				in.RemoveAdmin = false
			}
		}
		if !in.RemoveAdmin {
			return in.RemoveAdmin, errors.New("app is used, do not delete it.")
		}
	}
	return in.RemoveAdmin, err
}
