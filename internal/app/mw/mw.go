package mw

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole := c.Request().Header.Get("User-Role")

		if strings.EqualFold(userRole, roleAdmin) {
			fmt.Println("red button user detected")
		}

		return next(c)
	}
}
