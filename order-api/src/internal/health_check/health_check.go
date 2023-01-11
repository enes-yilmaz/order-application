package health_check

import "github.com/labstack/echo/v4"

func HealthCheck(c echo.Context) error {
	return c.NoContent(200)
}
