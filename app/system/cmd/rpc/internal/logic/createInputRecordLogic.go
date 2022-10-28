package logic

import (
	"context"
	"encoding/json"

	"orginone/app/system/cmd/rpc/internal/svc"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInputRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateInputRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInputRecordLogic {
	return &CreateInputRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 插入文件导入记录
func (l *CreateInputRecordLogic) CreateInputRecord(in *system.CommonRpcReq) (*system.CommonRpcRes, error) {
	inputData := new(schema.AsInputData)
	err := json.Unmarshal([]byte(in.Json), &inputData)
	if err != nil {
		return &system.CommonRpcRes{}, err
	}
	creater := l.svcCtx.SystemStore.AsInputData.Create()
	tools.SchemaUpdateAny(creater.Mutation(), inputData, "id")
	return system.JsonResult(creater.Save(l.ctx))
}
