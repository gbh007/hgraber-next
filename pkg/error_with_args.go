package pkg

import (
	"bytes"
	"fmt"
)

type errorArgument struct {
	Key   string
	Value any
}

type errorWithArgs struct { //nolint:errname // ложно-положительное
	origin error
	msg    string
	args   []errorArgument
}

func ErrorArgument(
	key string,
	value any,
) errorArgument { //nolint:revive // будет исправлено позднее
	return errorArgument{
		Key:   key,
		Value: value,
	}
}

func WrapError(
	origin error,
	msg string,
	args ...errorArgument,
) error {
	if origin == nil {
		return nil
	}

	return &errorWithArgs{
		origin: origin,
		msg:    msg,
		args:   args,
	}
}

func ErrorWithArgs(
	msg string,
	args ...errorArgument,
) error {
	if msg == "" &&
		len(args) == 0 {
		return nil
	}

	return &errorWithArgs{
		msg:  msg,
		args: args,
	}
}

func (err *errorWithArgs) Error() string {
	buf := bytes.Buffer{}

	buf.WriteString(err.msg)

	if len(err.args) > 0 {
		if buf.Len() != 0 {
			buf.WriteString(" ")
		}

		buf.WriteString("(")

		for i, arg := range err.args {
			if i != 0 {
				buf.WriteString(", ")
			}

			buf.WriteString(arg.Key)
			buf.WriteString("=")
			buf.WriteString(fmt.Sprint(arg.Value))
		}

		buf.WriteString(")")
	}

	if err.origin != nil {
		if buf.Len() != 0 {
			buf.WriteString(": ")
		}

		buf.WriteString(err.origin.Error())
	}

	return buf.String()
}

func (err *errorWithArgs) Unwrap() error {
	return err.origin
}
