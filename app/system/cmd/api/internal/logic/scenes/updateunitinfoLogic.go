package scenes

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"
	"orginone/common/rpcs/system"
	"orginone/common/schema"
	"orginone/common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateunitinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateunitinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateunitinfoLogic {
	return UpdateunitinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateunitinfoLogic) Updateunitinfo(req types.UpdateUnitInfoReq) (resp *types.CommonResponse, err error) {
	unitEntity := new(schema.AsUnit)
	err = tools.Copy(unitEntity, req)
	if err != nil {
		return types.Failed(http.StatusBadRequest, err)
	}
	reqJson, _ := json.Marshal(req)
	isUnitBasicPropertyExist := strings.Contains(string(reqJson), "\"unitBasicProperty\"")
	isAccountingSystemExist := strings.Contains(string(reqJson), "\"accountingSystem\"")
	isUnitReformExist := strings.Contains(string(reqJson), "\"unitReform\"")
	isDepartmentCategoryNameExist := strings.Contains(string(reqJson), "\"departmentCategoryName\"")
	//phone regexp
	re := regexp.MustCompile("^1|0(3|4|5|7|8|9)[0-9]{9}$")
	if !tools.IsNilOrEmpty(unitEntity.LinkPhone) && !re.Match([]byte(unitEntity.LinkPhone)) {
		return types.Failed(http.StatusBadRequest, errors.New("手机号不符合规范!"))
	}
	if isUnitBasicPropertyExist && unitEntity.UnitBasicProperty != 23 && unitEntity.BudgetCode == 0 {
		return types.Failed(http.StatusBadRequest, errors.New("财政预算代码不能为空!"))
	}
	if !isAccountingSystemExist {
		return types.Failed(http.StatusBadRequest, errors.New("单位执行会计制度不能为空!"))
	}
	if (unitEntity.AccountingSystem == 30 && isUnitBasicPropertyExist && unitEntity.UnitBasicProperty != 31) ||
		(unitEntity.AccountingSystem == 50 && tools.ArrIndexOf([]int{11, 12, 13, 14, 15, 16, 17, 21}, unitEntity.UnitBasicProperty) != -1) {
		return types.Failed(http.StatusBadRequest, errors.New("单位基本性质与单位执行会计制度不匹配!"))
	}
	if tools.ArrIndexOf([]int{11, 12, 13, 14, 15, 16, 17, 31}, unitEntity.UnitBasicProperty) != -1 && isUnitReformExist {
		return types.Failed(http.StatusBadRequest, errors.New("事业单位改革类型应为空!"))
	}
	if tools.ArrIndexOf([]int{21, 22, 23}, unitEntity.UnitBasicProperty) != -1 && !isUnitReformExist {
		return types.Failed(http.StatusBadRequest, errors.New("事业单位改革类型不应为空!"))
	}
	if !isDepartmentCategoryNameExist {
		return types.Failed(http.StatusBadRequest, errors.New("部门标识不能为空!"))
	}
	buf, err := json.Marshal(unitEntity)
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	_, err = l.svcCtx.SystemRpc.UpdateUnitInfo(l.ctx, &system.CommonRpcReq{Json: string(buf)})
	if err != nil {
		return types.Failed(http.StatusInternalServerError, err)
	}
	return types.Successful(true)
}
