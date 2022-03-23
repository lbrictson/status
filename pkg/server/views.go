package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) indexView(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Name": "Bob",
	})
}
