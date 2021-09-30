package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber"

	"github.com/emadghaffari/agileful/domain/entity"
	"github.com/emadghaffari/agileful/service"
)

func (f filter) Get(c *fiber.Ctx) {
	c.Type("json", "utf-8")
	req := entity.Filter{}
	if err := c.BodyParser(&req); err != nil {
		c.SendStatus(http.StatusInternalServerError)
		bts, _ := json.Marshal(entity.Error{Message: err.Error()})
		c.SendBytes(bts)
		return
	}

	if id := entity.QueryValidate[req.Query]; id == 0 {
		c.SendStatus(http.StatusBadRequest)
		bts, _ := json.Marshal(entity.Error{Message: "Query is invalid [SELECT,INSERT,UPDATE,DELETE]"})
		c.SendBytes(bts)
		return
	}

	if id := entity.OrderByValidate[req.Order]; id == 0 {
		c.SendStatus(http.StatusBadRequest)
		bts, _ := json.Marshal(entity.Error{Message: "Order By is invalid [desc,asc]"})
		c.SendBytes(bts)
		return
	}

	if req.Limit == 0 {
		c.SendStatus(http.StatusBadRequest)
		bts, _ := json.Marshal(entity.Error{Message: "limit is invalid"})
		c.SendBytes(bts)
		return
	}

	resp, err := service.PGActivity.Get(req)
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		bts, _ := json.Marshal(entity.Error{Message: err.Error()})
		c.SendBytes(bts)
		return
	}

	bts, _ := json.Marshal(entity.PgStatActivityResponse{Count: len(resp), Data: resp})
	c.SendStatus(http.StatusOK)
	c.SendBytes(bts)
}
