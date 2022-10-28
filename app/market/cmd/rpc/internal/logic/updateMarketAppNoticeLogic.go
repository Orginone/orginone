package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"orginone/app/market/cmd/rpc/internal/svc"
	"orginone/common/rpcs/market"
	"orginone/common/schema/asgrouptenantrelations"
	"orginone/common/tools"
	"orginone/common/tools/linq"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMarketAppNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMarketAppNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMarketAppNoticeLogic {
	return &UpdateMarketAppNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type unitId struct {
	Type        int64  `json:"type,optional,default=-1"`
	Id          string `json:"id,optional"`
	Name        string `json:"name,optional"`
	Check       int64  `json:"check,default=-1"`
	Label       string `json:"label,optional"`
	GroupId     string `json:"groupId,optional"`
	GroupOrUnit int64  `json:"groupOrUnit,optional,default=-1"`
}

// 新增或修改通知
func (l *UpdateMarketAppNoticeLogic) UpdateMarketAppNotice(in *market.UpdateMarketAppNoticeReq) (*market.CommonRpcRes, error) {
	var queryIdsStr string
	if !tools.IsNilOrEmpty(in.NoticeUnitIds) {
		uids := make([]unitId, 0)
		err := json.Unmarshal([]byte(in.NoticeUnitIds), &uids)
		if err == nil {
			queryIds := make([]int, 0)
			for _, item := range uids {
				if !tools.IsNilOrEmpty(item.Id) {
					id, err := strconv.Atoi(item.Id)
					if err != nil {
						continue
					}
					if item.Check == 2 {
						//关联直接下级
						queryIds, err = l.svcCtx.MarketStore.AsGroupTenantRelations.Query().
							Where(
								asgrouptenantrelations.ParentID(int64(id)),
								asgrouptenantrelations.StatusIn(102, 106),
								asgrouptenantrelations.ExpiresTimeIsNil(),
							).Select(asgrouptenantrelations.FieldSonID).Ints(l.ctx)
						if err != nil {
							continue
						}
					} else if item.Check == 3 {
						//关联全部下级
						groupEntity, err := l.svcCtx.MarketStore.AsAllGroup.Get(l.ctx, int64(id))
						if err != nil {
							continue
						}
						queryIds, err = l.svcCtx.MarketStore.AsGroupTenantRelations.Query().
							Where(
								asgrouptenantrelations.GroupCodeContains(groupEntity.GroupCode),
								asgrouptenantrelations.StatusIn(102, 106),
								asgrouptenantrelations.ExpiresTimeIsNil(),
							).Select(asgrouptenantrelations.FieldSonID).Ints(l.ctx)
						if err != nil {
							continue
						}
					}
					if item.GroupOrUnit != -1 && item.Check != -1 && tools.ArrIndexOf(queryIds, id) == -1 {
						queryIds = append(queryIds, id)
					}
				}
			}
			linq.From(queryIds).Distinct().ToSlice(&queryIds)
			bytes, err := json.Marshal(queryIds)
			if err == nil {
				queryIdsStr = string(bytes)
			}
		}
	}
	if in.Id == -1 {
		releaseUnitId, err := strconv.Atoi(in.NoticeReleaseUnitId)
		if err != nil {
			return &market.CommonRpcRes{}, err
		}
		return market.JsonResult(l.svcCtx.MarketStore.AsMarketAppNotice.Create().
			SetNoticeTitle(in.NoticeTitle).
			SetGroupOrUnit(in.GroupOrUnit).
			SetNoticeContent(in.NoticeContent).
			SetNoticeRoleIds(in.NoticeRoleIds).
			SetNoticeReleaseUnitID(int64(releaseUnitId)).
			SetNoticeUnitIds(in.NoticeUnitIds).
			SetNoticeReleaseStatus(in.NoticeReleaseStatus).
			SetUnitQueryIds(queryIdsStr).
			Save(l.ctx))
	} else {
		update := l.svcCtx.MarketStore.AsMarketAppNotice.UpdateOneID(in.Id)
		if !tools.IsNilOrEmpty(in.NoticeContent) {
			update.SetNoticeContent(in.NoticeContent)
		}
		if !tools.IsNilOrEmpty(in.NoticeTitle) {
			update.SetNoticeTitle(in.NoticeTitle)
		}
		if in.NoticeReleaseStatus != -1 {
			update.SetNoticeReleaseStatus(in.NoticeReleaseStatus)
		}
		return market.JsonResult(update.SetUnitQueryIds(queryIdsStr).Save(l.ctx))
	}
}
