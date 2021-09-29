package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		desc  string
		debug bool
	}{
		{
			desc:  "a",
			debug: false,
		},
		{
			desc:  "b",
			debug: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Confs.SetDebug(tC.debug)
			cnfs := Confs.Get()
			if cnfs.Debug != tC.debug {
				assert.Equal(t, tC.debug, cnfs.Debug)
			}
		})
	}
}

func TestGetDebug(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "a",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Confs.SetDebug(true)
		})
	}
}

func TestSet(t *testing.T) {
	testCases := []struct {
		desc string
		data []byte
		err  error
	}{
		{
			desc: "a",
			data: []byte("{}"),
			err:  nil,
		},
		{
			desc: "b",
			data: nil,
			err:  nil,
		},
		{
			desc: "c",
			data: []byte{byte(12)},
			err:  fmt.Errorf("failed to unmarshal the config: yaml: control characters are not allowed"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := Confs.Set(tC.data)
			if err != nil {
				assert.Equal(t, tC.err.Error(), err.Error())
			}
		})
	}
}
