# github.com/gbh007/hgraber-next

## Main packages

|       Name       |                       Path                       |
|:----------------:|:------------------------------------------------:|
|  configremaper   |    [/cmd/configremaper](cmd/configremaper.md)    |
| grafanagenerator | [/cmd/grafanagenerator](cmd/grafanagenerator.md) |
|      server      |           [/cmd/server](cmd/server.md)           |

## Inner packages

|        Name         |                                             Path                                             |
|:-------------------:|:--------------------------------------------------------------------------------------------:|
|        agent        |                             [/adapters/agent](adapters/agent.md)                             |
|       adapter       |            [/adapters/agent/internal/adapter](adapters/agent/internal/adapter.md)            |
|       agentfs       |                           [/adapters/agentfs](adapters/agentfs.md)                           |
|     filestorage     |                       [/adapters/filestorage](adapters/filestorage.md)                       |
|     localfiles      |                        [/adapters/localfiles](adapters/localfiles.md)                        |
|       metric        |                            [/adapters/metric](adapters/metric.md)                            |
|      generator      |                  [/adapters/metric/generator](adapters/metric/generator.md)                  |
|    bookandpages     |     [/adapters/metric/generator/bookandpages](adapters/metric/generator/bookandpages.md)     |
|    databasepanel    |    [/adapters/metric/generator/databasepanel](adapters/metric/generator/databasepanel.md)    |
|    generatorcore    |    [/adapters/metric/generator/generatorcore](adapters/metric/generator/generatorcore.md)    |
|   httpserverpanel   |  [/adapters/metric/generator/httpserverpanel](adapters/metric/generator/httpserverpanel.md)  |
|     logspannel      |       [/adapters/metric/generator/logspannel](adapters/metric/generator/logspannel.md)       |
|     otherpanel      |       [/adapters/metric/generator/otherpanel](adapters/metric/generator/otherpanel.md)       |
|     simpleinfo      |       [/adapters/metric/generator/simpleinfo](adapters/metric/generator/simpleinfo.md)       |
|      statistic      |        [/adapters/metric/generator/statistic](adapters/metric/generator/statistic.md)        |
|  workersandagents   | [/adapters/metric/generator/workersandagents](adapters/metric/generator/workersandagents.md) |
|     metricagent     |                [/adapters/metric/metricagent](adapters/metric/metricagent.md)                |
|     metriccore      |                 [/adapters/metric/metriccore](adapters/metric/metriccore.md)                 |
|   metricdatabase    |             [/adapters/metric/metricdatabase](adapters/metric/metricdatabase.md)             |
|      metricfs       |                   [/adapters/metric/metricfs](adapters/metric/metricfs.md)                   |
|     metrichttp      |                 [/adapters/metric/metrichttp](adapters/metric/metrichttp.md)                 |
|    metricserver     |               [/adapters/metric/metricserver](adapters/metric/metricserver.md)               |
|   metricstatistic   |            [/adapters/metric/metricstatistic](adapters/metric/metricstatistic.md)            |
|     postgresql      |                        [/adapters/postgresql](adapters/postgresql.md)                        |
|        agent        |         [/adapters/postgresql/internal/agent](adapters/postgresql/internal/agent.md)         |
|      attribute      |     [/adapters/postgresql/internal/attribute](adapters/postgresql/internal/attribute.md)     |
|        book         |          [/adapters/postgresql/internal/book](adapters/postgresql/internal/book.md)          |
|      deadhash       |      [/adapters/postgresql/internal/deadhash](adapters/postgresql/internal/deadhash.md)      |
|        file         |          [/adapters/postgresql/internal/file](adapters/postgresql/internal/file.md)          |
|        label        |         [/adapters/postgresql/internal/label](adapters/postgresql/internal/label.md)         |
|      massload       |      [/adapters/postgresql/internal/massload](adapters/postgresql/internal/massload.md)      |
|        model        |         [/adapters/postgresql/internal/model](adapters/postgresql/internal/model.md)         |
|        other        |         [/adapters/postgresql/internal/other](adapters/postgresql/internal/other.md)         |
|        page         |          [/adapters/postgresql/internal/page](adapters/postgresql/internal/page.md)          |
|     repository      |    [/adapters/postgresql/internal/repository](adapters/postgresql/internal/repository.md)    |
|      urlmirror      |     [/adapters/postgresql/internal/urlmirror](adapters/postgresql/internal/urlmirror.md)     |
|       tmpdata       |                           [/adapters/tmpdata](adapters/tmpdata.md)                           |
|    configremaper    |                  [/application/configremaper](application/configremaper.md)                  |
|       server        |                         [/application/server](application/server.md)                         |
|    configremaper    |                          [/cmd/configremaper](cmd/configremaper.md)                          |
|  grafanagenerator   |                       [/cmd/grafanagenerator](cmd/grafanagenerator.md)                       |
|       server        |                                 [/cmd/server](cmd/server.md)                                 |
|       config        |                                     [/config](config.md)                                     |
|      apiagent       |                       [/controllers/apiagent](controllers/apiagent.md)                       |
|      apiserver      |                      [/controllers/apiserver](controllers/apiserver.md)                      |
|    agenthandlers    |        [/controllers/apiserver/agenthandlers](controllers/apiserver/agenthandlers.md)        |
|    apiservercore    |        [/controllers/apiserver/apiservercore](controllers/apiserver/apiservercore.md)        |
|  attributehandlers  |    [/controllers/apiserver/attributehandlers](controllers/apiserver/attributehandlers.md)    |
|    bookhandlers     |         [/controllers/apiserver/bookhandlers](controllers/apiserver/bookhandlers.md)         |
| deduplicatehandlers |  [/controllers/apiserver/deduplicatehandlers](controllers/apiserver/deduplicatehandlers.md)  |
|     fshandlers      |           [/controllers/apiserver/fshandlers](controllers/apiserver/fshandlers.md)           |
|   hproxyhandlers    |       [/controllers/apiserver/hproxyhandlers](controllers/apiserver/hproxyhandlers.md)       |
|    labelhandlers    |        [/controllers/apiserver/labelhandlers](controllers/apiserver/labelhandlers.md)        |
|  massloadhandlers   |     [/controllers/apiserver/massloadhandlers](controllers/apiserver/massloadhandlers.md)     |
|   systemhandlers    |       [/controllers/apiserver/systemhandlers](controllers/apiserver/systemhandlers.md)       |
|        async        |                          [/controllers/async](controllers/async.md)                          |
|       worker        |                [/controllers/internal/worker](controllers/internal/worker.md)                |
|    workermanager    |                  [/controllers/workermanager](controllers/workermanager.md)                  |
|     agentmodel      |                          [/domain/agentmodel](domain/agentmodel.md)                          |
|         bff         |                                 [/domain/bff](domain/bff.md)                                 |
|        core         |                                [/domain/core](domain/core.md)                                |
|       fsmodel       |                             [/domain/fsmodel](domain/fsmodel.md)                             |
|     hproxymodel     |                         [/domain/hproxymodel](domain/hproxymodel.md)                         |
|    massloadmodel    |                       [/domain/massloadmodel](domain/massloadmodel.md)                       |
|       parsing       |                             [/domain/parsing](domain/parsing.md)                             |
|     systemmodel     |                         [/domain/systemmodel](domain/systemmodel.md)                         |
|      external       |                                   [/external](external.md)                                   |
|      agentapi       |                           [/openapi/agentapi](openapi/agentapi.md)                           |
|      serverapi      |                          [/openapi/serverapi](openapi/serverapi.md)                          |
|         pkg         |                                        [/pkg](pkg.md)                                        |
|    agentusecase     |                      [/usecases/agentusecase](usecases/agentusecase.md)                      |
|  attributeusecase   |                  [/usecases/attributeusecase](usecases/attributeusecase.md)                  |
|     bffusecase      |                        [/usecases/bffusecase](usecases/bffusecase.md)                        |
|     bookusecase     |                       [/usecases/bookusecase](usecases/bookusecase.md)                       |
|   cleanupusecase    |                    [/usecases/cleanupusecase](usecases/cleanupusecase.md)                    |
| deduplicatorusecase |               [/usecases/deduplicatorusecase](usecases/deduplicatorusecase.md)               |
|    exportusecase    |                     [/usecases/exportusecase](usecases/exportusecase.md)                     |
|  filesystemusecase  |                 [/usecases/filesystemusecase](usecases/filesystemusecase.md)                 |
|    hproxyusecase    |                     [/usecases/hproxyusecase](usecases/hproxyusecase.md)                     |
|    labelusecase     |                      [/usecases/labelusecase](usecases/labelusecase.md)                      |
|   massloadusecase   |                   [/usecases/massloadusecase](usecases/massloadusecase.md)                   |
|   parsingusecase    |                    [/usecases/parsingusecase](usecases/parsingusecase.md)                    |
|  rebuilderusecase   |                  [/usecases/rebuilderusecase](usecases/rebuilderusecase.md)                  |
|    systemusecase    |                     [/usecases/systemusecase](usecases/systemusecase.md)                     |
|       version       |                                    [/version](version.md)                                    |

## External imports

|     Name      |                              Path                               | Count |
|:-------------:|:---------------------------------------------------------------:|:-----:|
|    context    |                             context                             |  395  |
|      fmt      |                               fmt                               |  307  |
|     uuid      |                     github.com/google/uuid                      |  164  |
|   squirrel    |                 github.com/Masterminds/squirrel                 |  111  |
|     time      |                              time                               |  93   |
|     slog      |                            log/slog                             |  90   |
|      url      |                             net/url                             |  69   |
|    errors     |                             errors                              |  64   |
|      io       |                               io                                |  45   |
|     trace     |                 go.opentelemetry.io/otel/trace                  |  43   |
|      sql      |                          database/sql                           |  36   |
|     http      |                            net/http                             |  30   |
|  timeseries   |     github.com/grafana/grafana-foundation-sdk/go/timeseries     |  26   |
|      v5       |                     github.com/jackc/pgx/v5                     |  24   |
|    strings    |                             strings                             |  22   |
|    errors     |                   github.com/go-faster/errors                   |  20   |
|   dashboard   |     github.com/grafana/grafana-foundation-sdk/go/dashboard      |  17   |
|    promql     |           github.com/grafana/promql-builder/go/promql           |  17   |
|  prometheus   |         github.com/prometheus/client_golang/prometheus          |  17   |
|  ogenerrors   |               github.com/ogen-go/ogen/ogenerrors                |  16   |
|    slices     |                             slices                              |  16   |
|      cog      |        github.com/grafana/grafana-foundation-sdk/go/cog         |  15   |
|   variants    |    github.com/grafana/grafana-foundation-sdk/go/cog/variants    |  15   |
|     bytes     |                              bytes                              |  14   |
|  prometheus   |     github.com/grafana/grafana-foundation-sdk/go/prometheus     |  14   |
|     http      |                  github.com/ogen-go/ogen/http                   |  12   |
|   validate    |                github.com/ogen-go/ogen/validate                 |  12   |
|      os       |                               os                                |  12   |
|      jx       |                     github.com/go-faster/jx                     |  10   |
|  middleware   |               github.com/ogen-go/ogen/middleware                |  10   |
|      uri      |                   github.com/ogen-go/ogen/uri                   |  10   |
|    strconv    |                             strconv                             |   9   |
|     conv      |                  github.com/ogen-go/ogen/conv                   |   8   |
|   attribute   |               go.opentelemetry.io/otel/attribute                |   8   |
|     path      |                              path                               |   8   |
|     sync      |                              sync                               |   8   |
|     otel      |                    go.opentelemetry.io/otel                     |   7   |
|      zip      |                           archive/zip                           |   6   |
|     json      |                          encoding/json                          |   6   |
|   barchart    |      github.com/grafana/grafana-foundation-sdk/go/barchart      |   6   |
|    common     |       github.com/grafana/grafana-foundation-sdk/go/common       |   6   |
|     codes     |                 go.opentelemetry.io/otel/codes                  |   6   |
|    metric     |                 go.opentelemetry.io/otel/metric                 |   6   |
|     mime      |                              mime                               |   6   |
|    runtime    |                             runtime                             |   5   |
|     stat      |        github.com/grafana/grafana-foundation-sdk/go/stat        |   4   |
|  propagation  |              go.opentelemetry.io/otel/propagation               |   4   |
|    v1.26.0    |            go.opentelemetry.io/otel/semconv/v1.26.0             |   4   |
|     flag      |                              flag                               |   3   |
|     table     |       github.com/grafana/grafana-foundation-sdk/go/table        |   3   |
| pyroscope-go  |                 github.com/grafana/pyroscope-go                 |   3   |
|     toml      |                   github.com/BurntSushi/toml                    |   2   |
|    heatmap    |      github.com/grafana/grafana-foundation-sdk/go/heatmap       |   2   |
|   piechart    |      github.com/grafana/grafana-foundation-sdk/go/piechart      |   2   |
|   envconfig   |              github.com/kelseyhightower/envconfig               |   2   |
|     json      |                  github.com/ogen-go/ogen/json                   |   2   |
|   otelogen    |                github.com/ogen-go/ogen/otelogen                 |   2   |
|    assert     |               github.com/stretchr/testify/assert                |   2   |
|    yaml.v3    |                        gopkg.in/yaml.v3                         |   2   |
|     bits      |                            math/bits                            |   2   |
|    atomic     |                           sync/atomic                           |   2   |
|    syscall    |                             syscall                             |   2   |
|    testing    |                             testing                             |   2   |
|      md5      |                           crypto/md5                            |   1   |
|    sha256     |                          crypto/sha256                          |   1   |
|     embed     |                              embed                              |   1   |
|    strfmt     |                  github.com/go-openapi/strfmt                   |   1   |
|    plugins    |    github.com/grafana/grafana-foundation-sdk/go/cog/plugins     |   1   |
|     logs      |        github.com/grafana/grafana-foundation-sdk/go/logs        |   1   |
|     loki      |        github.com/grafana/grafana-foundation-sdk/go/loki        |   1   |
|     units     |       github.com/grafana/grafana-foundation-sdk/go/units        |   1   |
|    client     |       github.com/grafana/grafana-openapi-client-go/client       |   1   |
|    models     |       github.com/grafana/grafana-openapi-client-go/models       |   1   |
|      cog      |            github.com/grafana/promql-builder/go/cog             |   1   |
|    pgxpool    |                 github.com/jackc/pgx/v5/pgxpool                 |   1   |
|    stdlib     |                 github.com/jackc/pgx/v5/stdlib                  |   1   |
|   godotenv    |                    github.com/joho/godotenv                     |   1   |
|      v3       |                   github.com/pressly/goose/v3                   |   1   |
|  collectors   |    github.com/prometheus/client_golang/prometheus/collectors    |   1   |
|   promhttp    |     github.com/prometheus/client_golang/prometheus/promhttp     |   1   |
|    require    |               github.com/stretchr/testify/require               |   1   |
| otlptracehttp | go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp |   1   |
|   resource    |              go.opentelemetry.io/otel/sdk/resource              |   1   |
|     trace     |               go.opentelemetry.io/otel/sdk/trace                |   1   |
|    v1.20.0    |            go.opentelemetry.io/otel/semconv/v1.20.0             |   1   |
|     noop      |               go.opentelemetry.io/otel/trace/noop               |   1   |
|     unix      |                      golang.org/x/sys/unix                      |   1   |
|     maps      |                              maps                               |   1   |
|     math      |                              math                               |   1   |
|    signal     |                            os/signal                            |   1   |
|    regexp     |                             regexp                              |   1   |
|   template    |                          text/template                          |   1   |

## Scheme

```mermaid
erDiagram
    "/adapters/agent" ||--|{ "/adapters/agent/internal/adapter" : x1
    "/adapters/agent" ||--|{ "/domain/agentmodel" : x1
    "/adapters/agent" ||--|{ "/domain/core" : x1
    "/adapters/agent" ||--|{ "/domain/fsmodel" : x1
    "/adapters/agent" ||--|{ "/domain/hproxymodel" : x1
    "/adapters/agent/internal/adapter" ||--|{ "/domain/agentmodel" : x10
    "/adapters/agent/internal/adapter" ||--|{ "/domain/core" : x1
    "/adapters/agent/internal/adapter" ||--|{ "/domain/fsmodel" : x1
    "/adapters/agent/internal/adapter" ||--|{ "/domain/hproxymodel" : x1
    "/adapters/agent/internal/adapter" ||--|{ "/openapi/agentapi" : x11
    "/adapters/agent/internal/adapter" ||--|{ "/pkg" : x6
    "/adapters/agentfs" ||--|{ "/domain/fsmodel" : x1
    "/adapters/filestorage" ||--|{ "/adapters/agentfs" : x1
    "/adapters/filestorage" ||--|{ "/adapters/localfiles" : x1
    "/adapters/filestorage" ||--|{ "/domain/core" : x3
    "/adapters/filestorage" ||--|{ "/domain/fsmodel" : x4
    "/adapters/localfiles" ||--|{ "/domain/core" : x2
    "/adapters/localfiles" ||--|{ "/domain/fsmodel" : x1
    "/adapters/metric" ||--|{ "/adapters/metric/metricagent" : x2
    "/adapters/metric" ||--|{ "/adapters/metric/metriccore" : x3
    "/adapters/metric" ||--|{ "/adapters/metric/metricdatabase" : x2
    "/adapters/metric" ||--|{ "/adapters/metric/metricfs" : x3
    "/adapters/metric" ||--|{ "/adapters/metric/metrichttp" : x2
    "/adapters/metric" ||--|{ "/adapters/metric/metricserver" : x3
    "/adapters/metric" ||--|{ "/adapters/metric/metricstatistic" : x1
    "/adapters/metric" ||--|{ "/domain/core" : x1
    "/adapters/metric" ||--|{ "/domain/systemmodel" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/bookandpages" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/databasepanel" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/generatorcore" : x2
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/httpserverpanel" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/logspannel" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/otherpanel" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/simpleinfo" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/statistic" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/generator/workersandagents" : x1
    "/adapters/metric/generator" ||--|{ "/adapters/metric/metriccore" : x2
    "/adapters/metric/generator/bookandpages" ||--|{ "/adapters/metric/generator/generatorcore" : x9
    "/adapters/metric/generator/bookandpages" ||--|{ "/adapters/metric/metriccore" : x8
    "/adapters/metric/generator/bookandpages" ||--|{ "/adapters/metric/metricfs" : x4
    "/adapters/metric/generator/bookandpages" ||--|{ "/adapters/metric/metricserver" : x4
    "/adapters/metric/generator/databasepanel" ||--|{ "/adapters/metric/generator/generatorcore" : x6
    "/adapters/metric/generator/databasepanel" ||--|{ "/adapters/metric/metricdatabase" : x5
    "/adapters/metric/generator/generatorcore" ||--|{ "/adapters/metric/metriccore" : x1
    "/adapters/metric/generator/generatorcore" ||--|{ "/pkg" : x1
    "/adapters/metric/generator/httpserverpanel" ||--|{ "/adapters/metric/generator/generatorcore" : x5
    "/adapters/metric/generator/httpserverpanel" ||--|{ "/adapters/metric/metriccore" : x4
    "/adapters/metric/generator/httpserverpanel" ||--|{ "/adapters/metric/metrichttp" : x4
    "/adapters/metric/generator/logspannel" ||--|{ "/adapters/metric/generator/generatorcore" : x2
    "/adapters/metric/generator/otherpanel" ||--|{ "/adapters/metric/generator/generatorcore" : x5
    "/adapters/metric/generator/otherpanel" ||--|{ "/adapters/metric/metriccore" : x3
    "/adapters/metric/generator/otherpanel" ||--|{ "/adapters/metric/metricfs" : x3
    "/adapters/metric/generator/otherpanel" ||--|{ "/adapters/metric/metricserver" : x1
    "/adapters/metric/generator/simpleinfo" ||--|{ "/adapters/metric/generator/generatorcore" : x13
    "/adapters/metric/generator/simpleinfo" ||--|{ "/adapters/metric/metricagent" : x1
    "/adapters/metric/generator/simpleinfo" ||--|{ "/adapters/metric/metriccore" : x12
    "/adapters/metric/generator/simpleinfo" ||--|{ "/adapters/metric/metricfs" : x5
    "/adapters/metric/generator/simpleinfo" ||--|{ "/adapters/metric/metricserver" : x6
    "/adapters/metric/generator/statistic" ||--|{ "/adapters/metric/generator/generatorcore" : x2
    "/adapters/metric/generator/statistic" ||--|{ "/adapters/metric/metricstatistic" : x1
    "/adapters/metric/generator/workersandagents" ||--|{ "/adapters/metric/generator/generatorcore" : x7
    "/adapters/metric/generator/workersandagents" ||--|{ "/adapters/metric/metricagent" : x3
    "/adapters/metric/generator/workersandagents" ||--|{ "/adapters/metric/metriccore" : x3
    "/adapters/metric/generator/workersandagents" ||--|{ "/adapters/metric/metricserver" : x3
    "/adapters/metric/metricagent" ||--|{ "/adapters/metric/metriccore" : x2
    "/adapters/metric/metriccore" ||--|{ "/version" : x1
    "/adapters/metric/metricdatabase" ||--|{ "/adapters/metric/metriccore" : x1
    "/adapters/metric/metricfs" ||--|{ "/adapters/metric/metriccore" : x2
    "/adapters/metric/metrichttp" ||--|{ "/adapters/metric/metriccore" : x1
    "/adapters/metric/metricserver" ||--|{ "/adapters/metric/metriccore" : x2
    "/adapters/metric/metricserver" ||--|{ "/domain/systemmodel" : x1
    "/adapters/metric/metricstatistic" ||--|{ "/adapters/metric/metriccore" : x2
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/agent" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/attribute" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/book" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/deadhash" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/file" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/label" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/massload" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/other" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/page" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql" ||--|{ "/adapters/postgresql/internal/urlmirror" : x1
    "/adapters/postgresql/internal/agent" ||--|{ "/adapters/postgresql/internal/model" : x2
    "/adapters/postgresql/internal/agent" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/agent" ||--|{ "/domain/core" : x5
    "/adapters/postgresql/internal/attribute" ||--|{ "/adapters/postgresql/internal/model" : x5
    "/adapters/postgresql/internal/attribute" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/attribute" ||--|{ "/domain/core" : x15
    "/adapters/postgresql/internal/book" ||--|{ "/adapters/postgresql/internal/model" : x7
    "/adapters/postgresql/internal/book" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/book" ||--|{ "/domain/core" : x15
    "/adapters/postgresql/internal/deadhash" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/deadhash" ||--|{ "/domain/core" : x5
    "/adapters/postgresql/internal/deadhash" ||--|{ "/pkg" : x1
    "/adapters/postgresql/internal/file" ||--|{ "/adapters/postgresql/internal/model" : x10
    "/adapters/postgresql/internal/file" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/file" ||--|{ "/domain/core" : x11
    "/adapters/postgresql/internal/file" ||--|{ "/domain/fsmodel" : x5
    "/adapters/postgresql/internal/label" ||--|{ "/adapters/postgresql/internal/model" : x3
    "/adapters/postgresql/internal/label" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/label" ||--|{ "/domain/core" : x9
    "/adapters/postgresql/internal/massload" ||--|{ "/adapters/postgresql/internal/model" : x16
    "/adapters/postgresql/internal/massload" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/massload" ||--|{ "/domain/massloadmodel" : x21
    "/adapters/postgresql/internal/model" ||--|{ "/domain/core" : x8
    "/adapters/postgresql/internal/model" ||--|{ "/domain/fsmodel" : x1
    "/adapters/postgresql/internal/model" ||--|{ "/domain/massloadmodel" : x1
    "/adapters/postgresql/internal/model" ||--|{ "/domain/parsing" : x1
    "/adapters/postgresql/internal/model" ||--|{ "/pkg" : x1
    "/adapters/postgresql/internal/other" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/other" ||--|{ "/domain/fsmodel" : x1
    "/adapters/postgresql/internal/other" ||--|{ "/domain/systemmodel" : x1
    "/adapters/postgresql/internal/page" ||--|{ "/adapters/postgresql/internal/model" : x11
    "/adapters/postgresql/internal/page" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/page" ||--|{ "/domain/core" : x17
    "/adapters/postgresql/internal/repository" ||--|{ "/pkg" : x1
    "/adapters/postgresql/internal/urlmirror" ||--|{ "/adapters/postgresql/internal/model" : x4
    "/adapters/postgresql/internal/urlmirror" ||--|{ "/adapters/postgresql/internal/repository" : x1
    "/adapters/postgresql/internal/urlmirror" ||--|{ "/domain/parsing" : x4
    "/adapters/tmpdata" ||--|{ "/domain/agentmodel" : x2
    "/adapters/tmpdata" ||--|{ "/domain/fsmodel" : x2
    "/adapters/tmpdata" ||--|{ "/domain/systemmodel" : x2
    "/adapters/tmpdata" ||--|{ "/pkg" : x1
    "/application/configremaper" ||--|{ "/config" : x1
    "/application/server" ||--|{ "/adapters/agent" : x2
    "/application/server" ||--|{ "/adapters/filestorage" : x2
    "/application/server" ||--|{ "/adapters/metric" : x2
    "/application/server" ||--|{ "/adapters/postgresql" : x2
    "/application/server" ||--|{ "/adapters/tmpdata" : x2
    "/application/server" ||--|{ "/config" : x4
    "/application/server" ||--|{ "/controllers/apiagent" : x1
    "/application/server" ||--|{ "/controllers/apiserver" : x1
    "/application/server" ||--|{ "/controllers/async" : x2
    "/application/server" ||--|{ "/controllers/workermanager" : x3
    "/application/server" ||--|{ "/domain/core" : x1
    "/application/server" ||--|{ "/domain/systemmodel" : x1
    "/application/server" ||--|{ "/pkg" : x1
    "/application/server" ||--|{ "/usecases/agentusecase" : x2
    "/application/server" ||--|{ "/usecases/attributeusecase" : x2
    "/application/server" ||--|{ "/usecases/bffusecase" : x2
    "/application/server" ||--|{ "/usecases/bookusecase" : x2
    "/application/server" ||--|{ "/usecases/cleanupusecase" : x2
    "/application/server" ||--|{ "/usecases/deduplicatorusecase" : x2
    "/application/server" ||--|{ "/usecases/exportusecase" : x2
    "/application/server" ||--|{ "/usecases/filesystemusecase" : x2
    "/application/server" ||--|{ "/usecases/hproxyusecase" : x2
    "/application/server" ||--|{ "/usecases/labelusecase" : x2
    "/application/server" ||--|{ "/usecases/massloadusecase" : x2
    "/application/server" ||--|{ "/usecases/parsingusecase" : x2
    "/application/server" ||--|{ "/usecases/rebuilderusecase" : x2
    "/application/server" ||--|{ "/usecases/systemusecase" : x2
    "/cmd/configremaper" ||--|{ "/application/configremaper" : x1
    "/cmd/configremaper" ||--|{ "/config" : x1
    "/cmd/grafanagenerator" ||--|{ "/adapters/metric/generator" : x1
    "/cmd/grafanagenerator" ||--|{ "/config" : x1
    "/cmd/server" ||--|{ "/application/server" : x1
    "/controllers/apiagent" ||--|{ "/domain/agentmodel" : x3
    "/controllers/apiagent" ||--|{ "/domain/core" : x2
    "/controllers/apiagent" ||--|{ "/openapi/agentapi" : x9
    "/controllers/apiagent" ||--|{ "/pkg" : x3
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/agenthandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/apiservercore" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/attributehandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/bookhandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/deduplicatehandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/fshandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/hproxyhandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/labelhandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/massloadhandlers" : x1
    "/controllers/apiserver" ||--|{ "/controllers/apiserver/systemhandlers" : x1
    "/controllers/apiserver" ||--|{ "/openapi/serverapi" : x3
    "/controllers/apiserver/agenthandlers" ||--|{ "/controllers/apiserver/apiservercore" : x7
    "/controllers/apiserver/agenthandlers" ||--|{ "/domain/agentmodel" : x2
    "/controllers/apiserver/agenthandlers" ||--|{ "/domain/core" : x6
    "/controllers/apiserver/agenthandlers" ||--|{ "/openapi/serverapi" : x6
    "/controllers/apiserver/agenthandlers" ||--|{ "/pkg" : x1
    "/controllers/apiserver/apiservercore" ||--|{ "/domain/bff" : x1
    "/controllers/apiserver/apiservercore" ||--|{ "/domain/core" : x2
    "/controllers/apiserver/apiservercore" ||--|{ "/domain/fsmodel" : x1
    "/controllers/apiserver/apiservercore" ||--|{ "/openapi/serverapi" : x3
    "/controllers/apiserver/apiservercore" ||--|{ "/pkg" : x1
    "/controllers/apiserver/attributehandlers" ||--|{ "/controllers/apiserver/apiservercore" : x13
    "/controllers/apiserver/attributehandlers" ||--|{ "/domain/core" : x9
    "/controllers/apiserver/attributehandlers" ||--|{ "/openapi/serverapi" : x12
    "/controllers/apiserver/attributehandlers" ||--|{ "/pkg" : x4
    "/controllers/apiserver/bookhandlers" ||--|{ "/controllers/apiserver/apiservercore" : x11
    "/controllers/apiserver/bookhandlers" ||--|{ "/domain/bff" : x3
    "/controllers/apiserver/bookhandlers" ||--|{ "/domain/core" : x10
    "/controllers/apiserver/bookhandlers" ||--|{ "/openapi/serverapi" : x10
    "/controllers/apiserver/bookhandlers" ||--|{ "/pkg" : x2
    "/controllers/apiserver/deduplicatehandlers" ||--|{ "/controllers/apiserver/apiservercore" : x7
    "/controllers/apiserver/deduplicatehandlers" ||--|{ "/domain/bff" : x5
    "/controllers/apiserver/deduplicatehandlers" ||--|{ "/domain/core" : x4
    "/controllers/apiserver/deduplicatehandlers" ||--|{ "/openapi/serverapi" : x6
    "/controllers/apiserver/deduplicatehandlers" ||--|{ "/pkg" : x5
    "/controllers/apiserver/fshandlers" ||--|{ "/controllers/apiserver/apiservercore" : x12
    "/controllers/apiserver/fshandlers" ||--|{ "/domain/core" : x4
    "/controllers/apiserver/fshandlers" ||--|{ "/domain/fsmodel" : x4
    "/controllers/apiserver/fshandlers" ||--|{ "/openapi/serverapi" : x11
    "/controllers/apiserver/fshandlers" ||--|{ "/pkg" : x1
    "/controllers/apiserver/hproxyhandlers" ||--|{ "/controllers/apiserver/apiservercore" : x4
    "/controllers/apiserver/hproxyhandlers" ||--|{ "/domain/hproxymodel" : x3
    "/controllers/apiserver/hproxyhandlers" ||--|{ "/openapi/serverapi" : x4
    "/controllers/apiserver/hproxyhandlers" ||--|{ "/pkg" : x2
    "/controllers/apiserver/labelhandlers" ||--|{ "/controllers/apiserver/apiservercore" : x9
    "/controllers/apiserver/labelhandlers" ||--|{ "/domain/core" : x7
    "/controllers/apiserver/labelhandlers" ||--|{ "/openapi/serverapi" : x8
    "/controllers/apiserver/labelhandlers" ||--|{ "/pkg" : x2
    "/controllers/apiserver/massloadhandlers" ||--|{ "/controllers/apiserver/apiservercore" : x17
    "/controllers/apiserver/massloadhandlers" ||--|{ "/domain/core" : x1
    "/controllers/apiserver/massloadhandlers" ||--|{ "/domain/massloadmodel" : x10
    "/controllers/apiserver/massloadhandlers" ||--|{ "/openapi/serverapi" : x16
    "/controllers/apiserver/massloadhandlers" ||--|{ "/pkg" : x3
    "/controllers/apiserver/systemhandlers" ||--|{ "/controllers/apiserver/apiservercore" : x11
    "/controllers/apiserver/systemhandlers" ||--|{ "/domain/core" : x1
    "/controllers/apiserver/systemhandlers" ||--|{ "/domain/parsing" : x5
    "/controllers/apiserver/systemhandlers" ||--|{ "/domain/systemmodel" : x4
    "/controllers/apiserver/systemhandlers" ||--|{ "/openapi/serverapi" : x12
    "/controllers/apiserver/systemhandlers" ||--|{ "/pkg" : x4
    "/controllers/internal/worker" ||--|{ "/pkg" : x1
    "/controllers/workermanager" ||--|{ "/controllers/internal/worker" : x10
    "/controllers/workermanager" ||--|{ "/domain/agentmodel" : x2
    "/controllers/workermanager" ||--|{ "/domain/core" : x3
    "/controllers/workermanager" ||--|{ "/domain/fsmodel" : x1
    "/controllers/workermanager" ||--|{ "/domain/massloadmodel" : x2
    "/controllers/workermanager" ||--|{ "/domain/parsing" : x1
    "/controllers/workermanager" ||--|{ "/domain/systemmodel" : x2
    "/controllers/workermanager" ||--|{ "/pkg" : x9
    "/domain/agentmodel" ||--|{ "/domain/core" : x2
    "/domain/bff" ||--|{ "/domain/core" : x3
    "/domain/bff" ||--|{ "/pkg" : x1
    "/domain/core" ||--|{ "/pkg" : x1
    "/domain/fsmodel" ||--|{ "/domain/core" : x1
    "/domain/parsing" ||--|{ "/domain/core" : x1
    "/external" ||--|{ "/domain/core" : x3
    "/external" ||--|{ "/pkg" : x2
    "/usecases/agentusecase" ||--|{ "/domain/agentmodel" : x2
    "/usecases/agentusecase" ||--|{ "/domain/core" : x6
    "/usecases/agentusecase" ||--|{ "/pkg" : x1
    "/usecases/attributeusecase" ||--|{ "/domain/core" : x5
    "/usecases/attributeusecase" ||--|{ "/domain/systemmodel" : x1
    "/usecases/bffusecase" ||--|{ "/domain/bff" : x4
    "/usecases/bffusecase" ||--|{ "/domain/core" : x3
    "/usecases/bffusecase" ||--|{ "/domain/fsmodel" : x1
    "/usecases/bffusecase" ||--|{ "/pkg" : x1
    "/usecases/bookusecase" ||--|{ "/domain/core" : x3
    "/usecases/cleanupusecase" ||--|{ "/domain/core" : x5
    "/usecases/cleanupusecase" ||--|{ "/domain/fsmodel" : x3
    "/usecases/cleanupusecase" ||--|{ "/domain/systemmodel" : x7
    "/usecases/cleanupusecase" ||--|{ "/pkg" : x2
    "/usecases/deduplicatorusecase" ||--|{ "/domain/bff" : x5
    "/usecases/deduplicatorusecase" ||--|{ "/domain/core" : x12
    "/usecases/deduplicatorusecase" ||--|{ "/domain/systemmodel" : x1
    "/usecases/deduplicatorusecase" ||--|{ "/external" : x1
    "/usecases/deduplicatorusecase" ||--|{ "/pkg" : x1
    "/usecases/exportusecase" ||--|{ "/domain/agentmodel" : x2
    "/usecases/exportusecase" ||--|{ "/domain/core" : x3
    "/usecases/exportusecase" ||--|{ "/domain/parsing" : x2
    "/usecases/exportusecase" ||--|{ "/external" : x2
    "/usecases/exportusecase" ||--|{ "/pkg" : x1
    "/usecases/filesystemusecase" ||--|{ "/domain/core" : x5
    "/usecases/filesystemusecase" ||--|{ "/domain/fsmodel" : x4
    "/usecases/filesystemusecase" ||--|{ "/pkg" : x1
    "/usecases/hproxyusecase" ||--|{ "/domain/agentmodel" : x4
    "/usecases/hproxyusecase" ||--|{ "/domain/core" : x5
    "/usecases/hproxyusecase" ||--|{ "/domain/hproxymodel" : x4
    "/usecases/hproxyusecase" ||--|{ "/domain/massloadmodel" : x2
    "/usecases/hproxyusecase" ||--|{ "/domain/parsing" : x3
    "/usecases/hproxyusecase" ||--|{ "/pkg" : x1
    "/usecases/labelusecase" ||--|{ "/domain/core" : x2
    "/usecases/labelusecase" ||--|{ "/domain/systemmodel" : x1
    "/usecases/massloadusecase" ||--|{ "/domain/agentmodel" : x2
    "/usecases/massloadusecase" ||--|{ "/domain/core" : x3
    "/usecases/massloadusecase" ||--|{ "/domain/hproxymodel" : x1
    "/usecases/massloadusecase" ||--|{ "/domain/massloadmodel" : x7
    "/usecases/massloadusecase" ||--|{ "/domain/parsing" : x2
    "/usecases/massloadusecase" ||--|{ "/domain/systemmodel" : x2
    "/usecases/massloadusecase" ||--|{ "/pkg" : x2
    "/usecases/parsingusecase" ||--|{ "/domain/agentmodel" : x9
    "/usecases/parsingusecase" ||--|{ "/domain/core" : x9
    "/usecases/parsingusecase" ||--|{ "/domain/parsing" : x5
    "/usecases/parsingusecase" ||--|{ "/pkg" : x4
    "/usecases/rebuilderusecase" ||--|{ "/domain/core" : x10
    "/usecases/rebuilderusecase" ||--|{ "/pkg" : x1
    "/usecases/systemusecase" ||--|{ "/domain/systemmodel" : x3
```

---

> Generated by [goArchLint](https://github.com/gbh007/goarchlint)
