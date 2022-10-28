package models

import (
	"orginone/common/cache"
)

const (
	// DevMode means development mode.
	DevMode = "dev"
	// ProMode means production mode.
	ProMode = "pro"
)

//数据存储配置
type DbStoreConfig struct {
	WorkerId     int64  `json:",optional,default=1"`
	DataCenterId int64  `json:",optional,default=1"`
	Driver       string `json:",optional,default=mysql"`
	Mode         string `json:",optional,default=dev"`
	DataSource   string
	Cache        cache.CacheConfig `json:",optional"`
}

//分页返回结构
type PageResponse struct {
	Current     int64       `json:"current"`
	Pages       int64       `json:"pages"`
	Records     interface{} `json:"records"`
	SearchCount bool        `json:"searchCount"`
	Size        int64       `json:"size"`
	Total       int64       `json:"total"`
}

type DistributionConfigItem struct {
	Check int64  `json:"check"`
	Id    int64  `json:"id,string"`
	Label string `json:"label"`
	Name  string `json:"name"`
	Type  int64  `json:"type"`
}

type DistributionConfigModel struct {
	DistributeType string                    `json:"distributeType"`
	Contain        int64                     `json:"contain"`
	ConfigList     []*DistributionConfigItem `json:"configList"`
}
