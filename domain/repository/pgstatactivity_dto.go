package repository

import (
	"fmt"

	"github.com/emadghaffari/agileful/client/postgres"
	"github.com/emadghaffari/agileful/domain/entity"
)

// Get: get the postgres activity by filters
func (a activity) Get(req entity.Filter) ([]entity.PgStatActivity, error) {
	var m = []entity.PgStatActivity{}
	_, err := postgres.Storage.Query(&m,
		fmt.Sprintf("SELECT current_timestamp - query_start as runtime,* FROM pg_stat_activity WHERE wait_event IS NOT NULL AND backend_type = 'client backend' AND query like '%s %%' order by runtime %s limit %d offset %d", req.Query, req.Order, req.Limit, req.Offset))
	if err != nil {
		return nil, err
	}
	return m, nil
}
