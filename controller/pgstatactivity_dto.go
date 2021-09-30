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
		c.SendBytes(entity.Error(err.Error()))
		return
	}

	if id := entity.QueryValidate[req.Query]; id == 0 {
		c.SendStatus(http.StatusBadRequest)
		c.SendBytes(entity.Error("Query is invalid [SELECT,INSERT,UPDATE,DELETE]"))
		return
	}


	if id := entity.OrderByValidate[req.Order]; id == 0 {
		c.SendStatus(http.StatusBadRequest)
		c.SendBytes(entity.Error("Order By is invalid [desc,asc]"))
		return
	}

	if req.Limit == 0 {
		c.SendStatus(http.StatusBadRequest)
		c.SendBytes(entity.Error("limit is invalid"))
		return
	}

	resp,err := service.PGActivity.Get(req)
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		c.SendBytes(entity.Error(err.Error()))
		return
	}

	bts,_:=json.Marshal(entity.PgStatActivityResponse{Count: len(resp),Data: resp})
	c.SendStatus(http.StatusOK)
	c.SendBytes(bts)
}

