package main

import (
	"github.com/gofiber/fiber"

	"github.com/emadghaffari/agileful/app"
)

func main() {
	fbr := fiber.New()
	app.Base.StartApplication(fbr)
}
