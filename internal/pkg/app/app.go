package app

import (
	"fmt"
	"log"

	"github.com/GTech1256/junior-test/internal/app/endpoint"
	"github.com/GTech1256/junior-test/internal/app/mw"
	"github.com/GTech1256/junior-test/internal/app/service"
	"github.com/labstack/echo/v4"
)

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
