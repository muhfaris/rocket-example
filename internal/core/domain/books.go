package domain

type CreateBook struct {
	Availablecopies int    `json:"availableCopies"`
	Isbn            string `json:"isbn"`
	Publisheddate   string `json:"publishedDate"`
	Title           string `json:"title"`
	Author          string `json:"author"`
}
type DeleteBookParams struct {
	Bookid string `params:"bookId"`
}
type GetBookByIdParams struct {
	Bookid string `params:"bookId"`
}
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
