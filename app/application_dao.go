package app

import "github.com/gofiber/fiber"

var (
	Base Application = &App{}
)

// Application interface for start application
type Application interface {
	StartApplication(fbr *fiber.App)
	initEndpoints(fbr *fiber.App)
	initConfigs(path string) error
	initPostgres() error
}

type App struct{}
