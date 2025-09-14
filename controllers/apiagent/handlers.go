package apiagent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

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
	_ = json.NewEncoder(w).Encode(agentapi.ErrorResponse{
		InnerCode: "method not allowed",
		Details:   agentapi.NewOptString("method not allowed, allowed methods " + allowed),
	})
}

func methodNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Специально не делаем 404, т.к. на нее может быть завязано особое поведение метода
	w.WriteHeader(http.StatusInternalServerError)

	if r.Method != http.MethodOptions {
		// TODO: не игнорировать ошибку
		_ = json.NewEncoder(w).Encode(agentapi.ErrorResponse{
			InnerCode: "method not found",
			Details:   agentapi.NewOptString("method not found"),
		})
	}
}

var errPanicDetected = errors.New("panic detected")

func stackTrace(skip, count int) []string {
	result := []string{}

	pc := make([]uintptr, count)
	n := runtime.Callers(skip, pc)

	pc = pc[:n]

	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()

		result = append(result, frame.File+":"+strconv.Itoa(frame.Line))

		if !more {
			break
		}
	}

	return result
}

func (c *Controller) methodErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var (
		httpCode         = http.StatusInternalServerError
		errorCode        = "internal error"
		errorDescription = "missing error"
	)

	if err != nil {
		errorDescription = err.Error()
	}

	if c.logErrorHandler {
		c.logger.ErrorContext(
			ctx, "handle api agent server error",
			slog.Any("error", err),
		)
	}

	validateError := new(validate.Error)

	switch {
	case errors.Is(err, ogenerrors.ErrSecurityRequirementIsNotSatisfied):
		httpCode = http.StatusUnauthorized
		errorCode = "unauthorized"
	case errors.Is(err, errorAccessForbidden):
		httpCode = http.StatusForbidden
		errorCode = "forbidden"
	case errors.Is(err, errPanicDetected):
		httpCode = http.StatusInternalServerError
		errorCode = "panic"
	case errors.As(err, &validateError):
		httpCode = http.StatusBadRequest
		errorCode = "validate"
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)

	if r.Method == http.MethodOptions {
		return
	}

	// TODO: не игнорировать ошибку
	_ = json.NewEncoder(w).Encode(agentapi.ErrorResponse{
		InnerCode: errorCode,
		Details:   agentapi.NewOptString(errorDescription),
	})
}

func (c *Controller) simplePanicRecover(
	req middleware.Request,
	next middleware.Next,
) (returnedResponse middleware.Response, returnedError error) {
	defer func() {
		p := recover()
		if p != nil {
			c.logger.WarnContext(
				req.Context, "panic detected",
				slog.Any("panic", p),
				slog.Any("trace", stackTrace(3, 50)),
			)

			returnedResponse = middleware.Response{}
			returnedError = fmt.Errorf("%w: %v", errPanicDetected, p)
		}
	}()

	return next(req)
}

func (c *Controller) metricsMiddleware(
	req middleware.Request,
	next middleware.Next,
) (returnedResponse middleware.Response, returnedError error) {
	tStart := time.Now()
	operation := req.OperationName
	addr := c.addr

	c.metricProvider.HTTPServerIncActive(addr, operation)

	defer func() {
		c.metricProvider.HTTPServerDecActive(addr, operation)

		// FIXME: Каст не сможет сработать из-за особеностей Go, необходимо реализовать другой подход.
		_, hasError := returnedResponse.Type.(*agentapi.ErrorResponse)
		success := !hasError && returnedError == nil

		c.metricProvider.HTTPServerAddHandle(addr, operation, success, time.Since(tStart))
	}()

	return next(req)
}
