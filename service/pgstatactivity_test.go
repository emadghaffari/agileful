package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emadghaffari/agileful/domain/entity"
	"github.com/emadghaffari/agileful/domain/repository"
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
		desc string
		resp []entity.PgStatActivity
		err  error
	}{
		{
			desc: "a",
			resp: nil,
			err:  fmt.Errorf("err"),
		},
		{
			desc: "b",
			resp: []entity.PgStatActivity{},
			err:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tst := MockActivity{}
			tst.er = tC.err
			tst.resp = tC.resp
			repository.PGActivity = &tst

			_, err := PGActivity.Get(entity.Filter{})
			if err != nil {
				assert.Equal(t, tC.err, err)
			}
		})
	}
}
