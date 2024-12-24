package middlewares

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	libcontext "github.com/muhfaris/rocket-hexagonal/shared/context"
)

func Latency() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var (
			start = time.Now()
			ctx   = c.UserContext()
		)

		ctx = context.WithValue(ctx, libcontext.ContextLatency, start)
		c.SetUserContext(ctx)
		err := c.Next()
		if err != nil {
			return err
		}

		// Calculate latency in milliseconds
		latencyMs := float64(time.Since(start).Nanoseconds()) / 1e6

		// Add latency to response header
		c.Set(fiber.HeaderServerTiming, fmt.Sprintf("%.2fms", latencyMs))
		return nil
	}
}
