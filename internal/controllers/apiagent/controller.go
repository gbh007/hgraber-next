package apiagent

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/ogen-go/ogen/ogenerrors"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
	"hgnext/open_api/agentAPI"
)

type parsingUseCases interface {
	CheckBooks(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error)
	ParseBook(ctx context.Context, u url.URL) (entities.AgentBookDetails, error)
	DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
	CheckPages(ctx context.Context, pages []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error)
}

type exportUseCases interface {
	ImportArchive(ctx context.Context, body io.Reader, deduplicate bool, autoVerify bool) (uuid.UUID, error)
}

type Controller struct {
	startAt time.Time
	logger  *slog.Logger
	tracer  trace.Tracer
	addr    string
	debug   bool

	ogenServer *agentAPI.Server

	parsingUseCases parsingUseCases
	exportUseCases  exportUseCases

	token string
}

func New(
	startAt time.Time,
	logger *slog.Logger,
	tracer trace.Tracer,
	parsingUseCases parsingUseCases,
	exportUseCases exportUseCases,
	addr string,
	debug bool,
	token string,
) (*Controller, error) {
	c := &Controller{
		startAt: startAt,
		logger:  logger,
		tracer:  tracer,
		addr:    addr,
		debug:   debug,
		token:   token,

		parsingUseCases: parsingUseCases,
		exportUseCases:  exportUseCases,
	}

	ogenServer, err := agentAPI.NewServer(
		c, c,
		agentAPI.WithErrorHandler(methodErrorHandler),
		agentAPI.WithMethodNotAllowed(methodNotAllowed),
		agentAPI.WithNotFound(methodNotFound),
	)
	if err != nil {
		return nil, err
	}

	c.ogenServer = ogenServer

	return c, nil
}

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t agentAPI.HeaderAuth) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request, allowed string) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", allowed)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.Header().Set("Allow", allowed)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusMethodNotAllowed)

	// TODO: не игнорировать ошибку
	_ = json.NewEncoder(w).Encode(agentAPI.ErrorResponse{
		InnerCode: "method not allowed",
		Details:   agentAPI.NewOptString("method not allowed, allowed methods " + allowed),
	})
}

func methodNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError) // Специально не делаем 404, т.к. на нее может быть завязано особое поведение метода

	if r.Method != http.MethodOptions {
		// TODO: не игнорировать ошибку
		_ = json.NewEncoder(w).Encode(agentAPI.ErrorResponse{
			InnerCode: "method not found",
			Details:   agentAPI.NewOptString("method not found"),
		})
	}
}

func methodErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var (
		httpCode         int    = http.StatusInternalServerError
		errorCode        string = "internal error"
		errorDescription string = "missing error"
	)

	if err != nil {
		errorDescription = err.Error()
	}

	switch {
	case errors.Is(err, ogenerrors.ErrSecurityRequirementIsNotSatisfied):
		httpCode = http.StatusUnauthorized
		errorCode = "unauthorized"
	case errors.Is(err, errorAccessForbidden):
		httpCode = http.StatusForbidden
		errorCode = "forbidden"
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)

	if r.Method == http.MethodOptions {
		return
	}

	// TODO: не игнорировать ошибку
	_ = json.NewEncoder(w).Encode(agentAPI.ErrorResponse{
		InnerCode: errorCode,
		Details:   agentAPI.NewOptString(errorDescription),
	})
}
