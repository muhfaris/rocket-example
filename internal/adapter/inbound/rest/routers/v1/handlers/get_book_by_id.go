package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers/v1/response"
)

type GetBookByIdParams struct {
	Bookid string `params:"bookId"`
}

// GET /api/books/{bookId} handler
// @Summary Retrieve details of a specific book
// @Tags [Books]
// @Param bookId path &amp;[string] true &#34;The ID of the book&#34;
// @Success 200 {object} #/components/schemas/Book &#34;Details of the book&#34;
// @Router /api/books/{bookId} [GET]

func (h *Handler) GetBookById() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		var bookId GetBookByIdParams
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
