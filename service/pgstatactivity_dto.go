package service

import (
	"github.com/emadghaffari/agileful/domain/entity"
	"github.com/emadghaffari/agileful/domain/repository"
)

// Get: get the postgres activity by filters
func (a activity) Get(req entity.Filter) ([]entity.PgStatActivity, error) {
	resp, err := repository.PGActivity.Get(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
