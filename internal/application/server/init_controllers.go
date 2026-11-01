package server

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/internal/controllers/agentcacheapi"
	"github.com/gbh007/hgraber-next/internal/controllers/serverapi"
	"github.com/gbh007/hgraber-next/internal/controllers/workermanager"
	"github.com/gbh007/hgraber-next/internal/domain/systemmodel"
	"github.com/gbh007/hgraber-next/internal/experimental/mcp"
)

//nolint:cyclop,funlen // будет исправлено позднее
func (a *App) initControllers(_ context.Context) error {
	workerUnits := make([]workermanager.WorkerUnit, 0, len(a.Config.Workers))

	for _, wCfg := range a.Config.Workers {
		var unit workermanager.WorkerUnit

		switch wCfg.GetName() {
		case systemmodel.WorkerNameBook:
			unit = workermanager.NewBookParser(a.parsingUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNamePage:
			unit = workermanager.NewPageDownloader(a.parsingUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameHasher:
			unit = workermanager.NewHasher(a.fsUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameExporter:
			unit = workermanager.NewExporter(a.exportUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameTasker:
			unit = workermanager.NewTasker(a.tmpStorage, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameFileValidator:
			unit = workermanager.NewFileValidator(a.fsUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameFileTransferer:
			unit = workermanager.NewFileTransfer(a.fsUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameMassloadSizer:
			unit = workermanager.NewMassloadSize(a.massloadUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameMassloadAttributeSizer:
			unit = workermanager.NewMassloadAttributeSize(
				a.massloadUseCases,
				a.Logger,
				wCfg,
				a.metricProvider,
			)

		case systemmodel.WorkerNameMassloadCalculation:
			unit = workermanager.NewMassloadCalculation(a.massloadUseCases, a.Logger, wCfg, a.metricProvider)

		case systemmodel.WorkerNameBookCalculation:
			unit = workermanager.NewBookCalculation(a.bookUseCases, a.Logger, wCfg, a.metricProvider)

		default:
			continue
		}

		workerUnits = append(workerUnits, unit)
	}

	a.workersController.AddUnits(workerUnits...)

	apiController, err := serverapi.New(
		a.Logger,
		a.Config.API,
		a.metricProvider,
		a.parsingUseCases,
		a.agentUseCases,
		a.exportUseCases,
		a.deduplicateUseCases,
		a.systemUseCases,
		a.reBuilderUseCases,
		a.fsUseCases,
		a.bffUseCases,
		a.attributeUseCases,
		a.labelUseCases,
		a.bookUseCases,
		a.hProxyUseCases,
		a.massloadUseCases,
	)
	if err != nil {
		return fmt.Errorf("fail to create api server: %w", err)
	}

	a.asyncController.RegisterRunner(apiController)

	if a.Config.AgentServer.Addr != "" {
		apiAgentController, err := agentcacheapi.New(
			a.Config.AgentServer,
			time.Now(),
			a.Logger,
			a.parsingUseCases,
			a.exportUseCases,
			a.metricProvider,
		)
		if err != nil {
			return fmt.Errorf("fail to create api agent: %w", err)
		}

		a.asyncController.RegisterRunner(apiAgentController)
	}

	if a.Config.Application.Metric.Enabled() {
		infoCollector, err := a.metricProvider.NewSystemInfoCollector(
			a.Logger,
			a.systemUseCases,
			a.storage,
			a.Config.Application.Metric,
		)
		if err != nil {
			return fmt.Errorf("fail to create info collector: %w", err)
		}

		a.asyncController.RegisterRunner(infoCollector)
	}

	if a.Config.MCPServer.Addr != "" {
		mcpController := mcp.New(
			a.Logger,
			a.Config.MCPServer.Addr,
			a.Config.MCPServer.Token,
			a.bffUseCases,
			a.attributeUseCases,
			a.hProxyUseCases,
			a.parsingUseCases,
			a.Config.MCPServer.Debug,
			a.Config.MCPServer.Mutate,
		)

		a.asyncController.RegisterRunner(mcpController)
	}

	return nil
}
