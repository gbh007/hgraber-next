# pkg

## Imports

|     Name     |                Path                 | Inner | Count |
|:------------:|:-----------------------------------:|:-----:|:-----:|
|     fmt      |                 fmt                 |  ❌   |   3   |
|    assert    | github.com/stretchr/testify/assert  |  ❌   |   2   |
|   testing    |               testing               |  ❌   |   2   |
|    bytes     |                bytes                |  ❌   |   1   |
|     sql      |            database/sql             |  ❌   |   1   |
|    errors    |               errors                |  ❌   |   1   |
| pyroscope-go |   github.com/grafana/pyroscope-go   |  ❌   |   1   |
|   require    | github.com/stretchr/testify/require |  ❌   |   1   |
|      io      |                 io                  |  ❌   |   1   |
|     slog     |              log/slog               |  ❌   |   1   |
|     sync     |                sync                 |  ❌   |   1   |

## Used by

|        Name         |                                            Path                                            |
|:-------------------:|:------------------------------------------------------------------------------------------:|
|       adapter       |           [/adapters/agent/internal/adapter](adapters/agent/internal/adapter.md)           |
|    generatorcore    |   [/adapters/metric/generator/generatorcore](adapters/metric/generator/generatorcore.md)   |
|      deadhash       |     [/adapters/postgresql/internal/deadhash](adapters/postgresql/internal/deadhash.md)     |
|        model        |        [/adapters/postgresql/internal/model](adapters/postgresql/internal/model.md)        |
|     repository      |   [/adapters/postgresql/internal/repository](adapters/postgresql/internal/repository.md)   |
|       tmpdata       |                          [/adapters/tmpdata](adapters/tmpdata.md)                          |
|       server        |                        [/application/server](application/server.md)                        |
|      apiagent       |                      [/controllers/apiagent](controllers/apiagent.md)                      |
|    agenthandlers    |       [/controllers/apiserver/agenthandlers](controllers/apiserver/agenthandlers.md)       |
|    apiservercore    |       [/controllers/apiserver/apiservercore](controllers/apiserver/apiservercore.md)       |
|  attributehandlers  |   [/controllers/apiserver/attributehandlers](controllers/apiserver/attributehandlers.md)   |
|    bookhandlers     |        [/controllers/apiserver/bookhandlers](controllers/apiserver/bookhandlers.md)        |
| deduplicatehandlers | [/controllers/apiserver/deduplicatehandlers](controllers/apiserver/deduplicatehandlers.md) |
|     fshandlers      |          [/controllers/apiserver/fshandlers](controllers/apiserver/fshandlers.md)          |
|   hproxyhandlers    |      [/controllers/apiserver/hproxyhandlers](controllers/apiserver/hproxyhandlers.md)      |
|    labelhandlers    |       [/controllers/apiserver/labelhandlers](controllers/apiserver/labelhandlers.md)       |
|  massloadhandlers   |    [/controllers/apiserver/massloadhandlers](controllers/apiserver/massloadhandlers.md)    |
|   systemhandlers    |      [/controllers/apiserver/systemhandlers](controllers/apiserver/systemhandlers.md)      |
|       worker        |               [/controllers/internal/worker](controllers/internal/worker.md)               |
|    workermanager    |                 [/controllers/workermanager](controllers/workermanager.md)                 |
|         bff         |                                [/domain/bff](domain/bff.md)                                |
|        core         |                               [/domain/core](domain/core.md)                               |
|      external       |                                  [/external](external.md)                                  |
|    agentusecase     |                     [/usecases/agentusecase](usecases/agentusecase.md)                     |
|     bffusecase      |                       [/usecases/bffusecase](usecases/bffusecase.md)                       |
|   cleanupusecase    |                   [/usecases/cleanupusecase](usecases/cleanupusecase.md)                   |
| deduplicatorusecase |              [/usecases/deduplicatorusecase](usecases/deduplicatorusecase.md)              |
|    exportusecase    |                    [/usecases/exportusecase](usecases/exportusecase.md)                    |
|  filesystemusecase  |                [/usecases/filesystemusecase](usecases/filesystemusecase.md)                |
|    hproxyusecase    |                    [/usecases/hproxyusecase](usecases/hproxyusecase.md)                    |
|   massloadusecase   |                  [/usecases/massloadusecase](usecases/massloadusecase.md)                  |
|   parsingusecase    |                   [/usecases/parsingusecase](usecases/parsingusecase.md)                   |
|  rebuilderusecase   |                 [/usecases/rebuilderusecase](usecases/rebuilderusecase.md)                 |

## Scheme

```mermaid
erDiagram
    "/adapters/agent/internal/adapter" ||--|{ "/pkg" : x6
    "/adapters/metric/generator/generatorcore" ||--|{ "/pkg" : x1
    "/adapters/postgresql/internal/deadhash" ||--|{ "/pkg" : x1
    "/adapters/postgresql/internal/model" ||--|{ "/pkg" : x1
    "/adapters/postgresql/internal/repository" ||--|{ "/pkg" : x1
    "/adapters/tmpdata" ||--|{ "/pkg" : x1
    "/application/server" ||--|{ "/pkg" : x1
    "/controllers/apiagent" ||--|{ "/pkg" : x3
    "/controllers/apiserver/agenthandlers" ||--|{ "/pkg" : x1
    "/controllers/apiserver/apiservercore" ||--|{ "/pkg" : x1
    "/controllers/apiserver/attributehandlers" ||--|{ "/pkg" : x4
    "/controllers/apiserver/bookhandlers" ||--|{ "/pkg" : x2
    "/controllers/apiserver/deduplicatehandlers" ||--|{ "/pkg" : x5
    "/controllers/apiserver/fshandlers" ||--|{ "/pkg" : x1
    "/controllers/apiserver/hproxyhandlers" ||--|{ "/pkg" : x2
    "/controllers/apiserver/labelhandlers" ||--|{ "/pkg" : x2
    "/controllers/apiserver/massloadhandlers" ||--|{ "/pkg" : x3
    "/controllers/apiserver/systemhandlers" ||--|{ "/pkg" : x4
    "/controllers/internal/worker" ||--|{ "/pkg" : x1
    "/controllers/workermanager" ||--|{ "/pkg" : x9
    "/domain/bff" ||--|{ "/pkg" : x1
    "/domain/core" ||--|{ "/pkg" : x1
    "/external" ||--|{ "/pkg" : x2
    "/usecases/agentusecase" ||--|{ "/pkg" : x1
    "/usecases/bffusecase" ||--|{ "/pkg" : x1
    "/usecases/cleanupusecase" ||--|{ "/pkg" : x2
    "/usecases/deduplicatorusecase" ||--|{ "/pkg" : x1
    "/usecases/exportusecase" ||--|{ "/pkg" : x1
    "/usecases/filesystemusecase" ||--|{ "/pkg" : x1
    "/usecases/hproxyusecase" ||--|{ "/pkg" : x1
    "/usecases/massloadusecase" ||--|{ "/pkg" : x2
    "/usecases/parsingusecase" ||--|{ "/pkg" : x4
    "/usecases/rebuilderusecase" ||--|{ "/pkg" : x1
```

---

> Generated by [goArchLint](https://github.com/gbh007/goarchlint)
