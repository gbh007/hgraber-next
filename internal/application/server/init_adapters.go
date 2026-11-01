package server

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/internal/dataproviders/agentapi"
	"github.com/gbh007/hgraber-next/internal/dataproviders/filestorage"
	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb"
	"github.com/gbh007/hgraber-next/internal/dataproviders/tmpdata"
	"github.com/gbh007/hgraber-next/internal/domain/core"
)

func (a *App) initAdapters(ctx context.Context) error {
	var err error

	a.tmpStorage = tmpdata.New()

	a.storage, err = serverdb.New(
		ctx,
		a.Logger,
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

	a.agentSystem, err = agentapi.New(agents, a.Config.Parsing.AgentTimeout)
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
