package cache

type CacheConfig struct {
	Network    string `json:",optional,default=tcp"`
	Addr       string
	Db         int    `json:",optional,default=0"`
	Username   string `json:",optional"`
	Password   string `json:",optional"`
	PoolSize   int    `json:",optional,default=5"`
	MaxEntries int    `json:",optional,default=-1"`
	TTL        int64  `json:",optional,default=1"`
}
