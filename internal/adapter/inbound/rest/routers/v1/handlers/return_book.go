package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-example/internal/adapter/inbound/rest/routers/v1/response"
)

type ReturnBookProperties struct {
	Bookid string `json:"bookId"`
	Userid string `json:"userId"`
}

func (h *Handler) ReturnBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		var bodyRequest ReturnBookProperties
		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.ReturnBook(ctx , bodyRequest )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
