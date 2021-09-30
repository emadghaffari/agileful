package controller

import "github.com/gofiber/fiber"


var Filter filterInterface = &filter{}

type filterInterface interface {
	Get(c *fiber.Ctx)
}

type filter struct{}
