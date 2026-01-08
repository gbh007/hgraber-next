package server

import (
	"context"
	"fmt"
	"time"

	"github.com/gbh007/hgraber-next/controllers/apiagent"
	"github.com/gbh007/hgraber-next/controllers/apiserver"
	"github.com/gbh007/hgraber-next/controllers/workermanager"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/experimental/mcp"
)

//nolint:cyclop,funlen // будет исправлено позднее
func (a *App) initControllers(_ context.Context) error {
	workerUnits := make([]workermanager.WorkerUnit, 0, len(a.Config.Workers))

	for _, wCfg := range a.Config.Workers {
		var unit workermanager.WorkerUnit

		switch wCfg.GetName() {
		case systemmodel.WorkerNameBook:
			unit = workermanager.NewBookParser(a.parsingUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNamePage:
			unit = workermanager.NewPageDownloader(a.parsingUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameHasher:
			unit = workermanager.NewHasher(a.fsUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameExporter:
			unit = workermanager.NewExporter(a.exportUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameTasker:
			unit = workermanager.NewTasker(a.tmpStorage, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameFileValidator:
			unit = workermanager.NewFileValidator(a.fsUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameFileTransferer:
			unit = workermanager.NewFileTransfer(a.fsUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameMassloadSizer:
			unit = workermanager.NewMassloadSize(a.massloadUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameMassloadAttributeSizer:
			unit = workermanager.NewMassloadAttributeSize(
				a.massloadUseCases,
				a.Logger,
				a.Tracer,
				wCfg,
				a.metricProvider,
			)

		case systemmodel.WorkerNameMassloadCalculation:
			unit = workermanager.NewMassloadCalculation(a.massloadUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		case systemmodel.WorkerNameBookCalculation:
			unit = workermanager.NewBookCalculation(a.bookUseCases, a.Logger, a.Tracer, wCfg, a.metricProvider)

		default:
			continue
		}

		workerUnits = append(workerUnits, unit)
	}

	a.workersController.AddUnits(workerUnits...)

	apiController, err := apiserver.New(
		a.Logger,
		a.Tracer,
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
		apiAgentController, err := apiagent.New(
			a.Config.AgentServer,
			time.Now(),
			a.Logger,
			a.Tracer,
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
			a.Tracer,
			a.Config.MCPServer.Addr,
			a.Config.MCPServer.Token,
			a.bffUseCases,
			a.Config.MCPServer.Debug,
		)

		a.asyncController.RegisterRunner(mcpController)
	}

	return nil
}
