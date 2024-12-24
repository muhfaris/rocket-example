package service

import (
	"context"

	"github.com/muhfaris/rocket-hexagonal/internal/core/domain"
	portservice "github.com/muhfaris/rocket-hexagonal/internal/core/port/inbound/service"
)

type BookService struct{}

func NewBookService() portservice.BookService {
	return &BookService{}
}
func (s *BookService) CreateBook(ctx context.Context, bodyRequest domain.CreateBook) error {
	return nil
}
func (s *BookService) GetAllBooks(ctx context.Context) error {
	return nil
}
func (s *BookService) DeleteBook(ctx context.Context, deleteBookParams domain.DeleteBookParams) error {
	return nil
}
func (s *BookService) GetBookById(ctx context.Context, bookId domain.GetBookByIdParams) error {
	return nil
}
func (s *BookService) UpdateBook(ctx context.Context, updateBookParams domain.UpdateBookParams, bodyRequest domain.UpdateBook) error {
	return nil
}
