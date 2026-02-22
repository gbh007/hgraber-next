package apiservercore

import (
	"strconv"
)

type APIError struct {
	Code    int
	Details string
}

func (ae APIError) Error() string {
	return strconv.Itoa(ae.Code) + ": " + ae.Details
}
