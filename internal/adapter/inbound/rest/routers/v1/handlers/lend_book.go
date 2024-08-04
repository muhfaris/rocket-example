package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-example/internal/adapter/inbound/rest/routers/v1/response"
)

type LendBookProperties struct {
	Bookid string `json:"bookId"`
	Userid string `json:"userId"`
}

func (h *Handler) LendBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		var bodyRequest LendBookProperties
		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.LendBook(ctx , bodyRequest )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
