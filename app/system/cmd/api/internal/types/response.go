package types

import (
	"encoding/json"
	"errors"
	"net/http"
	"orginone/common/models"
	"orginone/common/rpcs/system"
)

func PageResult(res *system.CommonRpcRes, err error) (*CommonResponse, error) {
	if err != nil {
		return Failed(http.StatusInternalServerError, err)
	}
	var data *models.PageResponse
	err = json.Unmarshal([]byte(res.Json), &data)
	return JsonResult(data, err)
}
func BoolResult(data interface{}, err error) (*CommonResponse, error) {
	if err != nil {
		return Failed(http.StatusInternalServerError, err)
	}
	return Successful(true)
}
func JsonResult(data interface{}, err error) (*CommonResponse, error) {
	if err != nil {
		return Failed(http.StatusInternalServerError, err)
	}
	if data == nil {
		return Failed(http.StatusNoContent, errors.New("not found"))
	}
	return Successful(data)
}
func CommonResult(data *system.CommonRpcRes, err error) (*CommonResponse, error) {
	if err != nil {
		return Failed(http.StatusInternalServerError, err)
	}
	if data == nil {
		return Failed(http.StatusNoContent, errors.New("not found"))
	}
	return Successful(data.Json)
}
func Successful(data interface{}) (*CommonResponse, error) {
	return &CommonResponse{
		Data:    data,
		Success: true,
		Code:    200,
		Msg:     "success!",
	}, nil
}

func Failed(code int64, err error) (*CommonResponse, error) {
	errmsg := "unknown."
	if err != nil {
		errmsg = err.Error()
	}
	return &CommonResponse{
		Data:    nil,
		Success: false,
		Code:    code,
		Msg:     "request is failed, error message:" + errmsg,
	}, nil
}

func BadRequest(err error) *CommonResponse {
	return &CommonResponse{
		Data:    nil,
		Success: false,
		Code:    http.StatusBadRequest,
		Msg:     "request is failed, error message:" + err.Error(),
	}
}
