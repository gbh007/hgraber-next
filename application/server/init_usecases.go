package server

import (
	"github.com/gbh007/hgraber-next/usecases/agentusecase"
	"github.com/gbh007/hgraber-next/usecases/attributeusecase"
	"github.com/gbh007/hgraber-next/usecases/bffusecase"
	"github.com/gbh007/hgraber-next/usecases/bookusecase"
	"github.com/gbh007/hgraber-next/usecases/cleanupusecase"
	"github.com/gbh007/hgraber-next/usecases/deduplicatorusecase"
	"github.com/gbh007/hgraber-next/usecases/exportusecase"
	"github.com/gbh007/hgraber-next/usecases/filesystemusecase"
	"github.com/gbh007/hgraber-next/usecases/hproxyusecase"
	"github.com/gbh007/hgraber-next/usecases/labelusecase"
	"github.com/gbh007/hgraber-next/usecases/massloadusecase"
	"github.com/gbh007/hgraber-next/usecases/parsingusecase"
	"github.com/gbh007/hgraber-next/usecases/rebuilderusecase"
	"github.com/gbh007/hgraber-next/usecases/systemusecase"
)

func (a *App) initUseCases() {
	a.bookUseCases = bookusecase.New(a.Logger, a.storage)
	a.parsingUseCases = parsingusecase.New(
		a.Logger,
		a.storage,
		a.agentSystem,
		a.fileStorageAdapter,
		a.bookUseCases,
		a.Config.Parsing.ParseBookTimeout,
		a.Config.AttributeRemap.Auto,
		a.Config.AttributeRemap.AllLower,
	)
	a.exportUseCases = exportusecase.New(
		a.Logger,
		a.storage,
		a.fileStorageAdapter,
		a.agentSystem,
		a.tmpStorage,
		a.bookUseCases,
		a.Config.AttributeRemap.Auto,
		a.Config.AttributeRemap.AllLower,
	)
	a.deduplicateUseCases = deduplicatorusecase.New(a.Logger, a.storage, a.Tracer)
	a.cleanupUseCases = cleanupusecase.New(a.Logger, a.Tracer, a.storage, a.fileStorageAdapter)
	a.reBuilderUseCases = rebuilderusecase.New(
		a.Logger,
		a.Tracer,
		a.storage,
		a.Config.AttributeRemap.Auto,
		a.Config.AttributeRemap.AllLower,
	)
	a.fsUseCases = filesystemusecase.New(a.Logger, a.storage, a.fileStorageAdapter, a.tmpStorage)
	a.bffUseCases = bffusecase.New(a.Logger, a.storage, a.deduplicateUseCases)
	a.attributeUseCases = attributeusecase.New(a.Logger, a.storage, a.Config.AttributeRemap.AllLower)
	a.labelUseCases = labelusecase.New(a.Logger, a.storage)
	a.massloadUseCases = massloadusecase.New(a.Logger, a.storage, a.tmpStorage, a.agentSystem)

	a.systemUseCases = systemusecase.New(
		a.Logger,
		a.storage,
		a.tmpStorage,
		a.deduplicateUseCases,
		a.cleanupUseCases,
		a.workersController,
		a.attributeUseCases,
	)

	a.agentUseCases = agentusecase.New(a.Logger, a.agentSystem, a.storage)
	a.hProxyUseCases = hproxyusecase.New(
		a.Logger,
		a.storage,
		a.agentSystem,
		a.Config.Parsing.ParseBookTimeout,
		a.Config.AttributeRemap.Auto,
		a.Config.AttributeRemap.AllLower,
	)
}
