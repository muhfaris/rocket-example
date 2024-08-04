package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-example/internal/adapter/inbound/rest/routers/v1/response"
)

type GetBookByIdParameters struct {
	Bookid string `params:"bookId"`
}

func (h *Handler) GetBookById() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		var bookId GetBookByIdParameters
		if err := c.ParamsParser(&bookId); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.GetBookById(ctx , bookId )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
