package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Status handles health checks from the orchestrator.
func Status(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
