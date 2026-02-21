package apiagent

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

type apiError struct {
	Code      int
	InnerCode string
	Details   string
}

func (ae apiError) Error() string {
	return strconv.Itoa(ae.Code) + ": " + ae.InnerCode
}

func (c *Controller) NewError(ctx context.Context, err error) *agentapi.ErrorResponseStatusCode {
	if err == nil {
		return &agentapi.ErrorResponseStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: agentapi.ErrorResponse{
				InnerCode: "unexpected",
				Details:   agentapi.NewOptString("empty error"),
			},
		}
	}

	var ae apiError

	if errors.As(err, &ae) {
		return &agentapi.ErrorResponseStatusCode{
			StatusCode: ae.Code,
			Response: agentapi.ErrorResponse{
				InnerCode: ae.InnerCode,
				Details: agentapi.OptString{
					Value: ae.Details,
					Set:   len(ae.Details) > 0,
				},
			},
		}
	}

	var (
		httpCode         = http.StatusInternalServerError
		errorCode        = "internal error"
		errorDescription = err.Error()
	)

	validateError := new(validate.Error)

	switch {
	case errors.Is(err, ogenerrors.ErrSecurityRequirementIsNotSatisfied):
		httpCode = http.StatusUnauthorized
		errorCode = "unauthorized"
	case errors.Is(err, errAccessForbidden):
		httpCode = http.StatusForbidden
		errorCode = "forbidden"
	case errors.Is(err, errPanicDetected):
		httpCode = http.StatusInternalServerError
		errorCode = "panic"
	case errors.As(err, &validateError):
		httpCode = http.StatusBadRequest
		errorCode = "validate"
	}

	return &agentapi.ErrorResponseStatusCode{
		StatusCode: httpCode,
		Response: agentapi.ErrorResponse{
			InnerCode: errorCode,
			Details:   agentapi.NewOptString(errorDescription),
		},
	}
}
