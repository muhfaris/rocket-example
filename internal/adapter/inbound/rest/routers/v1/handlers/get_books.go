package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-example/internal/adapter/inbound/rest/routers/v1/response"
)

func (h *Handler) GetBooks() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		/*
			// Transform request into domain model
			result, err:=  svc.GetBooks(ctx )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
