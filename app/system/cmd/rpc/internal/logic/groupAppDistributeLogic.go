package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/models"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asappgroupdistributiondata"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/schema/asmarketappgroupdistribution"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketroledistribution"
	"orginone/common/schema/astenant"
	"orginone/common/schema/asunit"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupAppDistributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupAppDistributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupAppDistributeLogic {
	return &GroupAppDistributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 集团应用分发
func (l *GroupAppDistributeLogic) GroupAppDistribute(in *system.GroupAppDistributeReq) (*system.CommonRpcRes, error) {
	codes, err := l.getTenantCodesByConfig(in)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	err = l.deleteDistributeNotInCodes(in, codes)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	localCodes, err := l.svcCtx.SystemStore.AsMarketAppGroupDistribution.Query().Where(asmarketappgroupdistribution.IsDeleted(0),
		asmarketappgroupdistribution.AppID(in.AppId), asmarketappgroupdistribution.GroupID(in.GroupId)).Select(asmarketappgroupdistribution.FieldTenantID).Strings(l.ctx)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	codes = tools.ArrCodesAddOrDelete(codes, localCodes, 2)
	if len(codes) > 0 {
		creates := make([]*schema.AsMarketAppGroupDistributionCreate, 0)
		for _, code := range codes {
			creates = append(creates, l.svcCtx.SystemStore.AsMarketAppGroupDistribution.Create().SetAppID(in.AppId).
				SetGroupID(in.GroupId).SetTenantID(code).SetCreateUser(in.UserId).SetUseStatus(1))
		}
		_, err = l.svcCtx.SystemStore.AsMarketAppGroupDistribution.CreateBulk(creates...).Save(l.ctx)
	}
	err = l.saveConfigList(in)
	return &system.CommonRpcRes{}, err
}

// 删除不包含的分配记录
func (l *GroupAppDistributeLogic) deleteDistributeNotInCodes(in *system.GroupAppDistributeReq, codes []string) error {
	_, err := l.svcCtx.SystemStore.AsMarketAppGroupDistribution.Update().Where(asmarketappgroupdistribution.IsDeleted(0),
		asmarketappgroupdistribution.AppID(in.AppId), asmarketappgroupdistribution.GroupID(in.GroupId),
		asmarketappgroupdistribution.TenantIDNotIn(codes...)).SetIsDeleted(0).Save(l.ctx)
	if err == nil {
		_, err = l.svcCtx.SystemStore.AsMarketRoleDistribution.Update().Where(asmarketroledistribution.IsDeleted(0),
			asmarketroledistribution.TenantCodeNotIn(codes...), asmarketroledistribution.HasRolexWith(asmarketapprole.AppID(in.AppId))).SetIsDeleted(0).Save(l.ctx)
	}
	return err
}

//获取租户编码根据分发请求数据
func (l *GroupAppDistributeLogic) getTenantCodesByConfig(in *system.GroupAppDistributeReq) ([]string, error) {
	codes := make([]string, 0)
	addModel, delModel, err := l.parseConfigModel(in)
	if err != nil {
		return codes, err
	}
	codes, err = l.mergeTenantCodes(codes, addModel, in.GroupId, 1)
	if err != nil {
		return codes, err
	}
	codes, err = l.mergeTenantCodes(codes, delModel, in.GroupId, 2)
	if err != nil {
		return codes, err
	}
	return codes, nil
}

// 合并处理查找到的tentcodes
func (l *GroupAppDistributeLogic) mergeTenantCodes(srcArr []string, confModel *models.DistributionConfigModel, groupId int64, opt int64) ([]string, error) {
	disTypes := strings.Split(confModel.DistributeType, ",")
	for _, v := range disTypes {
		if v == "1" || v == "2" || v == "3" {
			codes, err := l.getTenantCodesByGroupId(groupId, v)
			if err != nil {
				return codes, err
			}
			srcArr = tools.ArrCodesAddOrDelete(srcArr, codes, opt)
			if v == "1" || v == "3" {
				return srcArr, err
			}
		}
	}
	for _, c := range confModel.ConfigList {
		switch c.Type {
		case 1: //集团
			code, err := l.getTentcodeByGroupId(c.Id)
			if err != nil {
				return srcArr, err
			}
			srcArr = tools.ArrCodesAddOrDelete(srcArr, []string{code}, opt)
			searchType := ""
			if c.Check == 2 {
				searchType = "3"
			} else if c.Check == 3 {
				searchType = "4"
			}
			if searchType != "" {
				codes, err := l.getTenantCodesByGroupId(c.Id, "3")
				if err != nil {
					return srcArr, err
				}
				srcArr = tools.ArrCodesAddOrDelete(srcArr, codes, opt)
			}
			break
		case 2: //单位
			code, err := l.GetTentcodeByUnitId(c.Id)
			if err != nil {
				return srcArr, err
			}
			srcArr = tools.ArrCodesAddOrDelete(srcArr, []string{code}, opt)
			break
		}
	}
	return srcArr, nil
}

func (l *GroupAppDistributeLogic) parseConfigModel(in *system.GroupAppDistributeReq) (addModel *models.DistributionConfigModel, delModel *models.DistributionConfigModel, err error) {
	configList := make([]*models.DistributionConfigModel, 0)
	err = json.Unmarshal([]byte(in.DistribeJson), &configList)
	if err != nil || len(configList) != 2 {
		return
	}
	for i, conf := range configList {
		if conf.Contain == 1 {
			addModel = configList[i]
		} else {
			delModel = configList[i]
		}
	}
	if addModel == nil || delModel == nil {
		err = errors.New("request error")
	}
	return
}

// 根据集团Id获取租户编号
// searchType 1-全部，2-主管单位，3-基层单位, 4-全部下属单位
func (l *GroupAppDistributeLogic) getTenantCodesByGroupId(groupId int64, searchType string) ([]string, error) {
	query := l.svcCtx.SystemStore.AsGroupTenantRelations.Query().Where(asgrouptenantrelations.IsDeleted(0),
		asgrouptenantrelations.Type(2), asgrouptenantrelations.StatusIn(102, 106), asgrouptenantrelations.ExpiresTimeIsNil())
	switch searchType { //1-全部，2-主管单位，3-基层单位
	case "1":
		groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, groupId)
		if err != nil || groupEntity == nil {
			return nil, errors.New("group not find.")
		}
		query = query.Where(asgrouptenantrelations.GroupCodeHasPrefix(groupEntity.GroupCode))
		codes, err := query.Clone().QueryTenant().Where(astenant.IsDeleted(0)).Select(astenant.FieldTenantCode).Strings(l.ctx)
		if err != nil {
			return codes, err
		}
		adminCodes, err := query.QueryGroup().Where(asallgroup.IsDeleted(0)).Select(asallgroup.FieldTenantCode).Strings(l.ctx)
		if err != nil {
			return codes, err
		}
		return tools.ArrCodesAddOrDelete(codes, adminCodes, 1), err
	case "2":
		groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, groupId)
		if err != nil || groupEntity == nil {
			return nil, errors.New("group not find.")
		}
		return query.Where(asgrouptenantrelations.GroupCodeHasPrefix(groupEntity.GroupCode)).QueryGroup().
			Where(asallgroup.IsDeleted(0)).Select(asallgroup.FieldTenantCode).Strings(l.ctx)
	case "3":
		return query.Where(asgrouptenantrelations.ParentID(groupId)).QueryTenant().
			Where(astenant.IsDeleted(0)).Select(astenant.FieldTenantCode).Strings(l.ctx)
	case "4":
		groupEntity, err := l.svcCtx.SystemStore.AsAllGroup.Get(l.ctx, groupId)
		if err != nil || groupEntity == nil {
			return nil, errors.New("group not find.")
		}
		return query.Where(asgrouptenantrelations.GroupCodeHasPrefix(groupEntity.GroupCode),
			asgrouptenantrelations.GroupCodeNEQ(groupEntity.GroupCode)).QueryTenant().
			Where(astenant.IsDeleted(0)).Select(astenant.FieldTenantCode).Strings(l.ctx)
	default:
		return nil, errors.New("type is not supper.")
	}
}

// 根据集团ID获取管理租户编号
func (l *GroupAppDistributeLogic) getTentcodeByGroupId(groupId int64) (string, error) {
	return l.svcCtx.SystemStore.AsAllGroup.Query().Where(asallgroup.IsDeleted(0), asallgroup.ID(groupId)).
		Select(astenant.FieldTenantCode).String(l.ctx)
}

// 根据单位ID获取租户编号
func (l *GroupAppDistributeLogic) GetTentcodeByUnitId(unitId int64) (string, error) {
	return l.svcCtx.SystemStore.AsUnit.Query().Where(asunit.ID(unitId), asunit.IsDeleted(0)).
		QueryTenantx().Where(astenant.IsDeleted(0)).Select(astenant.FieldTenantCode).String(l.ctx)
}

// 缓存前端配置
func (l *GroupAppDistributeLogic) saveConfigList(in *system.GroupAppDistributeReq) error {
	entitys, err := l.svcCtx.SystemStore.AsAppGroupDistributionData.Query().Where(asappgroupdistributiondata.AppID(in.AppId),
		asappgroupdistributiondata.GroupID(in.GroupId), asappgroupdistributiondata.IsDeleted(0)).Limit(1).All(l.ctx)
	if err != nil {
		return err
	}
	if len(entitys) > 0 {
		_, err = l.svcCtx.SystemStore.AsAppGroupDistributionData.UpdateOne(entitys[0]).SetConfig(in.DistribeJson).Save(l.ctx)
		if err != nil {
			return err
		}
	} else {
		_, err = l.svcCtx.SystemStore.AsAppGroupDistributionData.Create().SetAppID(in.AppId).SetGroupID(in.GroupId).SetConfig(in.DistribeJson).Save(l.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
