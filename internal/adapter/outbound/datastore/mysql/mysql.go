package mysqladapter

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhfaris/rocket-hexagonal/internal/core/port/outbound/repository"
)

var (
	mysqlConn *sql.DB
	once      sync.Once
)

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

func New(opt MySQLConfig) repository.MySQLRepository {
	var (
		ctx = context.Background()
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", opt.Username, opt.Password, opt.Host, opt.Port, opt.DB)
	)

	if mysqlConn != nil {
		return &Client{
			conn: mysqlConn,
		}
	}

	once.Do(func() {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		// Verify connection
		if err = db.PingContext(ctx); err != nil {
			panic(err)
		}

		// Configure connection pool
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)

		mysqlConn = db
	})

	return &Client{
		conn: mysqlConn,
	}
}
