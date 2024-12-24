package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers/v1/response"
)

type DeleteBookParams struct {
	Bookid string `params:"bookId"`
}

// DELETE /api/books/{bookId} handler
// @Summary Delete a specific book
// @Tags [Books]
// @Param bookId path &amp;[string] true &#34;The ID of the book&#34;
// @Router /api/books/{bookId} [DELETE]

func (h *Handler) DeleteBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		var deleteBookParams DeleteBookParams
		if err := c.ParamsParser(&deleteBookParams); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.DeleteBook(ctx , deleteBookParams )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
