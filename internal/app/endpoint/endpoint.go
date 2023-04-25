package endpoint

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GTech1256/junior-test/internal/app/service"
	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft(endTime time.Time) (int32, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Handler(c echo.Context) error {

	daysLeft, err := e.s.DaysLeft(service.TimeMyDate)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("%v days left until 1 Jan 2025", daysLeft))
}
