package response

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	libcontext "github.com/muhfaris/rocket-hexagonal/shared/context"
)

type R struct {
	Data     any      `json:"data,omitempty"`
	Message  string   `json:"message,omitempty"`
	Errors   []Error  `json:"errors,omitempty"`
	Metadata Metadata `json:"metadata"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Metadata struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Latency    string      `json:"latency"`
	RequestID  string      `json:"request_id"`
}

type Pagination struct {
	Next  *string `json:"next,omitempty"`
	Prev  *string `json:"prev,omitempty"`
	Limit *int    `json:"limit,omitempty"`
	Total *int    `json:"total,omitempty"`
	Page  *int    `json:"page,omitempty"`
}

func Success(c *fiber.Ctx, data any) error {
	var (
		ctx          = c.UserContext()
		startLatency = ctx.Value(libcontext.ContextLatency)
		defaultCode  = fiber.StatusOK
		requestID    = c.Get(fiber.HeaderXRequestID)
		latency      string
	)

	c.SendStatus(defaultCode)

	start, ok := startLatency.(time.Time)
	if ok {
		latencyMs := float64(time.Since(start).Nanoseconds() / 1e6)
		latency = fmt.Sprintf("%.2fms", latencyMs)
	}

	result := R{
		Data: data,
		Metadata: Metadata{
			RequestID: requestID,
			Latency:   latency,
		},
	}

	return c.JSON(result)
}

func Fail(c *fiber.Ctx, err error) error {
	var (
		ctx          = c.UserContext()
		startLatency = ctx.Value(libcontext.ContextLatency)
		defaultCode  = fiber.StatusOK
		requestID    = c.Get(fiber.HeaderXRequestID)
		latency      string
	)

	c.SendStatus(defaultCode)

	start, ok := startLatency.(time.Time)
	if ok {
		latencyMs := float64(time.Since(start).Nanoseconds() / 1e6)
		latency = fmt.Sprintf("%.2fms", latencyMs)
	}

	result := R{
		Errors: []Error{
			{
				Message: err.Error(),
			},
		},
		Metadata: Metadata{
			RequestID: requestID,
			Latency:   latency,
		},
	}

	return c.JSON(result)
}
