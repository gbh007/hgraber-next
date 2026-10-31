package debugserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cnt *Controller[T]) pageMain(c echo.Context) error {
	return c.Render(http.StatusOK, "main.html.gotmpl", nil)
}
