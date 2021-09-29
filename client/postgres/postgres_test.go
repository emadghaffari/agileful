package postgres

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emadghaffari/agileful/config"
)

func TestConnect(t *testing.T) {
	// Storage.Connect(config)
	tests := []struct {
		step string
		conf config.Config
		err  error
	}{
		{
			step: "A",
			conf: config.Config{
				Debug: true,
			},
			err: fmt.Errorf("dial tcp 127.0.0.1:5432: connect: connection refused"),
		},
		{
			step: "B",
			conf: config.Config{
				Debug: true,
				POSTGRES: config.Database{
					Username: "admin",
					Password: "password",
					Host:     "127.0.0.1:5432",
					Schema:   "schema",
				},
			},
			err: nil,
		}}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			err := Storage.Connect(tc.conf)
			if tc.err != nil {
				assert.Equal(t, tc.err.Error(), err.Error())
			}
		})
	}
}

func TestDB(t *testing.T) {
	db := Storage.DB()
	if db == nil {
		assert.Error(t, fmt.Errorf("error in database get DB"))
	}
}

func TestGet(t *testing.T) {
	db, _ := Storage.DB().Begin()
	if db == nil {
		assert.Error(t, fmt.Errorf("error in database get DB"))
	}
}

func TestClose(t *testing.T) {
	err := Storage.Close()
	if err != nil {
		assert.Error(t, fmt.Errorf("error in close DB"))
	}
}
