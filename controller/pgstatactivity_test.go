package controller

import (
	"fmt"
	"testing"

	"github.com/emadghaffari/agileful/domain/entity"
	"github.com/emadghaffari/agileful/service"
)

type MockActivity struct {
	resp []entity.PgStatActivity
	er   error
}

func (m MockActivity) Get(req entity.Filter) ([]entity.PgStatActivity, error) {
	return m.resp, m.er
}

func TestGet(t *testing.T) {
	testCases := []struct {
		desc         string
		order        string
		query        string
		limit        int
		offset       int
		resp         []entity.PgStatActivity
		serviceError error
		err          error
	}{
		{
			desc:         "a",
			order:        "",
			query:        "",
			limit:        0,
			offset:       0,
			resp:         []entity.PgStatActivity{},
			serviceError: nil,
			err:          fmt.Errorf("x"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tst := MockActivity{}
			tst.er = tC.serviceError
			tst.resp = tC.resp
			service.PGActivity = &tst

		})
	}
}
