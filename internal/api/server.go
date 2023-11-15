package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nxy7/tcp-assignment/internal/api/handlers"
)

// Returns server using default handler created from env variables
func MakeDefaultServer() *echo.Echo {
	h := handlers.FromEnv()
	return MakeServer(h)
}

func MakeServer(h handlers.HandlerCtx) *echo.Echo {
	e := echo.New()

	e.GET("/health", h.HandleHealth)

	return e

}
