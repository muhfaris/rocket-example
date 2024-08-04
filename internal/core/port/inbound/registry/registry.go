package portregistry

import portservice "github.com/muhfaris/rocket-example/internal/core/port/inbound/service"

type Service interface {
	BookService() portservice.BookService
}
