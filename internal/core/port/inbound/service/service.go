package portservice

import (
	"context"

	"github.com/muhfaris/rocket-hexagonal/internal/core/domain"
)

type BookService interface {
	CreateBook(ctx context.Context, bodyRequest domain.CreateBook) error
	GetAllBooks(ctx context.Context) error
	DeleteBook(ctx context.Context, deleteBookParams domain.DeleteBookParams) error
	GetBookById(ctx context.Context, bookId domain.GetBookByIdParams) error
	UpdateBook(ctx context.Context, updateBookParams domain.UpdateBookParams, bodyRequest domain.UpdateBook) error
}
