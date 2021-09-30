package postgres

import (
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/emadghaffari/agileful/config"
)

var (
	Storage store = &psql{}
	once    sync.Once
)

// store interface is interface for store things into postgres
type store interface {
	Connect(config config.Config) error
	DB() *pg.DB
	Close() error
	Query(model interface{}, query interface{}, params ...interface{}) (res orm.Result, err error)
}

// postgres struct
type psql struct {
	db *pg.DB
}
