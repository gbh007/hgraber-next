package server

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/adapters/agent"
	"github.com/gbh007/hgraber-next/adapters/filestorage"
	"github.com/gbh007/hgraber-next/adapters/postgresql"
	"github.com/gbh007/hgraber-next/adapters/tmpdata"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (a *App) initAdapters(ctx context.Context) error {
	var err error

	a.tmpStorage = tmpdata.New()

	a.storage, err = postgresql.New(
		ctx,
		a.Logger,
		a.Tracer,
		a.metricProvider,
		a.Config.Storage.DebugPGX,
		a.Config.Storage.Connection,
		a.Config.Storage.MaxConnections,
	)
	if err != nil {
		return fmt.Errorf("fail init postgres: %w", err)
	}

	agents, err := a.storage.Agents(ctx, core.AgentFilter{})
	if err != nil {
		return fmt.Errorf("fail load agents from storage: %w", err)
	}

	a.agentSystem, err = agent.New(agents, a.Config.Parsing.AgentTimeout)
	if err != nil {
		return fmt.Errorf("fail init agent system: %w", err)
	}

	a.fileStorageAdapter = filestorage.New(
		a.Logger,
		a.agentSystem,
		a.storage,
		a.metricProvider,
		a.Config.FileStorage.TryReconnect,
	)

	err = a.fileStorageAdapter.Init(ctx, true)
	if err != nil {
		return fmt.Errorf("fail init file system: %w", err)
	}

	return nil
}
