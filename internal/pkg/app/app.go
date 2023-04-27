package app

import (
	"fmt"
	"log"

	"github.com/GTech1256/junior-test/internal/metrics"

	"github.com/GTech1256/junior-test/internal/app/endpoint"
	"github.com/GTech1256/junior-test/internal/app/mw"
	"github.com/GTech1256/junior-test/internal/app/service"
	"github.com/labstack/echo/v4"
)

// TODO: перенести в env
const METRICS_ADDRESS = "127.0.0.1:8082"

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	// Root level middleware
	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/", a.e.Handler)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server running")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (a *App) StartMetrics() {
	go func() {
		metrics.Listen(METRICS_ADDRESS)
	}()
}
