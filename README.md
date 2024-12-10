# HGraber next

**Внимание:** в данный момент версия нестабильная и может изменятся без сохранения совместимости и данных

Это пятая итерация HGraber, которая не является обратно совместимой с предыдущими и создана с целью обработки большего количества данных более эффективным способом.

[Пример агента](https://github.com/gbh007/hgraber-next-agent-example) для системы

**Важно:** сейчас в разработки целевая версия UI на React TS, [ссылка на репозиторий](https://github.com/gbh007/hgraber-next-react-ui)

## Roadmap

1. Воркеры (или другой подход) для обработки крайне долгих задач с этапами (в основном для очистки системы)
2. Дедупликация
   - [x] Файлов
   - Аттрибутов
   - Книг
3. Добавить новый функционал:
   - Создание перестроенных книг - на основании существующих
     - Книга будет в отдельной таблице
     - Данные страниц будут вида `buildedBookID | buildedPageNumber | originBookID | originPageNumber | fileID`
   - Полноценные фильтры и поиск
   - Расширение данных оценок
     - [x] Полный отказ от текущей системы оценок (будет потеря данных рейтингов)
     - Добавление более детальных оценок для книги (как общий обзор на нее)
     - Добавление более простых оценок для страницы - не оценено, понравилось, идеально
       - Пояснение - при 5-ти бальной системе оценки как правило 3 и выше, а оценки 1 и 2 не ставятся и вместо них страница просто пропускается

## Словарь терминов используемых в приложении

| В коде | Слово    | Значение                                                                    | Примечание |
| ------ | -------- | --------------------------------------------------------------------------- | ---------- |
| book   | Книга    | Минимальная структурированная единица данных в системе состоящая из страниц |            |
| agent  | Агент    | Система для первичной обработки и загрузки данных                           |            |
| page   | Страница | Изображение и дополнительная информация связанная с ним                     |            |

## Прошлые версии

- HGraber (1-4) [Github](https://github.com/gbh007/hgraber)/[Gitlab](https://gitlab.com/gbh007/hgraber)

Отличия новой версии:

1. Изменения архитектуры БД с целью:
   - Уменьшения дублирования данных
   - Возможность использования системы несколькими пользователями
2. Переход от PULL модели агентов к PUSH, для переноса основной логики в корневой сервер
   - Примечание: т.к. у некоторых сайтов есть зеркала и т.п. для отслеживания дубликатов требовалась отдельная логика в агентах
3. Переход на "промышленные" библиотеки и стандарты
   - Примечание: изначально система писалась как более близкая к пользователю (логи вида plain text и т.п.) и максимально "чистая"
