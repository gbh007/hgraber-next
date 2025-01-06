package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"strconv"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"

	"hgnext/open_api/serverAPI"
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
	_ = json.NewEncoder(w).Encode(serverAPI.ErrorResponse{
		InnerCode: "method not allowed",
		Details:   serverAPI.NewOptString("method not allowed, allowed methods " + allowed),
	})
}

func methodNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError) // Специально не делаем 404, т.к. на нее может быть завязано особое поведение метода

	if r.Method != http.MethodOptions {
		// TODO: не игнорировать ошибку
		_ = json.NewEncoder(w).Encode(serverAPI.ErrorResponse{
			InnerCode: "method not found",
			Details:   serverAPI.NewOptString("method not found"),
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
	case errors.Is(err, errPanicDetected):
		httpCode = http.StatusInternalServerError
		errorCode = "panic"
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)

	if r.Method == http.MethodOptions {
		return
	}

	// TODO: не игнорировать ошибку
	_ = json.NewEncoder(w).Encode(serverAPI.ErrorResponse{
		InnerCode: errorCode,
		Details:   serverAPI.NewOptString(errorDescription),
	})
}

func (c *Controller) simplePanicRecover(req middleware.Request, next middleware.Next) (returnedResponse middleware.Response, returnedError error) {
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
