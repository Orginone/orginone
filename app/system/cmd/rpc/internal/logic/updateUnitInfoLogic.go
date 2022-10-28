package logic

import (
	"context"
	"encoding/json"
	"errors"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asunit"
	"orginone/common/schema/baseinfoadministrativeareaall"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUnitInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUnitInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUnitInfoLogic {
	return &UpdateUnitInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新单位信息
func (l *UpdateUnitInfoLogic) UpdateUnitInfo(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	var unitEntity *schema.AsUnit
	err := json.Unmarshal([]byte(in.Json), &unitEntity)
	if err != nil {
		return system.Result("false", err)
	}
	local, err := l.svcCtx.SystemStore.AsUnit.Get(l.ctx, unitEntity.ID)
	if err != nil {
		return system.Result("false", err)
	}
	if local == nil {
		return system.Result("false", errors.New("更新的单位不存在！"))
	}
	if local.UnitName != unitEntity.UnitName {
		isExist, err := l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.UnitName(unitEntity.UnitName)).Exist(l.ctx)
		if err != nil {
			return system.Result("false", err)
		}
		if isExist {
			return system.Result("false", errors.New("单位名称已存在!"))
		}
	}
	if local.SocialCreditCode != unitEntity.SocialCreditCode && !tools.IsNilOrEmpty(local.SocialCreditCode) {
		if unitEntity.VirtualUnit == 0 && !tools.SocialCreditCodeValidator(unitEntity.SocialCreditCode) {
			return system.Result("false", errors.New("实体单位统一社会信用代码不合法!"))
		}
		if isSCcodeExist, err := l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.SocialCreditCode(unitEntity.SocialCreditCode)).Exist(l.ctx); isSCcodeExist {
			if err != nil {
				return system.Result("false", err)
			}
			return system.Result("false", errors.New("单位统一社会信用代码重复!"))
		}
	}
	if tools.IsNilOrEmpty(unitEntity.SocialCreditCode) {
		unitEntity.SocialCreditCode = tools.CreateSocialCreditCode()
	}
	//查询邮编
	if !tools.IsNilOrEmpty(unitEntity.AdministrationDivisionName) {
		areaEntitys, err := l.svcCtx.SystemStore.Baseinfoadministrativeareaall.Query().
			Where(baseinfoadministrativeareaall.AllName(unitEntity.AdministrationDivisionName),
				baseinfoadministrativeareaall.IsDeleted(0)).Limit(1).All(l.ctx)
		if err == nil && len(areaEntitys) > 0 && len(areaEntitys[0].Code) >= 6 {
			unitEntity.ChargeSectionCode = string(areaEntitys[0].Code[0:6])
			unitEntity.AdministrationDivisionName = areaEntitys[0].AllName
		}
	}
	//开启事务
	tx, err := l.svcCtx.SystemStore.Tx(l.ctx)
	if err != nil {
		return system.Result("false", err)
	}
	updater := tx.AsUnit.UpdateOne(local)
	tools.SchemaUpdateAny(updater.Mutation(), unitEntity, "id")
	_, err = updater.Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return system.Result("false", err)
	}
	//更新组织
	_, err = tx.AsTenant.UpdateOneID(unitEntity.TenantID).
		SetIsVirtual(unitEntity.VirtualUnit).
		SetTenantName(unitEntity.UnitName).
		Save(l.ctx)
	if err != nil {
		tx.Rollback()
		return system.Result("false", err)
	}
	err = tx.Commit()
	return system.Result("true", err)
}
