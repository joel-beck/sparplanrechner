package middleware

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ZerologMiddleware() echo.MiddlewareFunc {
	// Configure zerolog to log human-readable output
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Caller().Logger()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Call the next handler
			err := next(c)
			if err != nil {
				c.Error(err)
			}

			// Log relevant fields (Method, URI, and Status)
			log.Debug().
				Str("method", c.Request().Method).
				Str("uri", c.Request().RequestURI).
				Int("status", c.Response().Status).
				Msg("Request-Response Log")

			return nil
		}
	}
}
