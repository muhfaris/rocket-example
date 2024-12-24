package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers/v1/response"
)

// GET /api/books handler
// @Summary Retrieve a list of all books
// @Tags [Books]
// @Success 200 {object}  &#34;A list of books&#34;
// @Router /api/books [GET]

func (h *Handler) GetAllBooks() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		/*
			// Transform request into domain model
			result, err:=  svc.GetAllBooks(ctx )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
