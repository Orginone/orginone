package user

import (
	"encoding/json"
	"errors"
	"math"
	"orginone/common/models"
)

func PageResult(page *PageRequest, total int64, data interface{}, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	return JsonResult(&models.PageResponse{
		Records:     data,
		Current:     page.Offset/page.Limit + 1,
		Size:        page.Limit,
		SearchCount: page.SearchCount,
		Total:       total,
		Pages:       int64(math.Ceil(float64(total / page.Limit))),
	}, err)
}

func JsonResult(data interface{}, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	if data == nil {
		return &CommonRpcRes{}, errors.New("not found")
	}
	resJson, err := json.Marshal(data)
	return Result(string(resJson), err)
}

func Result(data string, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	return &CommonRpcRes{
		Json: data,
	}, nil
}
