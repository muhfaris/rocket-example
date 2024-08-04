package portservice

import (
	"context"

	"github.com/muhfaris/rocket-example/internal/core/domain"
)

type BookService interface {
	GetBooks(ctx context.Context) error
	GetBookById(ctx context.Context, bookId domain.GetBookByIdParameters) error
	LendBook(ctx context.Context, bodyRequest domain.LendBookProperties) error
	ReturnBook(ctx context.Context, bodyRequest domain.ReturnBookProperties) error
}
