package postgres

import (
	"context"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/emadghaffari/agileful/config"
)

// Connect method job is connect to postgres database and check migration
func (p *psql) Connect(cnf config.Config) error {
	var err error

	once.Do(func() {
		p.db = pg.Connect(&pg.Options{
			User:                  cnf.POSTGRES.Username,
			Password:              cnf.POSTGRES.Password,
			Addr:                  cnf.POSTGRES.Host,
			Database:              cnf.POSTGRES.Schema,
			RetryStatementTimeout: true,
		})

		if err = p.db.Ping(context.Background()); err != nil {
			return
		}
	})

	return err
}

// DB: return database
func (p *psql) DB() *pg.DB {
	return p.db
}

// Close: close the db pool
func (p *psql) Close() error {
	once = sync.Once{}
	return p.db.Close()
}

// Query: create custom query
func (p *psql) Query(model interface{}, query interface{}, params ...interface{}) (res orm.Result, err error) {
	return p.DB().Query(model, query, params)
}
