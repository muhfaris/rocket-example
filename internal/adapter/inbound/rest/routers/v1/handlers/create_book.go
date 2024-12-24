package handlersv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers/v1/response"
)

type CreateBook struct {
	Availablecopies int    `json:"availableCopies"`
	Isbn            string `json:"isbn"`
	Publisheddate   string `json:"publishedDate"`
	Title           string `json:"title"`
	Author          string `json:"author"`
}

// POST /api/books handler
// @Summary Create a new book
// @Tags [Books]
// @Accept application/json
// @Param body body &amp;[object] true &#34;Request body&#34;
// @Success 201 {object} #/components/schemas/Book &#34;Book created successfully&#34;
// @Router /api/books [POST]

func (h *Handler) CreateBook() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		/*
			var (
				ctx = c.UserContext()
				svc = h.services.BookService()
				)
		*/

		var bodyRequest CreateBook
		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		/*
			// Transform request into domain model
			result, err:=  svc.CreateBook(ctx , bodyRequest )
			if err != nil {
				return err
			}
		*/

		return response.Success(c, "I'm Alive!")
	}
}
