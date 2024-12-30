package api

import (
	"fmt"
	"github.com/loveuer/uzone/pkg/opt"
	"os"
	"runtime/debug"
	"time"

	"github.com/loveuer/uzone/pkg/tool"
)

func NewRecover(enableStackTrace bool) HandlerFunc {
	return func(c *Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				if enableStackTrace {
					os.Stderr.WriteString(fmt.Sprintf("recovered from panic: %v\n%s\n", r, debug.Stack()))
				} else {
					os.Stderr.WriteString(fmt.Sprintf("recovered from panic: %v\n", r))
				}

				_ = c.Status(500).SendString(fmt.Sprint(r))
			}
		}()

		return c.Next()
	}
}

func NewLogger() HandlerFunc {
	return func(c *Ctx) error {
		var (
			now = time.Now()
			ip  = c.IP()
		)

		traceId := c.Context().Value(opt.TraceKey)
		c.Locals(opt.TraceKey, traceId)

		err := c.Next()

		c.Writer.Header().Set(opt.TraceKey, fmt.Sprint(traceId))

		//status, _ := strconv.Atoi(c.Writer.Header().Get(opt.HttpStatusHeader))
		duration := time.Since(now)

		//msg := fmt.Sprintf(" %15s | %d[%3d] | %s | %6s | %s", ip, c.StatusCode, status, tool.HumanDuration(duration.Nanoseconds()), c.Method(), c.Path())
		fields := []any{
			"ip", ip,
			"status", c.StatusCode,
			"duration", tool.HumanDuration(duration.Nanoseconds()),
			"method", c.Method(),
			"path", c.Path(),
		}

		switch {
		case c.StatusCode >= 500:
			c.UseLogger().With(fields...).Error("")
		case c.StatusCode >= 400:
			c.UseLogger().With(fields...).Warn("")
		default:
			c.UseLogger().With(fields...).Info("")
		}

		return err
	}
}
