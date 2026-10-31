package agentdebugserver

import (
	"bytes"
	"net/http"

	"github.com/gbh007/hgraber-next/external/config"
	"github.com/labstack/echo/v4"
)

func (cnt *Controller[T]) pageConfig(c echo.Context) error {
	buff := &bytes.Buffer{}

	err := config.ExportToWriter(buff, &cnt.config, ".yaml")
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "config.html.gotmpl", buff.String())
}
