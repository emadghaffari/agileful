package repository

import (
	"fmt"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/stretchr/testify/assert"

	"github.com/emadghaffari/agileful/client/postgres"
	"github.com/emadghaffari/agileful/config"
	"github.com/emadghaffari/agileful/domain/entity"
)

type sqlMock struct {
	err error
}

func (s *sqlMock) Connect(config config.Config) error {
	return s.err
}

func (s sqlMock) DB() *pg.DB {
	return pg.Connect(&pg.Options{})
}

func (s sqlMock) Close() error {
	return nil
}

func (s *sqlMock) Query(model interface{}, query interface{}, params ...interface{}) (res orm.Result, err error) {
	return nil, s.err
}

func TestGet(t *testing.T) {
	testCases := []struct {
		desc   string
		order  string
		query  string
		limit  int
		offset int
		resp   []entity.PgStatActivity
		err    error
	}{
		{
			desc:   "a",
			order:  "",
			query:  "",
			limit:  0,
			offset: 0,
			resp:   nil,
			err:    fmt.Errorf("dial tcp: lookup postgres: Temporary failure in name resolution"),
		},
		{
			desc:   "a",
			order:  "",
			query:  "",
			limit:  0,
			offset: 0,
			resp:   []entity.PgStatActivity{},
			err:    nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			test := sqlMock{}
			test.err = tC.err
			postgres.Storage = &test

			_, err := PGActivity.Get(entity.Filter{
				Query:  tC.query,
				Order:  tC.order,
				Limit:  tC.limit,
				Offset: tC.offset,
			})
			if err != nil {
				assert.Equal(t, tC.err.Error(), err.Error())
			}
		})
	}

}
