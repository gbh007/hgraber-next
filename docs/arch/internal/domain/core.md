# core

## Imports

### Inner imports

| Name | Path                       | Count |
|:-----|:---------------------------|------:|
| pkg  | [/internal/pkg](../pkg.md) |     1 |

### External imports

| Name | Path                   | Count |
|:-----|:-----------------------|------:|
| uuid | github.com/google/uuid |     8 |

### Std imports

| Name    | Path          | Count |
|:--------|:--------------|------:|
| time    | time          |     8 |
| url     | net/url       |     4 |
| strings | strings       |     3 |
| errors  | errors        |     2 |
| strconv | strconv       |     2 |
| md5     | crypto/md5    |     1 |
| sha256  | crypto/sha256 |     1 |
| fmt     | fmt           |     1 |
| io      | io            |     1 |
| maps    | maps          |     1 |

## Used by

| Name                | Path                                                                                                   |
|:--------------------|:-------------------------------------------------------------------------------------------------------|
| parserdto           | [/external/parserdto](../../external/parserdto.md)                                                     |
| legacyadapter       | [/external/parserdto/legacyadapter](../../external/parserdto/legacyadapter.md)                         |
| server              | [/internal/application/server](../application/server.md)                                               |
| agentapi            | [/internal/controllers/agentapi](../controllers/agentapi.md)                                           |
| agentcacheapi       | [/internal/controllers/agentcacheapi](../controllers/agentcacheapi.md)                                 |
| agenthandlers       | [/internal/controllers/serverapi/agenthandlers](../controllers/serverapi/agenthandlers.md)             |
| apiservercore       | [/internal/controllers/serverapi/apiservercore](../controllers/serverapi/apiservercore.md)             |
| attributehandlers   | [/internal/controllers/serverapi/attributehandlers](../controllers/serverapi/attributehandlers.md)     |
| bookhandlers        | [/internal/controllers/serverapi/bookhandlers](../controllers/serverapi/bookhandlers.md)               |
| deduplicatehandlers | [/internal/controllers/serverapi/deduplicatehandlers](../controllers/serverapi/deduplicatehandlers.md) |
| fshandlers          | [/internal/controllers/serverapi/fshandlers](../controllers/serverapi/fshandlers.md)                   |
| labelhandlers       | [/internal/controllers/serverapi/labelhandlers](../controllers/serverapi/labelhandlers.md)             |
| massloadhandlers    | [/internal/controllers/serverapi/massloadhandlers](../controllers/serverapi/massloadhandlers.md)       |
| systemhandlers      | [/internal/controllers/serverapi/systemhandlers](../controllers/serverapi/systemhandlers.md)           |
| workermanager       | [/internal/controllers/workermanager](../controllers/workermanager.md)                                 |
| agentapi            | [/internal/dataproviders/agentapi](../dataproviders/agentapi.md)                                       |
| filestorage         | [/internal/dataproviders/filestorage](../dataproviders/filestorage.md)                                 |
| localfs             | [/internal/dataproviders/localfs](../dataproviders/localfs.md)                                         |
| metric              | [/internal/dataproviders/metric](../dataproviders/metric.md)                                           |
| agent               | [/internal/dataproviders/serverdb/internal/agent](../dataproviders/serverdb/internal/agent.md)         |
| attribute           | [/internal/dataproviders/serverdb/internal/attribute](../dataproviders/serverdb/internal/attribute.md) |
| book                | [/internal/dataproviders/serverdb/internal/book](../dataproviders/serverdb/internal/book.md)           |
| deadhash            | [/internal/dataproviders/serverdb/internal/deadhash](../dataproviders/serverdb/internal/deadhash.md)   |
| file                | [/internal/dataproviders/serverdb/internal/file](../dataproviders/serverdb/internal/file.md)           |
| label               | [/internal/dataproviders/serverdb/internal/label](../dataproviders/serverdb/internal/label.md)         |
| model               | [/internal/dataproviders/serverdb/internal/model](../dataproviders/serverdb/internal/model.md)         |
| page                | [/internal/dataproviders/serverdb/internal/page](../dataproviders/serverdb/internal/page.md)           |
| archivemodel        | [/internal/domain/archivemodel](archivemodel.md)                                                       |
| bff                 | [/internal/domain/bff](bff.md)                                                                         |
| fsmodel             | [/internal/domain/fsmodel](fsmodel.md)                                                                 |
| parsing             | [/internal/domain/parsing](parsing.md)                                                                 |
| serveragentdto      | [/internal/domain/serveragentdto](serveragentdto.md)                                                   |
| mcp                 | [/internal/experimental/mcp](../experimental/mcp.md)                                                   |
| agentmanagesuc      | [/internal/usecases/agentmanagesuc](../usecases/agentmanagesuc.md)                                     |
| attributeusecase    | [/internal/usecases/attributeusecase](../usecases/attributeusecase.md)                                 |
| bffusecase          | [/internal/usecases/bffusecase](../usecases/bffusecase.md)                                             |
| bookusecase         | [/internal/usecases/bookusecase](../usecases/bookusecase.md)                                           |
| cleanupusecase      | [/internal/usecases/cleanupusecase](../usecases/cleanupusecase.md)                                     |
| deduplicatorusecase | [/internal/usecases/deduplicatorusecase](../usecases/deduplicatorusecase.md)                           |
| exportusecase       | [/internal/usecases/exportusecase](../usecases/exportusecase.md)                                       |
| filesystemusecase   | [/internal/usecases/filesystemusecase](../usecases/filesystemusecase.md)                               |
| hproxyusecase       | [/internal/usecases/hproxyusecase](../usecases/hproxyusecase.md)                                       |
| labelusecase        | [/internal/usecases/labelusecase](../usecases/labelusecase.md)                                         |
| massloadusecase     | [/internal/usecases/massloadusecase](../usecases/massloadusecase.md)                                   |
| parsingusecase      | [/internal/usecases/parsingusecase](../usecases/parsingusecase.md)                                     |
| rebuilderusecase    | [/internal/usecases/rebuilderusecase](../usecases/rebuilderusecase.md)                                 |

## Scheme

```mermaid
erDiagram
    "/external/parserdto" ||--|{ "/internal/domain/core" : x1
    "/external/parserdto/legacyadapter" ||--|{ "/internal/domain/core" : x2
    "/internal/application/server" ||--|{ "/internal/domain/core" : x1
    "/internal/controllers/agentapi" ||--|{ "/internal/domain/core" : x4
    "/internal/controllers/agentcacheapi" ||--|{ "/internal/domain/core" : x2
    "/internal/controllers/serverapi/agenthandlers" ||--|{ "/internal/domain/core" : x6
    "/internal/controllers/serverapi/apiservercore" ||--|{ "/internal/domain/core" : x2
    "/internal/controllers/serverapi/attributehandlers" ||--|{ "/internal/domain/core" : x9
    "/internal/controllers/serverapi/bookhandlers" ||--|{ "/internal/domain/core" : x10
    "/internal/controllers/serverapi/deduplicatehandlers" ||--|{ "/internal/domain/core" : x4
    "/internal/controllers/serverapi/fshandlers" ||--|{ "/internal/domain/core" : x4
    "/internal/controllers/serverapi/labelhandlers" ||--|{ "/internal/domain/core" : x7
    "/internal/controllers/serverapi/massloadhandlers" ||--|{ "/internal/domain/core" : x1
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/internal/domain/core" : x1
    "/internal/controllers/workermanager" ||--|{ "/internal/domain/core" : x3
    "/internal/dataproviders/agentapi" ||--|{ "/internal/domain/core" : x1
    "/internal/dataproviders/filestorage" ||--|{ "/internal/domain/core" : x3
    "/internal/dataproviders/localfs" ||--|{ "/internal/domain/core" : x3
    "/internal/dataproviders/metric" ||--|{ "/internal/domain/core" : x1
    "/internal/dataproviders/serverdb/internal/agent" ||--|{ "/internal/domain/core" : x5
    "/internal/dataproviders/serverdb/internal/attribute" ||--|{ "/internal/domain/core" : x15
    "/internal/dataproviders/serverdb/internal/book" ||--|{ "/internal/domain/core" : x15
    "/internal/dataproviders/serverdb/internal/deadhash" ||--|{ "/internal/domain/core" : x5
    "/internal/dataproviders/serverdb/internal/file" ||--|{ "/internal/domain/core" : x11
    "/internal/dataproviders/serverdb/internal/label" ||--|{ "/internal/domain/core" : x9
    "/internal/dataproviders/serverdb/internal/model" ||--|{ "/internal/domain/core" : x12
    "/internal/dataproviders/serverdb/internal/page" ||--|{ "/internal/domain/core" : x17
    "/internal/domain/archivemodel" ||--|{ "/internal/domain/core" : x3
    "/internal/domain/bff" ||--|{ "/internal/domain/core" : x3
    "/internal/domain/core" ||--|{ "/internal/pkg" : x1
    "/internal/domain/fsmodel" ||--|{ "/internal/domain/core" : x1
    "/internal/domain/parsing" ||--|{ "/internal/domain/core" : x1
    "/internal/domain/serveragentdto" ||--|{ "/internal/domain/core" : x2
    "/internal/experimental/mcp" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/agentmanagesuc" ||--|{ "/internal/domain/core" : x6
    "/internal/usecases/attributeusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/bffusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/bookusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/cleanupusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/deduplicatorusecase" ||--|{ "/internal/domain/core" : x12
    "/internal/usecases/exportusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/filesystemusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/labelusecase" ||--|{ "/internal/domain/core" : x2
    "/internal/usecases/massloadusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/parsingusecase" ||--|{ "/internal/domain/core" : x9
    "/internal/usecases/rebuilderusecase" ||--|{ "/internal/domain/core" : x10
```

---

> Generated by [goArchLint](https://github.com/gbh007/goarchlint)
