package app

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/emadghaffari/agileful/client/postgres"
	"github.com/emadghaffari/agileful/config"
)

func configPath() string {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	return strings.TrimSpace(string(path)) + "/config.yaml"
}

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
	return nil, nil
}

func TestInitPostgres(t *testing.T) {

	testCases := []struct {
		desc string
		err  error
	}{
		{
			desc: "first",
			err:  nil,
		},
		{
			desc: "second",
			err:  fmt.Errorf("some error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			test := sqlMock{}
			test.err = tc.err
			postgres.Storage = &test
			if err := Base.initPostgres(); err != nil {
				assert.Equal(t, err.Error(), tc.err.Error())
			}
		})
	}
}

type cfMock struct {
	err error
}

func (cf *cfMock) Get() config.Config {
	return config.Config{}
}
func (cf *cfMock) SetDebug(bool)       {}
func (cf *cfMock) Set(bts []byte) bool { return true }

func TestInitConfigs(t *testing.T) {
	testCases := []struct {
		desc string
		path string
		err  error
	}{
		{
			desc: "a",
			path: configPath(),
			err:  nil,
		},
		{
			desc: "a",
			path: "",
			err:  fmt.Errorf("init config read file err  open : no such file or directory"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			test := &cfMock{}
			test.err = tc.err
			err := Base.initConfigs(tc.path)
			if err != nil {
				assert.Equal(t, err, tc.err)
			}
		})
	}
}

func TestInitEndpoints(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Base.initEndpoints(fiber.New())
		})
	}
}

type BaseMock struct {
	errConf error
	errDB   error
}

func (base BaseMock) StartApplication(fbr *fiber.App) {}
func (base BaseMock) initEndpoints(fbr *fiber.App)    {}
func (base BaseMock) initConfigs(path string) error   { return base.errConf }
func (base BaseMock) initPostgres() error             { return base.errDB }

func TestStartApplication(t *testing.T) {
	testCases := []struct {
		desc    string
		errConf error
		errDB   error
	}{
		{
			desc:    "a",
			errConf: nil,
			errDB:   nil,
		},
		{
			desc:    "b",
			errConf: fmt.Errorf("error"),
			errDB:   nil,
		},
		{
			desc:    "c",
			errConf: nil,
			errDB:   fmt.Errorf("error"),
		},
	}
	for _, tC := range testCases {
		app := App{}
		t.Run(tC.desc, func(t *testing.T) {
			psql := sqlMock{}
			psql.err = nil
			postgres.Storage = &psql

			base := BaseMock{}
			base.errConf = tC.errConf
			base.errDB = tC.errDB
			Base = &base

			fbr := fiber.New()
			go func() {
				app.StartApplication(fbr)
			}()
			time.Sleep(time.Second * 2)
			fbr.Shutdown()
		})
	}
}
