package app

import (
	"github.com/muhfaris/rocket-hexagonal/config"
	redisadapter "github.com/muhfaris/rocket-hexagonal/internal/adapter/outbound/cache/redis"
	mysqladapter "github.com/muhfaris/rocket-hexagonal/internal/adapter/outbound/datastore/mysql"
	portregistry "github.com/muhfaris/rocket-hexagonal/internal/core/port/inbound/registry"
	"github.com/muhfaris/rocket-hexagonal/internal/core/port/outbound/repository"
)

func InitializeRepository() portregistry.Repository {
	return &App{

		cacheRepository: redisadapter.New(
			redisadapter.RedisOptions{
				Addr:     config.App.Cache.Redis.Addr,
				Username: config.App.Cache.Redis.Username,
				Password: config.App.Cache.Redis.Password,
				DB:       config.App.Cache.Redis.DB,
			},
		),

		mysqlRepository: mysqladapter.New(mysqladapter.MySQLConfig{
			Host:     config.App.Datastore.MySQL.Host,
			Port:     config.App.Datastore.MySQL.Port,
			Username: config.App.Datastore.MySQL.Username,
			Password: config.App.Datastore.MySQL.Password,
			DB:       config.App.Datastore.MySQL.DB,
		}),
	}
}

// App struct would need to be updated to include new repository types
type App struct {
	cacheRepository repository.CacheRepository
	mysqlRepository repository.MySQLRepository
}

func (a *App) CacheRepository() repository.CacheRepository {
	return a.cacheRepository
}

func (a *App) MySQLRepository() repository.MySQLRepository {
	return a.mysqlRepository
}
