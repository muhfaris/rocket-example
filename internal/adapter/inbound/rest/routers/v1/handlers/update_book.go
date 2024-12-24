package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers/v1/response"
)

type UpdateBookParams struct {
	Bookid string `params:"bookId"`
}
type UpdateBook struct {
	Author          string `json:"author"`
	Availablecopies int    `json:"availableCopies"`
	Isbn            string `json:"isbn"`
	Publisheddate   string `json:"publishedDate"`
	Title           string `json:"title"`
}

// PUT /api/books/{bookId} handler
// @Summary Update details of a specific book
// @Tags [Books]
// @Param bookId path &amp;[string] true &#34;The ID of the book&#34;
// @Accept application/json
// @Param body body &amp;[object] true &#34;Request body&#34;
// @Success 200 {object} #/components/schemas/Book &#34;Book updated successfully&#34;
// @Router /api/books/{bookId} [PUT]

func (h *Handler) UpdateBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		var updateBookParams UpdateBookParams
		if err := c.ParamsParser(&updateBookParams); err != nil {
			return err
		}

		var bodyRequest UpdateBook
		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.UpdateBook(ctx , updateBookParams , bodyRequest )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
