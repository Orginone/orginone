package logic

import (
	"context"
	"encoding/json"
	"errors"
	"orginone/common/schema"
	"orginone/common/schema/asdict"
	"orginone/common/tools"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitDictLogic {
	return &SubmitDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 保存字典
func (l *SubmitDictLogic) SubmitDict(in *system.CommonRpcRes) (*system.CommonRpcRes, error) {
	var entity *schema.AsDict
	err := json.Unmarshal([]byte(in.Json), &entity)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	var num int64
	if entity.ID != 0 {
		num, err = l.svcCtx.SystemStore.AsDict.Query().Where(asdict.IsDeleted(0), asdict.IDNEQ(entity.ID), asdict.Code(entity.Code), asdict.DictKey(entity.DictKey)).Count(l.ctx)
	} else {
		num, err = l.svcCtx.SystemStore.AsDict.Query().Where(asdict.IsDeleted(0), asdict.Code(entity.Code), asdict.DictKey(entity.DictKey)).Count(l.ctx)
	}
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	if num > 0 {
		return &system.CommonRpcRes{}, errors.New("当前字典键值已存在!")
	}

	updater := l.svcCtx.SystemStore.AsDict.UpdateOneID(entity.ID)
	tools.SchemaUpdateAny(updater.Mutation(), entity, "id", "parent_id")
	obj, err := updater.Save(l.ctx)
	if obj == nil {
		create := l.svcCtx.SystemStore.AsDict.Create().SetParentID(entity.ParentID).SetCode(entity.Code).
			SetDictKey(entity.DictKey).SetDictValue(entity.DictValue).SetSort(entity.Sort).SetRemark(entity.Remark).
			SetCurrversion(entity.Currversion).SetVersion(entity.Version).SetDictparentID(entity.DictparentID)
		obj, err = create.Save(l.ctx)
	}
	return system.JsonResult(obj, err)
}
