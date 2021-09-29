package main

import (
	"log"

	"github.com/gofiber/fiber"

	"github.com/emadghaffari/agileful/app"
)

func main() {
	fbr := fiber.New()

	app.Base.StartApplication(fbr)

	log.Fatal(fbr.Listen(":3000"))
}
