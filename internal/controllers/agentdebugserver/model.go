package agentdebugserver

import "errors"

var errMissingValidator = errors.New("missing validator")

type modelValidator interface {
	Validate() error
}

type validator struct{}

func (validator) Validate(i any) error {
	if v, ok := i.(modelValidator); ok {
		return v.Validate()
	}

	return errMissingValidator
}
