package portregistry

import (
	portservice "github.com/muhfaris/rocket-hexagonal/internal/core/port/inbound/service"
	"github.com/muhfaris/rocket-hexagonal/internal/core/port/outbound/repository"
)

type Service interface {
	BookService() portservice.BookService
}

type Repository interface {
	CacheRepository() repository.CacheRepository
}
