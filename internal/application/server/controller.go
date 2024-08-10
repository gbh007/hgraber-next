package server

import (
	"context"
	"fmt"
	"log/slog"
)

type Runner interface {
	Start(context.Context) (chan struct{}, error)
	Name() string
}

func New(logger *slog.Logger) *Controller {
	return &Controller{
		logger: logger,
	}
}

type Controller struct {
	logger *slog.Logger

	runnerChannels []chan struct{}
	runners        []Runner
	after          []func()
}

func (c *Controller) RegisterRunner(runner Runner) {
	c.runners = append(c.runners, runner)
}

func (c *Controller) RegisterAfterStop(handler func()) {
	c.after = append(c.after, handler)
}

func (c *Controller) Serve(parentCtx context.Context) error {
	ctx, cnl := context.WithCancel(parentCtx)
	defer cnl()

	for _, r := range c.runners {
		exitCh, err := r.Start(ctx)
		if err != nil {
			err = fmt.Errorf("start %s: %w", r.Name(), err)

			c.logger.ErrorContext(ctx, err.Error())

			return err
		}

		c.runnerChannels = append(c.runnerChannels, exitCh)
	}

	// Дожидаемся завершения потоков
	for _, exitCh := range c.runnerChannels {
		<-exitCh
	}

	// Проходим по всем послеостановочным функциям
	for _, handler := range c.after {
		handler()
	}

	return nil
}
