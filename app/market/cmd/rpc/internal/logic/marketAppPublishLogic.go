package logic

import (
	"context"
	"encoding/json"
	"errors"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/app/market/model"
	"orginone/common/rpcs/market"
	"orginone/common/schema"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/tools"
	"orginone/common/tools/date"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketAppPublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarketAppPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketAppPublishLogic {
	return &MarketAppPublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用发布
func (l *MarketAppPublishLogic) MarketAppPublish(in *market.CommonRpcReq) (*market.CommonRpcRes, error) {
	publishModel := new(model.PublishAppModel)
	err := json.Unmarshal([]byte(in.Json), &publishModel)
	if err != nil || publishModel == nil {
		return &market.CommonRpcRes{}, errors.New("request error.")
	}
	tx, err := l.svcCtx.MarketStore.Tx(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	_, err = tx.AsMarketApp.UpdateOneID(publishModel.AppId).SetStatus(6).SetSaleStatus(0).
		SetPublishTime(date.Now()).SetApplyTime(date.Now()).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	roleList := make([]*schema.AsMarketAppRole, 0)
	_, err = tx.AsMarketAppRole.Update().Where(asmarketapprole.AppID(publishModel.AppId)).ClearRoleMenus().SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	for _, role := range publishModel.RoleList {
		roleEntity, err := tx.AsMarketAppRole.Create().SetAppID(publishModel.AppId).SetRoleName(role.RoleName).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
		roleList = append(roleList, roleEntity)
	}
	_, err = tx.AsMarketMenu.Update().Where(asmarketmenu.AppID(publishModel.AppId)).SetIsDeleted(1).Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return &market.CommonRpcRes{}, err
	}
	for _, menu := range publishModel.MenuList {
		err = l.createMenuOfMenuTree(menu, -1, roleList, publishModel.AppId, tx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
	}
	for _, component := range publishModel.ComponentList {
		if component.ID > 0 {
			err = tx.AsMarketAppComponent.DeleteOneID(component.ID).Exec(l.ctx)
			if err != nil {
				tx.Rollback()
				return &market.CommonRpcRes{}, err
			}
		}
		create := tx.AsMarketAppComponent.Create()
		tools.SchemaUpdateAny(create.Mutation(), component)
		_, err = create.SetAppID(publishModel.AppId).Save(l.ctx)
		if err != nil {
			tx.Rollback()
			return &market.CommonRpcRes{}, err
		}
	}
	err = tx.Commit()
	return &market.CommonRpcRes{Json: "publish success."}, err
}

//创建菜单及权限关系
func (l *MarketAppPublishLogic) createMenuOfMenuTree(m *model.MenuTreeModel, pId int64,
	roleList []*schema.AsMarketAppRole, appId int64, tx *schema.Tx) error {
	create := tx.AsMarketMenu.Create().SetAppID(appId).SetIcon(m.Icon).SetMenuURL(m.MenuUrl).
		SetMenuName(m.MenuName).SetOutIPMenuURL(m.InnerUrl).SetHTTPSMenuURL(m.HttpsMenuUrl)
	if pId > 0 {
		create.SetParentID(pId)
	}
	menuEntity, err := create.Save(l.ctx)
	if menuEntity != nil {
		for _, r := range m.RoleList {
			for _, s := range roleList {
				if r.RoleName == s.RoleName {
					if err = tx.AsMarketRoleMenu.Create().SetMenuID(menuEntity.ID).
						SetRoleID(s.ID).Exec(l.ctx); err != nil {
						return err
					}
				}
			}
		}
		for _, c := range m.Children {
			if err = l.createMenuOfMenuTree(c, menuEntity.ID, roleList, appId, tx); err != nil {
				return err
			}
		}
	}
	return err
}
