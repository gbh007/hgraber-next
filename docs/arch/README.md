# github.com/gbh007/hgraber-next

## Main packages

| Name             | Path                                             | Description |
|:-----------------|:-------------------------------------------------|:------------|
| agent            | [/cmd/agent](cmd/agent.md)                       |             |
| grafanagenerator | [/cmd/grafanagenerator](cmd/grafanagenerator.md) |             |
| server           | [/cmd/server](cmd/server.md)                     |             |

## Inner packages

| Name                      | Path                                                                                                                       | Description |
|:--------------------------|:---------------------------------------------------------------------------------------------------------------------------|:------------|
| agentsdk                  | [/external/agentsdk](external/agentsdk.md)                                                                                 |             |
| config                    | [/external/config](external/config.md)                                                                                     |             |
| parserdto                 | [/external/parserdto](external/parserdto.md)                                                                               |             |
| legacyadapter             | [/external/parserdto/legacyadapter](external/parserdto/legacyadapter.md)                                                   |             |
| agent                     | [/internal/application/agent](internal/application/agent.md)                                                               |             |
| agnconfig                 | [/internal/application/agent/agnconfig](internal/application/agent/agnconfig.md)                                           |             |
| server                    | [/internal/application/server](internal/application/server.md)                                                             |             |
| srvconfig                 | [/internal/application/server/srvconfig](internal/application/server/srvconfig.md)                                         |             |
| agentapi                  | [/internal/controllers/agentapi](internal/controllers/agentapi.md)                                                         |             |
| agentcacheapi             | [/internal/controllers/agentcacheapi](internal/controllers/agentcacheapi.md)                                               |             |
| agentdebugserver          | [/internal/controllers/agentdebugserver](internal/controllers/agentdebugserver.md)                                         |             |
| async                     | [/internal/controllers/async](internal/controllers/async.md)                                                               |             |
| serverapi                 | [/internal/controllers/serverapi](internal/controllers/serverapi.md)                                                       |             |
| agenthandlers             | [/internal/controllers/serverapi/agenthandlers](internal/controllers/serverapi/agenthandlers.md)                           |             |
| apiservercore             | [/internal/controllers/serverapi/apiservercore](internal/controllers/serverapi/apiservercore.md)                           |             |
| attributehandlers         | [/internal/controllers/serverapi/attributehandlers](internal/controllers/serverapi/attributehandlers.md)                   |             |
| bookhandlers              | [/internal/controllers/serverapi/bookhandlers](internal/controllers/serverapi/bookhandlers.md)                             |             |
| deduplicatehandlers       | [/internal/controllers/serverapi/deduplicatehandlers](internal/controllers/serverapi/deduplicatehandlers.md)               |             |
| fshandlers                | [/internal/controllers/serverapi/fshandlers](internal/controllers/serverapi/fshandlers.md)                                 |             |
| hproxyhandlers            | [/internal/controllers/serverapi/hproxyhandlers](internal/controllers/serverapi/hproxyhandlers.md)                         |             |
| labelhandlers             | [/internal/controllers/serverapi/labelhandlers](internal/controllers/serverapi/labelhandlers.md)                           |             |
| massloadhandlers          | [/internal/controllers/serverapi/massloadhandlers](internal/controllers/serverapi/massloadhandlers.md)                     |             |
| systemhandlers            | [/internal/controllers/serverapi/systemhandlers](internal/controllers/serverapi/systemhandlers.md)                         |             |
| workermanager             | [/internal/controllers/workermanager](internal/controllers/workermanager.md)                                               |             |
| worker                    | [/internal/controllers/workermanager/internal/worker](internal/controllers/workermanager/internal/worker.md)               |             |
| agentapi                  | [/internal/dataproviders/agentapi](internal/dataproviders/agentapi.md)                                                     |             |
| adapter                   | [/internal/dataproviders/agentapi/internal/adapter](internal/dataproviders/agentapi/internal/adapter.md)                   |             |
| agentdb                   | [/internal/dataproviders/agentdb](internal/dataproviders/agentdb.md)                                                       |             |
| agentfs                   | [/internal/dataproviders/agentfs](internal/dataproviders/agentfs.md)                                                       |             |
| agentimportfs             | [/internal/dataproviders/agentimportfs](internal/dataproviders/agentimportfs.md)                                           |             |
| agentparsecache           | [/internal/dataproviders/agentparsecache](internal/dataproviders/agentparsecache.md)                                       |             |
| filestorage               | [/internal/dataproviders/filestorage](internal/dataproviders/filestorage.md)                                               |             |
| localfs                   | [/internal/dataproviders/localfs](internal/dataproviders/localfs.md)                                                       |             |
| masterapi                 | [/internal/dataproviders/masterapi](internal/dataproviders/masterapi.md)                                                   |             |
| metric                    | [/internal/dataproviders/metric](internal/dataproviders/metric.md)                                                         |             |
| generator                 | [/internal/dataproviders/metric/generator](internal/dataproviders/metric/generator.md)                                     |             |
| bookandpages              | [/internal/dataproviders/metric/generator/bookandpages](internal/dataproviders/metric/generator/bookandpages.md)           |             |
| databasepanel             | [/internal/dataproviders/metric/generator/databasepanel](internal/dataproviders/metric/generator/databasepanel.md)         |             |
| generatorcore             | [/internal/dataproviders/metric/generator/generatorcore](internal/dataproviders/metric/generator/generatorcore.md)         |             |
| httpserverpanel           | [/internal/dataproviders/metric/generator/httpserverpanel](internal/dataproviders/metric/generator/httpserverpanel.md)     |             |
| logspannel                | [/internal/dataproviders/metric/generator/logspannel](internal/dataproviders/metric/generator/logspannel.md)               |             |
| otherpanel                | [/internal/dataproviders/metric/generator/otherpanel](internal/dataproviders/metric/generator/otherpanel.md)               |             |
| simpleinfo                | [/internal/dataproviders/metric/generator/simpleinfo](internal/dataproviders/metric/generator/simpleinfo.md)               |             |
| statistic                 | [/internal/dataproviders/metric/generator/statistic](internal/dataproviders/metric/generator/statistic.md)                 |             |
| workersandagents          | [/internal/dataproviders/metric/generator/workersandagents](internal/dataproviders/metric/generator/workersandagents.md)   |             |
| metricagent               | [/internal/dataproviders/metric/metricagent](internal/dataproviders/metric/metricagent.md)                                 |             |
| metriccore                | [/internal/dataproviders/metric/metriccore](internal/dataproviders/metric/metriccore.md)                                   |             |
| metricdatabase            | [/internal/dataproviders/metric/metricdatabase](internal/dataproviders/metric/metricdatabase.md)                           |             |
| metricfs                  | [/internal/dataproviders/metric/metricfs](internal/dataproviders/metric/metricfs.md)                                       |             |
| metrichttp                | [/internal/dataproviders/metric/metrichttp](internal/dataproviders/metric/metrichttp.md)                                   |             |
| metricserver              | [/internal/dataproviders/metric/metricserver](internal/dataproviders/metric/metricserver.md)                               |             |
| metricstatistic           | [/internal/dataproviders/metric/metricstatistic](internal/dataproviders/metric/metricstatistic.md)                         |             |
| parserloader              | [/internal/dataproviders/parserloader](internal/dataproviders/parserloader.md)                                             |             |
| hgraber_local             | [/internal/dataproviders/parserloader/parsers/hgraber_local](internal/dataproviders/parserloader/parsers/hgraber_local.md) |             |
| mock                      | [/internal/dataproviders/parserloader/parsers/mock](internal/dataproviders/parserloader/parsers/mock.md)                   |             |
| request                   | [/internal/dataproviders/request](internal/dataproviders/request.md)                                                       |             |
| serverdb                  | [/internal/dataproviders/serverdb](internal/dataproviders/serverdb.md)                                                     |             |
| agent                     | [/internal/dataproviders/serverdb/internal/agent](internal/dataproviders/serverdb/internal/agent.md)                       |             |
| attribute                 | [/internal/dataproviders/serverdb/internal/attribute](internal/dataproviders/serverdb/internal/attribute.md)               |             |
| book                      | [/internal/dataproviders/serverdb/internal/book](internal/dataproviders/serverdb/internal/book.md)                         |             |
| deadhash                  | [/internal/dataproviders/serverdb/internal/deadhash](internal/dataproviders/serverdb/internal/deadhash.md)                 |             |
| file                      | [/internal/dataproviders/serverdb/internal/file](internal/dataproviders/serverdb/internal/file.md)                         |             |
| label                     | [/internal/dataproviders/serverdb/internal/label](internal/dataproviders/serverdb/internal/label.md)                       |             |
| massload                  | [/internal/dataproviders/serverdb/internal/massload](internal/dataproviders/serverdb/internal/massload.md)                 |             |
| model                     | [/internal/dataproviders/serverdb/internal/model](internal/dataproviders/serverdb/internal/model.md)                       |             |
| other                     | [/internal/dataproviders/serverdb/internal/other](internal/dataproviders/serverdb/internal/other.md)                       |             |
| page                      | [/internal/dataproviders/serverdb/internal/page](internal/dataproviders/serverdb/internal/page.md)                         |             |
| repository                | [/internal/dataproviders/serverdb/internal/repository](internal/dataproviders/serverdb/internal/repository.md)             |             |
| urlmirror                 | [/internal/dataproviders/serverdb/internal/urlmirror](internal/dataproviders/serverdb/internal/urlmirror.md)               |             |
| tmpdata                   | [/internal/dataproviders/tmpdata](internal/dataproviders/tmpdata.md)                                                       |             |
| agentdto                  | [/internal/domain/agentdto](internal/domain/agentdto.md)                                                                   |             |
| archivemodel              | [/internal/domain/archivemodel](internal/domain/archivemodel.md)                                                           |             |
| bff                       | [/internal/domain/bff](internal/domain/bff.md)                                                                             |             |
| core                      | [/internal/domain/core](internal/domain/core.md)                                                                           |             |
| fsmodel                   | [/internal/domain/fsmodel](internal/domain/fsmodel.md)                                                                     |             |
| hproxymodel               | [/internal/domain/hproxymodel](internal/domain/hproxymodel.md)                                                             |             |
| massloadmodel             | [/internal/domain/massloadmodel](internal/domain/massloadmodel.md)                                                         |             |
| parsing                   | [/internal/domain/parsing](internal/domain/parsing.md)                                                                     |             |
| serveragentdto            | [/internal/domain/serveragentdto](internal/domain/serveragentdto.md)                                                       |             |
| systemmodel               | [/internal/domain/systemmodel](internal/domain/systemmodel.md)                                                             |             |
| mcp                       | [/internal/experimental/mcp](internal/experimental/mcp.md)                                                                 |             |
| pkg                       | [/internal/pkg](internal/pkg.md)                                                                                           |             |
| agenthighway              | [/internal/usecases/agenthighway](internal/usecases/agenthighway.md)                                                       |             |
| agentimportapiuc          | [/internal/usecases/agentimportapiuc](internal/usecases/agentimportapiuc.md)                                               |             |
| agentimportdeduplicatoruc | [/internal/usecases/agentimportdeduplicatoruc](internal/usecases/agentimportdeduplicatoruc.md)                             |             |
| agentmanagesuc            | [/internal/usecases/agentmanagesuc](internal/usecases/agentmanagesuc.md)                                                   |             |
| agentparseruc             | [/internal/usecases/agentparseruc](internal/usecases/agentparseruc.md)                                                     |             |
| attributeusecase          | [/internal/usecases/attributeusecase](internal/usecases/attributeusecase.md)                                               |             |
| bffusecase                | [/internal/usecases/bffusecase](internal/usecases/bffusecase.md)                                                           |             |
| bookusecase               | [/internal/usecases/bookusecase](internal/usecases/bookusecase.md)                                                         |             |
| cleanupusecase            | [/internal/usecases/cleanupusecase](internal/usecases/cleanupusecase.md)                                                   |             |
| deduplicatorusecase       | [/internal/usecases/deduplicatorusecase](internal/usecases/deduplicatorusecase.md)                                         |             |
| exportusecase             | [/internal/usecases/exportusecase](internal/usecases/exportusecase.md)                                                     |             |
| filesystemusecase         | [/internal/usecases/filesystemusecase](internal/usecases/filesystemusecase.md)                                             |             |
| hproxyusecase             | [/internal/usecases/hproxyusecase](internal/usecases/hproxyusecase.md)                                                     |             |
| labelusecase              | [/internal/usecases/labelusecase](internal/usecases/labelusecase.md)                                                       |             |
| massloadusecase           | [/internal/usecases/massloadusecase](internal/usecases/massloadusecase.md)                                                 |             |
| parsingusecase            | [/internal/usecases/parsingusecase](internal/usecases/parsingusecase.md)                                                   |             |
| rebuilderusecase          | [/internal/usecases/rebuilderusecase](internal/usecases/rebuilderusecase.md)                                               |             |
| systemusecase             | [/internal/usecases/systemusecase](internal/usecases/systemusecase.md)                                                     |             |
| version                   | [/internal/version](internal/version.md)                                                                                   |             |
| agentapi                  | [/openapi/agentapi](openapi/agentapi.md)                                                                                   |             |
| serverapi                 | [/openapi/serverapi](openapi/serverapi.md)                                                                                 |             |

## External imports

| Name       | Path                                                      | Count |
|:-----------|:----------------------------------------------------------|------:|
| uuid       | github.com/google/uuid                                    |   176 |
| squirrel   | github.com/Masterminds/squirrel                           |   127 |
| v5         | github.com/jackc/pgx/v5                                   |    31 |
| timeseries | github.com/grafana/grafana-foundation-sdk/go/timeseries   |    26 |
| ogenerrors | github.com/ogen-go/ogen/ogenerrors                        |    21 |
| errors     | github.com/go-faster/errors                               |    20 |
| prometheus | github.com/prometheus/client_golang/prometheus            |    18 |
| dashboard  | github.com/grafana/grafana-foundation-sdk/go/dashboard    |    17 |
| promql     | github.com/grafana/promql-builder/go/promql               |    17 |
| validate   | github.com/ogen-go/ogen/validate                          |    16 |
| cog        | github.com/grafana/grafana-foundation-sdk/go/cog          |    15 |
| variants   | github.com/grafana/grafana-foundation-sdk/go/cog/variants |    15 |
| prometheus | github.com/grafana/grafana-foundation-sdk/go/prometheus   |    14 |
| http       | github.com/ogen-go/ogen/http                              |    12 |
| middleware | github.com/ogen-go/ogen/middleware                        |    11 |
| jx         | github.com/go-faster/jx                                   |    10 |
| uri        | github.com/ogen-go/ogen/uri                               |    10 |
| conv       | github.com/ogen-go/ogen/conv                              |     8 |
| server     | github.com/mark3labs/mcp-go/server                        |     7 |
| barchart   | github.com/grafana/grafana-foundation-sdk/go/barchart     |     6 |
| common     | github.com/grafana/grafana-foundation-sdk/go/common       |     6 |
| mcp        | github.com/mark3labs/mcp-go/mcp                           |     6 |
| lo         | github.com/samber/lo                                      |     6 |
| v4         | github.com/labstack/echo/v4                               |     5 |
| stat       | github.com/grafana/grafana-foundation-sdk/go/stat         |     4 |
| table      | github.com/grafana/grafana-foundation-sdk/go/table        |     3 |
| toml       | github.com/BurntSushi/toml                                |     2 |
| heatmap    | github.com/grafana/grafana-foundation-sdk/go/heatmap      |     2 |
| piechart   | github.com/grafana/grafana-foundation-sdk/go/piechart     |     2 |
| json       | github.com/ogen-go/ogen/json                              |     2 |
| v3         | github.com/pressly/goose/v3                               |     2 |
| promhttp   | github.com/prometheus/client_golang/prometheus/promhttp   |     2 |
| assert     | github.com/stretchr/testify/assert                        |     2 |
| go-sqlite  | github.com/glebarez/go-sqlite                             |     1 |
| strfmt     | github.com/go-openapi/strfmt                              |     1 |
| plugins    | github.com/grafana/grafana-foundation-sdk/go/cog/plugins  |     1 |
| logs       | github.com/grafana/grafana-foundation-sdk/go/logs         |     1 |
| loki       | github.com/grafana/grafana-foundation-sdk/go/loki         |     1 |
| units      | github.com/grafana/grafana-foundation-sdk/go/units        |     1 |
| client     | github.com/grafana/grafana-openapi-client-go/client       |     1 |
| models     | github.com/grafana/grafana-openapi-client-go/models       |     1 |
| cog        | github.com/grafana/promql-builder/go/cog                  |     1 |
| pgxpool    | github.com/jackc/pgx/v5/pgxpool                           |     1 |
| stdlib     | github.com/jackc/pgx/v5/stdlib                            |     1 |
| sqlx       | github.com/jmoiron/sqlx                                   |     1 |
| middleware | github.com/labstack/echo/v4/middleware                    |     1 |
| collectors | github.com/prometheus/client_golang/prometheus/collectors |     1 |
| v2         | github.com/qustavo/sqlhooks/v2                            |     1 |
| require    | github.com/stretchr/testify/require                       |     1 |
| proxy      | golang.org/x/net/proxy                                    |     1 |
| unix       | golang.org/x/sys/unix                                     |     1 |

## Std imports

| Name     | Path            | Count |
|:---------|:----------------|------:|
| context  | context         |   446 |
| fmt      | fmt             |   332 |
| time     | time            |   118 |
| slog     | log/slog        |   115 |
| url      | net/url         |    96 |
| errors   | errors          |    92 |
| http     | net/http        |    79 |
| io       | io              |    68 |
| sql      | database/sql    |    42 |
| strings  | strings         |    37 |
| bytes    | bytes           |    28 |
| os       | os              |    23 |
| slices   | slices          |    19 |
| strconv  | strconv         |    17 |
| path     | path            |    15 |
| json     | encoding/json   |     9 |
| sync     | sync            |     9 |
| zip      | archive/zip     |     7 |
| mime     | mime            |     7 |
| runtime  | runtime         |     5 |
| regexp   | regexp          |     4 |
| md5      | crypto/md5      |     3 |
| embed    | embed           |     3 |
| flag     | flag            |     3 |
| maps     | maps            |     3 |
| syscall  | syscall         |     3 |
| testing  | testing         |     3 |
| sha256   | crypto/sha256   |     2 |
| template | html/template   |     2 |
| bits     | math/bits       |     2 |
| signal   | os/signal       |     2 |
| atomic   | sync/atomic     |     2 |
| base64   | encoding/base64 |     1 |
| binary   | encoding/binary |     1 |
| hex      | encoding/hex    |     1 |
| html     | html            |     1 |
| math     | math            |     1 |
| debug    | runtime/debug   |     1 |

## Scheme

```mermaid
erDiagram
    "/cmd/agent" ||--|{ "/external/agentsdk" : x1
    "/cmd/agent" ||--|{ "/external/parserdto" : x1
    "/cmd/agent" ||--|{ "/internal/dataproviders/parserloader/parsers/hgraber_local" : x1
    "/cmd/agent" ||--|{ "/internal/dataproviders/parserloader/parsers/mock" : x1
    "/cmd/grafanagenerator" ||--|{ "/external/config" : x1
    "/cmd/grafanagenerator" ||--|{ "/internal/dataproviders/metric/generator" : x1
    "/cmd/server" ||--|{ "/internal/application/server" : x1
    "/external/agentsdk" ||--|{ "/internal/application/agent" : x1
    "/external/parserdto" ||--|{ "/internal/domain/core" : x1
    "/external/parserdto/legacyadapter" ||--|{ "/external/parserdto" : x2
    "/external/parserdto/legacyadapter" ||--|{ "/internal/domain/core" : x2
    "/internal/application/agent" ||--|{ "/external/config" : x1
    "/internal/application/agent" ||--|{ "/external/parserdto" : x1
    "/internal/application/agent" ||--|{ "/internal/application/agent/agnconfig" : x1
    "/internal/application/agent" ||--|{ "/internal/controllers/agentapi" : x1
    "/internal/application/agent" ||--|{ "/internal/controllers/agentdebugserver" : x1
    "/internal/application/agent" ||--|{ "/internal/controllers/async" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/agentdb" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/agentimportfs" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/agentparsecache" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/localfs" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/masterapi" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/metric" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/parserloader" : x1
    "/internal/application/agent" ||--|{ "/internal/dataproviders/request" : x1
    "/internal/application/agent" ||--|{ "/internal/domain/serveragentdto" : x1
    "/internal/application/agent" ||--|{ "/internal/pkg" : x1
    "/internal/application/agent" ||--|{ "/internal/usecases/agenthighway" : x1
    "/internal/application/agent" ||--|{ "/internal/usecases/agentimportapiuc" : x1
    "/internal/application/agent" ||--|{ "/internal/usecases/agentimportdeduplicatoruc" : x1
    "/internal/application/agent" ||--|{ "/internal/usecases/agentparseruc" : x1
    "/internal/application/server" ||--|{ "/external/config" : x1
    "/internal/application/server" ||--|{ "/internal/application/server/srvconfig" : x3
    "/internal/application/server" ||--|{ "/internal/controllers/agentcacheapi" : x1
    "/internal/application/server" ||--|{ "/internal/controllers/async" : x2
    "/internal/application/server" ||--|{ "/internal/controllers/serverapi" : x1
    "/internal/application/server" ||--|{ "/internal/controllers/workermanager" : x3
    "/internal/application/server" ||--|{ "/internal/dataproviders/agentapi" : x2
    "/internal/application/server" ||--|{ "/internal/dataproviders/filestorage" : x2
    "/internal/application/server" ||--|{ "/internal/dataproviders/metric" : x2
    "/internal/application/server" ||--|{ "/internal/dataproviders/serverdb" : x2
    "/internal/application/server" ||--|{ "/internal/dataproviders/tmpdata" : x2
    "/internal/application/server" ||--|{ "/internal/domain/core" : x1
    "/internal/application/server" ||--|{ "/internal/domain/systemmodel" : x1
    "/internal/application/server" ||--|{ "/internal/experimental/mcp" : x1
    "/internal/application/server" ||--|{ "/internal/usecases/agentmanagesuc" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/attributeusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/bffusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/bookusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/cleanupusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/deduplicatorusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/exportusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/filesystemusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/hproxyusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/labelusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/massloadusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/parsingusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/rebuilderusecase" : x2
    "/internal/application/server" ||--|{ "/internal/usecases/systemusecase" : x2
    "/internal/controllers/agentapi" ||--|{ "/external/parserdto" : x3
    "/internal/controllers/agentapi" ||--|{ "/internal/domain/agentdto" : x2
    "/internal/controllers/agentapi" ||--|{ "/internal/domain/core" : x4
    "/internal/controllers/agentapi" ||--|{ "/internal/domain/fsmodel" : x2
    "/internal/controllers/agentapi" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/controllers/agentapi" ||--|{ "/internal/pkg" : x4
    "/internal/controllers/agentapi" ||--|{ "/openapi/agentapi" : x16
    "/internal/controllers/agentcacheapi" ||--|{ "/internal/domain/core" : x2
    "/internal/controllers/agentcacheapi" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/controllers/agentcacheapi" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/agentcacheapi" ||--|{ "/openapi/agentapi" : x9
    "/internal/controllers/agentdebugserver" ||--|{ "/external/config" : x1
    "/internal/controllers/agentdebugserver" ||--|{ "/external/parserdto" : x2
    "/internal/controllers/agentdebugserver" ||--|{ "/internal/application/agent/agnconfig" : x1
    "/internal/controllers/agentdebugserver" ||--|{ "/internal/domain/archivemodel" : x1
    "/internal/controllers/agentdebugserver" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/agenthandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/apiservercore" : x3
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/attributehandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/bookhandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/deduplicatehandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/fshandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/hproxyhandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/labelhandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/massloadhandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/internal/controllers/serverapi/systemhandlers" : x1
    "/internal/controllers/serverapi" ||--|{ "/openapi/serverapi" : x4
    "/internal/controllers/serverapi/agenthandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x5
    "/internal/controllers/serverapi/agenthandlers" ||--|{ "/internal/domain/core" : x6
    "/internal/controllers/serverapi/agenthandlers" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/controllers/serverapi/agenthandlers" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi/agenthandlers" ||--|{ "/openapi/serverapi" : x6
    "/internal/controllers/serverapi/apiservercore" ||--|{ "/internal/domain/bff" : x1
    "/internal/controllers/serverapi/apiservercore" ||--|{ "/internal/domain/core" : x2
    "/internal/controllers/serverapi/apiservercore" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/controllers/serverapi/apiservercore" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi/apiservercore" ||--|{ "/openapi/serverapi" : x3
    "/internal/controllers/serverapi/attributehandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x5
    "/internal/controllers/serverapi/attributehandlers" ||--|{ "/internal/domain/core" : x9
    "/internal/controllers/serverapi/attributehandlers" ||--|{ "/internal/pkg" : x4
    "/internal/controllers/serverapi/attributehandlers" ||--|{ "/openapi/serverapi" : x12
    "/internal/controllers/serverapi/bookhandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x10
    "/internal/controllers/serverapi/bookhandlers" ||--|{ "/internal/domain/bff" : x3
    "/internal/controllers/serverapi/bookhandlers" ||--|{ "/internal/domain/core" : x10
    "/internal/controllers/serverapi/bookhandlers" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/serverapi/bookhandlers" ||--|{ "/openapi/serverapi" : x10
    "/internal/controllers/serverapi/deduplicatehandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x3
    "/internal/controllers/serverapi/deduplicatehandlers" ||--|{ "/internal/domain/bff" : x5
    "/internal/controllers/serverapi/deduplicatehandlers" ||--|{ "/internal/domain/core" : x4
    "/internal/controllers/serverapi/deduplicatehandlers" ||--|{ "/internal/pkg" : x5
    "/internal/controllers/serverapi/deduplicatehandlers" ||--|{ "/openapi/serverapi" : x6
    "/internal/controllers/serverapi/fshandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x7
    "/internal/controllers/serverapi/fshandlers" ||--|{ "/internal/domain/core" : x4
    "/internal/controllers/serverapi/fshandlers" ||--|{ "/internal/domain/fsmodel" : x4
    "/internal/controllers/serverapi/fshandlers" ||--|{ "/internal/pkg" : x1
    "/internal/controllers/serverapi/fshandlers" ||--|{ "/openapi/serverapi" : x11
    "/internal/controllers/serverapi/hproxyhandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x3
    "/internal/controllers/serverapi/hproxyhandlers" ||--|{ "/internal/domain/hproxymodel" : x3
    "/internal/controllers/serverapi/hproxyhandlers" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/serverapi/hproxyhandlers" ||--|{ "/openapi/serverapi" : x4
    "/internal/controllers/serverapi/labelhandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x3
    "/internal/controllers/serverapi/labelhandlers" ||--|{ "/internal/domain/core" : x7
    "/internal/controllers/serverapi/labelhandlers" ||--|{ "/internal/pkg" : x2
    "/internal/controllers/serverapi/labelhandlers" ||--|{ "/openapi/serverapi" : x8
    "/internal/controllers/serverapi/massloadhandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x4
    "/internal/controllers/serverapi/massloadhandlers" ||--|{ "/internal/domain/core" : x1
    "/internal/controllers/serverapi/massloadhandlers" ||--|{ "/internal/domain/massloadmodel" : x10
    "/internal/controllers/serverapi/massloadhandlers" ||--|{ "/internal/pkg" : x3
    "/internal/controllers/serverapi/massloadhandlers" ||--|{ "/openapi/serverapi" : x16
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/internal/controllers/serverapi/apiservercore" : x5
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/internal/domain/core" : x1
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/internal/domain/parsing" : x5
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/internal/domain/systemmodel" : x4
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/internal/pkg" : x4
    "/internal/controllers/serverapi/systemhandlers" ||--|{ "/openapi/serverapi" : x12
    "/internal/controllers/workermanager" ||--|{ "/internal/controllers/workermanager/internal/worker" : x10
    "/internal/controllers/workermanager" ||--|{ "/internal/domain/core" : x3
    "/internal/controllers/workermanager" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/controllers/workermanager" ||--|{ "/internal/domain/massloadmodel" : x2
    "/internal/controllers/workermanager" ||--|{ "/internal/domain/parsing" : x1
    "/internal/controllers/workermanager" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/controllers/workermanager" ||--|{ "/internal/domain/systemmodel" : x2
    "/internal/controllers/workermanager" ||--|{ "/internal/pkg" : x9
    "/internal/controllers/workermanager/internal/worker" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/agentapi" ||--|{ "/internal/dataproviders/agentapi/internal/adapter" : x1
    "/internal/dataproviders/agentapi" ||--|{ "/internal/domain/core" : x1
    "/internal/dataproviders/agentapi" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/dataproviders/agentapi" ||--|{ "/internal/domain/hproxymodel" : x1
    "/internal/dataproviders/agentapi" ||--|{ "/internal/domain/serveragentdto" : x1
    "/internal/dataproviders/agentapi/internal/adapter" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/dataproviders/agentapi/internal/adapter" ||--|{ "/internal/domain/hproxymodel" : x2
    "/internal/dataproviders/agentapi/internal/adapter" ||--|{ "/internal/domain/serveragentdto" : x6
    "/internal/dataproviders/agentapi/internal/adapter" ||--|{ "/internal/pkg" : x3
    "/internal/dataproviders/agentapi/internal/adapter" ||--|{ "/openapi/agentapi" : x10
    "/internal/dataproviders/agentdb" ||--|{ "/internal/domain/agentdto" : x1
    "/internal/dataproviders/agentfs" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/dataproviders/agentimportfs" ||--|{ "/internal/domain/agentdto" : x1
    "/internal/dataproviders/agentimportfs" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/filestorage" ||--|{ "/internal/dataproviders/agentfs" : x1
    "/internal/dataproviders/filestorage" ||--|{ "/internal/dataproviders/localfs" : x1
    "/internal/dataproviders/filestorage" ||--|{ "/internal/domain/core" : x3
    "/internal/dataproviders/filestorage" ||--|{ "/internal/domain/fsmodel" : x4
    "/internal/dataproviders/localfs" ||--|{ "/internal/domain/core" : x3
    "/internal/dataproviders/localfs" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/dataproviders/masterapi" ||--|{ "/internal/domain/agentdto" : x1
    "/internal/dataproviders/masterapi" ||--|{ "/internal/domain/serveragentdto" : x1
    "/internal/dataproviders/masterapi" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/masterapi" ||--|{ "/openapi/serverapi" : x3
    "/internal/dataproviders/metric" ||--|{ "/internal/dataproviders/metric/metricagent" : x2
    "/internal/dataproviders/metric" ||--|{ "/internal/dataproviders/metric/metriccore" : x3
    "/internal/dataproviders/metric" ||--|{ "/internal/dataproviders/metric/metricdatabase" : x2
    "/internal/dataproviders/metric" ||--|{ "/internal/dataproviders/metric/metricfs" : x3
    "/internal/dataproviders/metric" ||--|{ "/internal/dataproviders/metric/metrichttp" : x2
    "/internal/dataproviders/metric" ||--|{ "/internal/dataproviders/metric/metricserver" : x3
    "/internal/dataproviders/metric" ||--|{ "/internal/dataproviders/metric/metricstatistic" : x1
    "/internal/dataproviders/metric" ||--|{ "/internal/domain/core" : x1
    "/internal/dataproviders/metric" ||--|{ "/internal/domain/systemmodel" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/bookandpages" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/databasepanel" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x2
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/httpserverpanel" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/logspannel" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/otherpanel" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/simpleinfo" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/statistic" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/generator/workersandagents" : x1
    "/internal/dataproviders/metric/generator" ||--|{ "/internal/dataproviders/metric/metriccore" : x2
    "/internal/dataproviders/metric/generator/bookandpages" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x9
    "/internal/dataproviders/metric/generator/bookandpages" ||--|{ "/internal/dataproviders/metric/metriccore" : x8
    "/internal/dataproviders/metric/generator/bookandpages" ||--|{ "/internal/dataproviders/metric/metricfs" : x4
    "/internal/dataproviders/metric/generator/bookandpages" ||--|{ "/internal/dataproviders/metric/metricserver" : x4
    "/internal/dataproviders/metric/generator/databasepanel" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x6
    "/internal/dataproviders/metric/generator/databasepanel" ||--|{ "/internal/dataproviders/metric/metricdatabase" : x5
    "/internal/dataproviders/metric/generator/generatorcore" ||--|{ "/internal/dataproviders/metric/metriccore" : x1
    "/internal/dataproviders/metric/generator/generatorcore" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/metric/generator/httpserverpanel" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x5
    "/internal/dataproviders/metric/generator/httpserverpanel" ||--|{ "/internal/dataproviders/metric/metriccore" : x4
    "/internal/dataproviders/metric/generator/httpserverpanel" ||--|{ "/internal/dataproviders/metric/metrichttp" : x4
    "/internal/dataproviders/metric/generator/logspannel" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x2
    "/internal/dataproviders/metric/generator/otherpanel" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x5
    "/internal/dataproviders/metric/generator/otherpanel" ||--|{ "/internal/dataproviders/metric/metriccore" : x3
    "/internal/dataproviders/metric/generator/otherpanel" ||--|{ "/internal/dataproviders/metric/metricfs" : x3
    "/internal/dataproviders/metric/generator/otherpanel" ||--|{ "/internal/dataproviders/metric/metricserver" : x1
    "/internal/dataproviders/metric/generator/simpleinfo" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x13
    "/internal/dataproviders/metric/generator/simpleinfo" ||--|{ "/internal/dataproviders/metric/metricagent" : x1
    "/internal/dataproviders/metric/generator/simpleinfo" ||--|{ "/internal/dataproviders/metric/metriccore" : x12
    "/internal/dataproviders/metric/generator/simpleinfo" ||--|{ "/internal/dataproviders/metric/metricfs" : x5
    "/internal/dataproviders/metric/generator/simpleinfo" ||--|{ "/internal/dataproviders/metric/metricserver" : x6
    "/internal/dataproviders/metric/generator/statistic" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x2
    "/internal/dataproviders/metric/generator/statistic" ||--|{ "/internal/dataproviders/metric/metricstatistic" : x1
    "/internal/dataproviders/metric/generator/workersandagents" ||--|{ "/internal/dataproviders/metric/generator/generatorcore" : x7
    "/internal/dataproviders/metric/generator/workersandagents" ||--|{ "/internal/dataproviders/metric/metricagent" : x3
    "/internal/dataproviders/metric/generator/workersandagents" ||--|{ "/internal/dataproviders/metric/metriccore" : x3
    "/internal/dataproviders/metric/generator/workersandagents" ||--|{ "/internal/dataproviders/metric/metricserver" : x3
    "/internal/dataproviders/metric/metricagent" ||--|{ "/internal/dataproviders/metric/metriccore" : x2
    "/internal/dataproviders/metric/metriccore" ||--|{ "/internal/version" : x1
    "/internal/dataproviders/metric/metricdatabase" ||--|{ "/internal/dataproviders/metric/metriccore" : x1
    "/internal/dataproviders/metric/metricfs" ||--|{ "/internal/dataproviders/metric/metriccore" : x2
    "/internal/dataproviders/metric/metrichttp" ||--|{ "/internal/dataproviders/metric/metriccore" : x1
    "/internal/dataproviders/metric/metricserver" ||--|{ "/internal/dataproviders/metric/metriccore" : x2
    "/internal/dataproviders/metric/metricserver" ||--|{ "/internal/domain/systemmodel" : x1
    "/internal/dataproviders/metric/metricstatistic" ||--|{ "/internal/dataproviders/metric/metriccore" : x2
    "/internal/dataproviders/parserloader" ||--|{ "/external/parserdto" : x2
    "/internal/dataproviders/parserloader" ||--|{ "/internal/application/agent/agnconfig" : x1
    "/internal/dataproviders/parserloader" ||--|{ "/internal/dataproviders/request" : x1
    "/internal/dataproviders/parserloader" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/parserloader/parsers/hgraber_local" ||--|{ "/external/parserdto" : x1
    "/internal/dataproviders/parserloader/parsers/hgraber_local" ||--|{ "/external/parserdto/legacyadapter" : x1
    "/internal/dataproviders/parserloader/parsers/mock" ||--|{ "/external/parserdto" : x1
    "/internal/dataproviders/parserloader/parsers/mock" ||--|{ "/external/parserdto/legacyadapter" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/agent" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/attribute" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/book" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/deadhash" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/file" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/label" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/massload" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/other" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/page" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb" ||--|{ "/internal/dataproviders/serverdb/internal/urlmirror" : x1
    "/internal/dataproviders/serverdb/internal/agent" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x5
    "/internal/dataproviders/serverdb/internal/agent" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/agent" ||--|{ "/internal/domain/core" : x5
    "/internal/dataproviders/serverdb/internal/attribute" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x21
    "/internal/dataproviders/serverdb/internal/attribute" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/attribute" ||--|{ "/internal/domain/core" : x15
    "/internal/dataproviders/serverdb/internal/book" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x17
    "/internal/dataproviders/serverdb/internal/book" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/book" ||--|{ "/internal/domain/core" : x15
    "/internal/dataproviders/serverdb/internal/deadhash" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x5
    "/internal/dataproviders/serverdb/internal/deadhash" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/deadhash" ||--|{ "/internal/domain/core" : x5
    "/internal/dataproviders/serverdb/internal/deadhash" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/serverdb/internal/file" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x17
    "/internal/dataproviders/serverdb/internal/file" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/file" ||--|{ "/internal/domain/core" : x11
    "/internal/dataproviders/serverdb/internal/file" ||--|{ "/internal/domain/fsmodel" : x5
    "/internal/dataproviders/serverdb/internal/label" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x11
    "/internal/dataproviders/serverdb/internal/label" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/label" ||--|{ "/internal/domain/core" : x9
    "/internal/dataproviders/serverdb/internal/massload" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x24
    "/internal/dataproviders/serverdb/internal/massload" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/massload" ||--|{ "/internal/domain/massloadmodel" : x21
    "/internal/dataproviders/serverdb/internal/model" ||--|{ "/internal/domain/core" : x12
    "/internal/dataproviders/serverdb/internal/model" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/dataproviders/serverdb/internal/model" ||--|{ "/internal/domain/massloadmodel" : x4
    "/internal/dataproviders/serverdb/internal/model" ||--|{ "/internal/domain/parsing" : x1
    "/internal/dataproviders/serverdb/internal/model" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/serverdb/internal/other" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x1
    "/internal/dataproviders/serverdb/internal/other" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/other" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/dataproviders/serverdb/internal/other" ||--|{ "/internal/domain/systemmodel" : x1
    "/internal/dataproviders/serverdb/internal/page" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x22
    "/internal/dataproviders/serverdb/internal/page" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/page" ||--|{ "/internal/domain/core" : x17
    "/internal/dataproviders/serverdb/internal/repository" ||--|{ "/internal/pkg" : x1
    "/internal/dataproviders/serverdb/internal/urlmirror" ||--|{ "/internal/dataproviders/serverdb/internal/model" : x5
    "/internal/dataproviders/serverdb/internal/urlmirror" ||--|{ "/internal/dataproviders/serverdb/internal/repository" : x1
    "/internal/dataproviders/serverdb/internal/urlmirror" ||--|{ "/internal/domain/parsing" : x4
    "/internal/dataproviders/tmpdata" ||--|{ "/internal/domain/fsmodel" : x2
    "/internal/dataproviders/tmpdata" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/dataproviders/tmpdata" ||--|{ "/internal/domain/systemmodel" : x2
    "/internal/dataproviders/tmpdata" ||--|{ "/internal/pkg" : x1
    "/internal/domain/archivemodel" ||--|{ "/internal/domain/core" : x3
    "/internal/domain/archivemodel" ||--|{ "/internal/pkg" : x2
    "/internal/domain/bff" ||--|{ "/internal/domain/core" : x3
    "/internal/domain/bff" ||--|{ "/internal/pkg" : x1
    "/internal/domain/core" ||--|{ "/internal/pkg" : x1
    "/internal/domain/fsmodel" ||--|{ "/internal/domain/core" : x1
    "/internal/domain/parsing" ||--|{ "/internal/domain/core" : x1
    "/internal/domain/serveragentdto" ||--|{ "/internal/domain/core" : x2
    "/internal/experimental/mcp" ||--|{ "/internal/domain/bff" : x2
    "/internal/experimental/mcp" ||--|{ "/internal/domain/core" : x5
    "/internal/experimental/mcp" ||--|{ "/internal/domain/hproxymodel" : x1
    "/internal/experimental/mcp" ||--|{ "/internal/domain/parsing" : x2
    "/internal/experimental/mcp" ||--|{ "/internal/pkg" : x3
    "/internal/usecases/agentimportapiuc" ||--|{ "/internal/domain/agentdto" : x2
    "/internal/usecases/agentimportdeduplicatoruc" ||--|{ "/internal/domain/agentdto" : x2
    "/internal/usecases/agentmanagesuc" ||--|{ "/internal/domain/core" : x6
    "/internal/usecases/agentmanagesuc" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/usecases/agentmanagesuc" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/agentparseruc" ||--|{ "/external/parserdto" : x3
    "/internal/usecases/agentparseruc" ||--|{ "/internal/domain/serveragentdto" : x1
    "/internal/usecases/attributeusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/attributeusecase" ||--|{ "/internal/domain/systemmodel" : x1
    "/internal/usecases/bffusecase" ||--|{ "/internal/domain/bff" : x4
    "/internal/usecases/bffusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/bffusecase" ||--|{ "/internal/domain/fsmodel" : x1
    "/internal/usecases/bffusecase" ||--|{ "/internal/domain/massloadmodel" : x2
    "/internal/usecases/bffusecase" ||--|{ "/internal/pkg" : x2
    "/internal/usecases/bookusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/cleanupusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/cleanupusecase" ||--|{ "/internal/domain/fsmodel" : x3
    "/internal/usecases/cleanupusecase" ||--|{ "/internal/domain/systemmodel" : x7
    "/internal/usecases/cleanupusecase" ||--|{ "/internal/pkg" : x2
    "/internal/usecases/deduplicatorusecase" ||--|{ "/internal/domain/archivemodel" : x1
    "/internal/usecases/deduplicatorusecase" ||--|{ "/internal/domain/bff" : x5
    "/internal/usecases/deduplicatorusecase" ||--|{ "/internal/domain/core" : x12
    "/internal/usecases/deduplicatorusecase" ||--|{ "/internal/domain/systemmodel" : x1
    "/internal/usecases/deduplicatorusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/exportusecase" ||--|{ "/internal/domain/archivemodel" : x2
    "/internal/usecases/exportusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/exportusecase" ||--|{ "/internal/domain/parsing" : x2
    "/internal/usecases/exportusecase" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/usecases/exportusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/filesystemusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/filesystemusecase" ||--|{ "/internal/domain/fsmodel" : x4
    "/internal/usecases/filesystemusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/domain/core" : x5
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/domain/hproxymodel" : x4
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/domain/massloadmodel" : x2
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/domain/parsing" : x3
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/domain/serveragentdto" : x4
    "/internal/usecases/hproxyusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/labelusecase" ||--|{ "/internal/domain/core" : x2
    "/internal/usecases/labelusecase" ||--|{ "/internal/domain/systemmodel" : x1
    "/internal/usecases/massloadusecase" ||--|{ "/internal/domain/core" : x3
    "/internal/usecases/massloadusecase" ||--|{ "/internal/domain/hproxymodel" : x1
    "/internal/usecases/massloadusecase" ||--|{ "/internal/domain/massloadmodel" : x7
    "/internal/usecases/massloadusecase" ||--|{ "/internal/domain/parsing" : x2
    "/internal/usecases/massloadusecase" ||--|{ "/internal/domain/serveragentdto" : x2
    "/internal/usecases/massloadusecase" ||--|{ "/internal/domain/systemmodel" : x2
    "/internal/usecases/massloadusecase" ||--|{ "/internal/pkg" : x2
    "/internal/usecases/parsingusecase" ||--|{ "/internal/domain/core" : x9
    "/internal/usecases/parsingusecase" ||--|{ "/internal/domain/hproxymodel" : x2
    "/internal/usecases/parsingusecase" ||--|{ "/internal/domain/parsing" : x5
    "/internal/usecases/parsingusecase" ||--|{ "/internal/domain/serveragentdto" : x7
    "/internal/usecases/parsingusecase" ||--|{ "/internal/pkg" : x4
    "/internal/usecases/rebuilderusecase" ||--|{ "/internal/domain/core" : x10
    "/internal/usecases/rebuilderusecase" ||--|{ "/internal/pkg" : x1
    "/internal/usecases/systemusecase" ||--|{ "/internal/domain/systemmodel" : x3
```

---

> Generated by [goArchLint](https://github.com/gbh007/goarchlint)
