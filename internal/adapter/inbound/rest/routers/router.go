package routers

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/rocket-hexagonal/config"
	handlersv1 "github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers/v1/handlers"
	"github.com/muhfaris/rocket-hexagonal/internal/adapter/inbound/rest/routers/v1/middlewares"
	portadapter "github.com/muhfaris/rocket-hexagonal/internal/core/port/inbound/adapter"

	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Router struct {
	Client *fiber.App
	Port   int
}

func Init(port int) portadapter.Rest {
	r := fiber.New(fiber.Config{
		AppName:                  "rocket-hexagonal",
		EnableSplittingOnParsers: config.App.Fiber.EnableSplittingOnParsers,
		EnablePrintRoutes:        config.App.Fiber.EnablePrintRoutes,
	})

	r.Use(requestid.New())
	r.Use(middlewares.Latency())

	// public
	// publicGroup := r.Group("/")
	// publicGroup.Get("/health", handlersv1.Health())

	h := handlersv1.New()

	bookGroup := r.Group("/api/books")
	bookGroup.Post("/api/books", h.CreateBook()).Name("CreateBook")
	bookGroup.Get("/api/books", h.GetAllBooks()).Name("GetAllBooks")
	bookGroup.Delete("/api/books/:bookId", h.DeleteBook()).Name("DeleteBook")
	bookGroup.Get("/api/books/:bookId", h.GetBookById()).Name("GetBookById")
	bookGroup.Put("/api/books/:bookId", h.UpdateBook()).Name("UpdateBook")

	return &Router{
		Client: r,
		Port:   port,
	}
}

func (r *Router) Run() error {
	// gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var err error

	go func() {
		port := fmt.Sprintf(":%d", r.Port)
		slog.Info("Listening on port", "port", port)
		if err := r.Client.Listen(port); err != nil {
			return
		}
	}()

	<-ctx.Done()
	err = r.Client.Shutdown()
	if err != nil {
		return fmt.Errorf("failed to shutdown: %v", err)
	}

	return err
}
