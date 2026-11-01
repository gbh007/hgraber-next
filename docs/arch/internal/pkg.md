# pkg

## Imports

### External imports

| Name    | Path                                | Count |
|:--------|:------------------------------------|------:|
| assert  | github.com/stretchr/testify/assert  |     2 |
| require | github.com/stretchr/testify/require |     1 |

### Std imports

| Name    | Path         | Count |
|:--------|:-------------|------:|
| fmt     | fmt          |     2 |
| testing | testing      |     2 |
| bytes   | bytes        |     1 |
| sql     | database/sql |     1 |
| errors  | errors       |     1 |
| io      | io           |     1 |
| sync    | sync         |     1 |

## Used by

| Name                | Path                                                                                                      |
|:--------------------|:----------------------------------------------------------------------------------------------------------|
| agent               | [/internal/application/agent](application/agent.md)                                                       |
| agentapi            | [/internal/controllers/agentapi](controllers/agentapi.md)                                                 |
| agentcacheapi       | [/internal/controllers/agentcacheapi](controllers/agentcacheapi.md)                                       |
| agentdebugserver    | [/internal/controllers/agentdebugserver](controllers/agentdebugserver.md)                                 |
| agenthandlers       | [/internal/controllers/serverapi/agenthandlers](controllers/serverapi/agenthandlers.md)                   |
| apiservercore       | [/internal/controllers/serverapi/apiservercore](controllers/serverapi/apiservercore.md)                   |
| attributehandlers   | [/internal/controllers/serverapi/attributehandlers](controllers/serverapi/attributehandlers.md)           |
| bookhandlers        | [/internal/controllers/serverapi/bookhandlers](controllers/serverapi/bookhandlers.md)                     |
| deduplicatehandlers | [/internal/controllers/serverapi/deduplicatehandlers](controllers/serverapi/deduplicatehandlers.md)       |
| fshandlers          | [/internal/controllers/serverapi/fshandlers](controllers/serverapi/fshandlers.md)                         |
| hproxyhandlers      | [/internal/controllers/serverapi/hproxyhandlers](controllers/serverapi/hproxyhandlers.md)                 |
| labelhandlers       | [/internal/controllers/serverapi/labelhandlers](controllers/serverapi/labelhandlers.md)                   |
| massloadhandlers    | [/internal/controllers/serverapi/massloadhandlers](controllers/serverapi/massloadhandlers.md)             |
| systemhandlers      | [/internal/controllers/serverapi/systemhandlers](controllers/serverapi/systemhandlers.md)                 |
| workermanager       | [/internal/controllers/workermanager](controllers/workermanager.md)                                       |
| worker              | [/internal/controllers/workermanager/internal/worker](controllers/workermanager/internal/worker.md)       |
| adapter             | [/internal/dataproviders/agentapi/internal/adapter](dataproviders/agentapi/internal/adapter.md)           |
| agentimportfs       | [/internal/dataproviders/agentimportfs](dataproviders/agentimportfs.md)                                   |
| masterapi           | [/internal/dataproviders/masterapi](dataproviders/masterapi.md)                                           |
| generatorcore       | [/internal/dataproviders/metric/generator/generatorcore](dataproviders/metric/generator/generatorcore.md) |
| parserloader        | [/internal/dataproviders/parserloader](dataproviders/parserloader.md)                                     |
| deadhash            | [/internal/dataproviders/serverdb/internal/deadhash](dataproviders/serverdb/internal/deadhash.md)         |
| model               | [/internal/dataproviders/serverdb/internal/model](dataproviders/serverdb/internal/model.md)               |
| repository          | [/internal/dataproviders/serverdb/internal/repository](dataproviders/serverdb/internal/repository.md)     |
| tmpdata             | [/internal/dataproviders/tmpdata](dataproviders/tmpdata.md)                                               |
| archivemodel        | [/internal/domain/archivemodel](domain/archivemodel.md)                                                   |
| bff                 | [/internal/domain/bff](domain/bff.md)                                                                     |
| core                | [/internal/domain/core](domain/core.md)                                                                   |
| mcp                 | [/internal/experimental/mcp](experimental/mcp.md)                                                         |
| agentmanagesuc      | [/internal/usecases/agentmanagesuc](usecases/agentmanagesuc.md)                                           |
| bffusecase          | [/internal/usecases/bffusecase](usecases/bffusecase.md)                                                   |
| cleanupusecase      | [/internal/usecases/cleanupusecase](usecases/cleanupusecase.md)                                           |
| deduplicatorusecase | [/internal/usecases/deduplicatorusecase](usecases/deduplicatorusecase.md)                                 |
| exportusecase       | [/internal/usecases/exportusecase](usecases/exportusecase.md)                                             |
| filesystemusecase   | [/internal/usecases/filesystemusecase](usecases/filesystemusecase.md)                                     |
| hproxyusecase       | [/internal/usecases/hproxyusecase](usecases/hproxyusecase.md)                                             |
| massloadusecase     | [/internal/usecases/massloadusecase](usecases/massloadusecase.md)                                         |
| parsingusecase      | [/internal/usecases/parsingusecase](usecases/parsingusecase.md)                                           |
| rebuilderusecase    | [/internal/usecases/rebuilderusecase](usecases/rebuilderusecase.md)                                       |

## Scheme

```mermaid
erDiagram
    "/internal/application/agent" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/agentapi" ||--|{ "/internal/pkg" : x4
    "/internal/controllers/agentcacheapi" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/agentdebugserver" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi/agenthandlers" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi/apiservercore" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi/attributehandlers" ||--|{ "/internal/pkg" : x4
    "/internal/controllers/serverapi/bookhandlers" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/serverapi/deduplicatehandlers" ||--|{ "/internal/pkg" : x5
    "/internal/controllers/serverapi/fshandlers" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi/hproxyhandlers" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/serverapi/labelhandlers" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/serverapi/massloadhandlers" ||--|{ "/internal/pkg" : x3
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/internal/pkg" : x4
    "/internal/controllers/workermanager" ||--|{ "/internal/pkg" : x9
    "/internal/controllers/workermanager/internal/worker" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/agentapi/internal/adapter" ||--|{ "/internal/pkg" : x3
    "/internal/dataproviders/agentimportfs" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/masterapi" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/metric/generator/generatorcore" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/parserloader" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/serverdb/internal/deadhash" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/serverdb/internal/model" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/serverdb/internal/repository" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/tmpdata" ||--|{ "/internal/pkg" : x1
    "/internal/domain/archivemodel" ||--|{ "/internal/pkg" : x2
    "/internal/domain/bff" ||--|{ "/internal/pkg" : x1
    "/internal/domain/core" ||--|{ "/internal/pkg" : x1
    "/internal/experimental/mcp" ||--|{ "/internal/pkg" : x3
    "/internal/usecases/agentmanagesuc" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/bffusecase" ||--|{ "/internal/pkg" : x2
    "/internal/usecases/cleanupusecase" ||--|{ "/internal/pkg" : x2
    "/internal/usecases/deduplicatorusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/exportusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/filesystemusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/massloadusecase" ||--|{ "/internal/pkg" : x2
    "/internal/usecases/parsingusecase" ||--|{ "/internal/pkg" : x4
    "/internal/usecases/rebuilderusecase" ||--|{ "/internal/pkg" : x1
```

---

> Generated by [goArchLint](https://github.com/gbh007/goarchlint)
