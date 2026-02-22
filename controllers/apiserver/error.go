package apiserver

import (
	"context"
	"errors"
	"net/http"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/validate"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *Controller) NewError(ctx context.Context, err error) *serverapi.ErrorResponseStatusCode {
	if err == nil {
		return &serverapi.ErrorResponseStatusCode{
			StatusCode: http.StatusInternalServerError,
			Response: serverapi.ErrorResponse{
				InnerCode: "unexpected",
				Details:   serverapi.NewOptString("missing error"),
			},
		}
	}

	var ae apiservercore.APIError

	if errors.As(err, &ae) {
		return &serverapi.ErrorResponseStatusCode{
			StatusCode: ae.Code,
			Response: serverapi.ErrorResponse{
				InnerCode: "api error",
				Details: serverapi.OptString{
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

	var (
		validateError      *validate.Error
		decodeBodyError    *ogenerrors.DecodeBodyError
		decodeParamError   *ogenerrors.DecodeParamError
		decodeParamsError  *ogenerrors.DecodeParamsError
		decodeRequestError *ogenerrors.DecodeRequestError
	)

	//nolint:goconst // будет исправлено позднее
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
	case errors.As(err, &decodeBodyError):
		httpCode = http.StatusBadRequest
		errorCode = "validate"
	case errors.As(err, &decodeParamError):
		httpCode = http.StatusBadRequest
		errorCode = "validate"
	case errors.As(err, &decodeParamsError):
		httpCode = http.StatusBadRequest
		errorCode = "validate"
	case errors.As(err, &decodeRequestError):
		httpCode = http.StatusBadRequest
		errorCode = "validate"
	}

	return &serverapi.ErrorResponseStatusCode{
		StatusCode: httpCode,
		Response: serverapi.ErrorResponse{
			InnerCode: errorCode,
			Details:   serverapi.NewOptString(errorDescription),
		},
	}
}
