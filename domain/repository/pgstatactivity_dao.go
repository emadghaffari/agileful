package repository

import "github.com/emadghaffari/agileful/domain/entity"

var PGActivity activityInterface = &activity{}

type activityInterface interface {
	Get(req entity.Filter) ([]entity.PgStatActivity,error)
}

type activity struct{}