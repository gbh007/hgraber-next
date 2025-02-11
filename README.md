# HGraber next

**Внимание:** в данный момент версия нестабильная и может изменятся без сохранения совместимости и данных

**Внимание:** legacy файловые системы будут удалены в следующем патче

Это пятая итерация HGraber, которая не является обратно совместимой с предыдущими и создана с целью обработки большего количества данных более эффективным способом.

[Базовый агент](https://github.com/gbh007/hgraber-next-agent-core) для системы

[UI на React TS](https://github.com/gbh007/hgraber-next-react-ui)

## Roadmap

1. Улучшение и отладка фич дедупликации
2. Улучшение и рефакторинг API
3. Стабилизация и релиз 1.0

## Словарь терминов используемых в приложении

| В коде    | Слово            | Значение                                                                                 | Примечание                                                    |
| --------- | ---------------- | ---------------------------------------------------------------------------------------- | ------------------------------------------------------------- |
| book      | Книга            | Минимальная структурированная единица данных в системе состоящая из страниц              |                                                               |
| agent     | Агент            | Система для первичной обработки и загрузки данных                                        |                                                               |
| page      | Страница         | Изображение и дополнительная информация связанная с ним                                  |                                                               |
| attribute | Атрибут          | Данные книги, для фильтрации (например автор)                                            |                                                               |
| label     | Метка            | Некая мета-информация о книге или ее странице                                            | всегда имеет вид пары ключ значение                           |
| deadHash  | Мертвый хеш      | Хеши файлов которые не представляют интереса для обработки и хранения                    |                                                               |
| fs        | Файловая система | Система для хранения файлов                                                              | Может быть как на агенте, так и на локальной файловой системе |
| highway   | Хайвей           | Способ получать данные напрямую с файловых агентов, без дополнительного проксирования    |                                                               |
| mirror    | Зеркало          | Дополнительный адрес для парсинга, совокупность адресов используется для проверки дублей | Ранее система проверки дублей по зеркалам была частью агента  |

## Прошлые версии

- HGraber (1-4) [Github](https://github.com/gbh007/hgraber)/[Gitlab](https://gitlab.com/gbh007/hgraber)

Отличия новой версии:

1. Изменения архитектуры БД с целью:
   - Уменьшения дублирования данных
2. Переход от PULL модели агентов к PUSH, для переноса основной логики в корневой сервер
   - Примечание: т.к. у некоторых сайтов есть зеркала и т.п. для отслеживания дубликатов требовалась отдельная логика в агентах
3. Переход на "промышленные" библиотеки и стандарты
   - Примечание: изначально система писалась как более близкая к пользователю (логи вида plain text и т.п.) и максимально "чистая"
4. Фичи дедупликации и очистки ненужных данных

## Схема взаимодействия

![schema](scheme.drawio.png)

На схеме изображено:

- 2 мастер системы A и B (важно, hgnext не имеет мультимастера)
  - Системы могут выгружать в друг друга данные и использовать друг друга как кеш (не будет похода в парсер)
  - Система A
    - Имеет локальную файловую систему A
    - Файловую систему через агента B
    - Базу данных A
  - Система B
    - Файловую систему через агента C
    - Базу данных B
- 3 Агента A, B, C
  - Агент A
    - Имеет локальную файловую систему B
    - Умеет парсить сайт A
    - Поддерживает экспорт
  - Агент B
    - Умеет парсить сайт B и C
  - Агент C
    - Имеет локальную файловую систему C
    - Умеет парсить сайт D и E

## Пример настройка логов и метрик (Grafana stack)

Генерация борды (`jsonnet/dashboard.json`) с кастомной конфигурацией

> HG_SERVICES="a,b,c" make jsonnet-custom

Docker compose

```yaml
services:
  main:
    container_name: hgraber-next-main
    logging:
      driver: loki
      options:
        loki-url: "http://localhost:3100/loki/api/v1/push"
```

Prometheus

```yaml
scrape_configs:
  - job_name: hgraber-next
    static_configs:
      - targets:
          - localhost:8080
        labels:
          service_name: hgraber-next-main
```

Promtail

```yaml
scrape_configs:
  - job_name: hgraber-next
    static_configs:
      - targets:
          - localhost
        labels:
          job: hgraber-next
          __path__: <path to logs>
          service_name: hgraber-next-main
```

## Пример настройки CD в Jenkins

В данном примере приведен пример пайплайна, не рекомендуется его использовать как готовый

```bash
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
bash build.bash

docker build --build-arg "BINARY_PATH=_build/server-linux-amd64" -t hgraber-next-server:latest .
docker compose -f "${DC_PATH}/docker-compose.yml" up -d --remove-orphans

cd jsonnet
go run github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb@latest install github.com/grafana/grafonnet/gen/grafonnet-latest@main

curl -X 'POST' \
  "${GRAFANA_API_HOST}/api/dashboards/import" \
  -H 'accept: application/json' \
  -H "Authorization: Bearer ${GRAFANA_API_TOKEN}" \
  -H 'Content-Type: application/json' \
  -d "{
  \"dashboard\": $(go run github.com/google/go-jsonnet/cmd/jsonnet@latest --ext-str services="${HG_SERVICES}" -J vendor dashboard.jsonnet),
  \"overwrite\": true
}"
```
