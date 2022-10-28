package logic

import (
	"context"
	"encoding/json"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/asredeploydata"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckRedeployLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckRedeployLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckRedeployLogic {
	return &CheckRedeployLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过或拒绝应用
func (l *CheckRedeployLogic) CheckRedeploy(in *market.CheckRedeployReq) (*market.CommonRpcRes, error) {
	if in.Pass {
		deployEntitys, err := l.svcCtx.MarketStore.AsRedeployData.Query().
			Where(
				asredeploydata.IDIn(in.RedeployIds...),
				asredeploydata.HasAppxWith(asmarketapp.IsDeleted(0)),
			).
			WithAppx().
			All(l.ctx)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
		tx, err := l.svcCtx.MarketStore.Tx(l.ctx)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
		for _, deploy := range deployEntitys {
			appDetail := new(model.AppModelDetail)
			if json.Unmarshal([]byte(deploy.Application), &appDetail) != nil {
				break
			}

			//移除现有角色列表
			if _, err = tx.AsMarketAppRole.Update().Where(asmarketapprole.AppID(deploy.AppID)).SetIsDeleted(1).Save(l.ctx); err != nil {
				tx.Rollback()
				return &market.CommonRpcRes{}, err
			}
			roleCreateBulk := make([]*schema.AsMarketAppRoleCreate, 0)
			//新增角色列表
			for _, role := range appDetail.RoleList {
				creater := tx.AsMarketAppRole.Create()
				tools.SchemaUpdateAny(creater.Mutation(), role.AsMarketAppRole, "id")
				roleCreateBulk = append(roleCreateBulk, creater)
			}
			if tx.AsMarketAppRole.CreateBulk(roleCreateBulk...).Exec(l.ctx) != nil {
				tx.Rollback()
				return &market.CommonRpcRes{}, err
			}
			//删除以前的菜单
			if _, err := tx.AsMarketMenu.Update().Where(asmarketmenu.AppID(deploy.AppID)).ClearRoleMenus().SetIsDeleted(1).Save(l.ctx); err != nil {
				tx.Rollback()
				return &market.CommonRpcRes{}, err
			}
			//新增菜单
			if AddMenus(appDetail.Menu, tx, l.ctx) != nil {
				tx.Rollback()
				return &market.CommonRpcRes{}, err
			}
		}
		_, err = tx.AsRedeployData.Update().
			Where(asredeploydata.IDIn(in.RedeployIds...)).
			SetStatus(1).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
		return market.JsonResult("true", tx.Commit())
	} else {
		return market.JsonResult(l.svcCtx.MarketStore.AsRedeployData.
			Update().
			Where(asredeploydata.IDIn(in.RedeployIds...)).
			SetStatus(2).
			Save(l.ctx))
	}
}

func AddMenus(menuList []*model.AppMarketMenuModel, tx *schema.Tx, ctx context.Context) (err error) {
	//新增菜单
	for _, menu := range menuList {
		create := tx.AsMarketMenu.Create()
		tools.SchemaUpdateAny(create.Mutation(), menu.AsMarketMenu, "id")
		menuEntity, err := create.Save(ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, role := range menu.RoleList {
			_, err = tx.AsMarketRoleMenu.Create().SetRoleID(role.ID).SetMenuID(menuEntity.ID).Save(ctx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		if len(menu.Children) > 0 {
			err = AddMenus(menu.Children, tx, ctx)
		}
	}
	return
}
