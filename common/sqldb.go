package common

import (
	"orginone/common/cache"
	"orginone/common/models"
	"orginone/common/schema"
	"orginone/common/tools/snowflake"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DbStore struct {
	*schema.Client
	Cache *cache.Redis
}

func NewDbStore(c models.DbStoreConfig) *DbStore {
	_, err := snowflake.NewSnowflake(c.DataCenterId, c.WorkerId)
	if err != nil {
		panic("connect database error")
	}
	db, err := sql.Open(c.Driver, c.DataSource)
	if err != nil {
		panic("connect database error")
	}
	// init cache driver
	drv, redis := cache.NewDriver(db, c.Cache)

	entClient := schema.NewClient(schema.Driver(drv))
	if c.Mode == models.DevMode {
		entClient = entClient.Debug()
	}
	return &DbStore{
		Client: entClient,
		Cache:  redis,
	}
}
