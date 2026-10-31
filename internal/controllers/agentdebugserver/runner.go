package agentdebugserver

import (
	"context"
	"errors"
	"html/template"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (c *Controller) Name() string {
	return "api server"
}

func (c *Controller) Start(parentCtx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	l := slog.NewLogLogger(c.logger.Handler(), slog.LevelError)

	echoRouter := echo.New()

	echoRouter.HideBanner = true
	echoRouter.HidePort = true
	echoRouter.Validator = validator{}
	echoRouter.Logger.SetOutput(l.Writer())
	echoRouter.StdLogger = l

	echoRouter.GET("/", c.pageMain)
	echoRouter.GET("/config", c.pageConfig)
	echoRouter.GET("/parsing", c.pageParsing)
	echoRouter.POST("/parsing", c.pageParsing)
	echoRouter.GET("/css/main.css", echo.StaticFileHandler("templates/main.css", TemplateDir))

	echoRouter.HTTPErrorHandler = func(err error, c echo.Context) {
		_ = c.Render(http.StatusOK, "error.html.gotmpl", err)
	}

	echoRouter.Renderer = &Template{
		templates: template.Must(template.ParseFS(TemplateDir, "templates/*.gotmpl")),
	}

	echoRouter.Use(echo.WrapMiddleware(cors))
	echoRouter.Use(middleware.Recover())

	server := &http.Server{
		Addr:     c.addr,
		ErrorLog: l,
	}

	go func() {
		defer close(done)

		c.logger.InfoContext(parentCtx, "debug server start")
		defer c.logger.InfoContext(parentCtx, "debug server stop")

		err := echoRouter.StartServer(server)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	go func() {
		<-parentCtx.Done()
		c.logger.InfoContext(parentCtx, "stopping debug server")

		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(parentCtx), time.Second*10)
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	return done, nil
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}
