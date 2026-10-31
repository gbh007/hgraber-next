package entities

import "context"

type Runner interface {
	Start(context.Context) (chan struct{}, error)
	Name() string
}
