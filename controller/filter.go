package controller

import (
	"fmt"

	"github.com/gofiber/fiber"
)

var Filter filterInterface = &filter{}

type filterInterface interface {
	Get(c *fiber.Ctx)
}

type filter struct{}

func (f filter) Get(c *fiber.Ctx) {
	msg := fmt.Sprintf("✋ %s", c.Params("*"))
	c.SendString(msg)
}
