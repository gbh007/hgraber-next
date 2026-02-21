package apiservercore

import (
	"strconv"
)

type APIError struct {
	Code      int
	InnerCode string
	Details   string
}

func (ae APIError) Error() string {
	return strconv.Itoa(ae.Code) + ": " + ae.InnerCode
}
