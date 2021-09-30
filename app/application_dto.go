package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gofiber/fiber"

	"github.com/emadghaffari/agileful/client/postgres"
	"github.com/emadghaffari/agileful/config"
	"github.com/emadghaffari/agileful/controller"
)

type Test struct {
	ID   int
	Name string
}

func (app App) StartApplication(fbr *fiber.App) {

	// init configs
	dir, _ := os.Getwd()
	if err := Base.initConfigs(dir + "/config.yaml"); err != nil {
		log.Println(err)
		return
	}

	if err := Base.initPostgres(); err != nil {
		log.Println(err)
		return
	}
	defer postgres.Storage.DB().Close()

	// init endpoints
	Base.initEndpoints(fbr)

	log.Print(fbr.Listen(":3000"))
}

func (app App) initConfigs(path string) error {
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("init config read file err  %s", err.Error())
	}

	return config.Confs.Set(bts)
}

func (app App) initPostgres() error {
	if err := postgres.Storage.Connect(config.Confs.Get()); err != nil {
		return err
	}

	fmt.Printf("postgres database loaded successfully \n")
	return nil
}

func (app App) initEndpoints(fbr *fiber.App) {
	fbr.Post("/search", controller.Filter.Get)
}
