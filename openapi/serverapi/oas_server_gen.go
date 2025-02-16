// Code generated by ogen, DO NOT EDIT.

package serverapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APIAgentDeletePost implements POST /api/agent/delete operation.
	//
	// Удаление агента.
	//
	// POST /api/agent/delete
	APIAgentDeletePost(ctx context.Context, req *APIAgentDeletePostReq) (APIAgentDeletePostRes, error)
	// APIAgentGetPost implements POST /api/agent/get operation.
	//
	// Получение данных агента.
	//
	// POST /api/agent/get
	APIAgentGetPost(ctx context.Context, req *APIAgentGetPostReq) (APIAgentGetPostRes, error)
	// APIAgentListPost implements POST /api/agent/list operation.
	//
	// Список агентов.
	//
	// POST /api/agent/list
	APIAgentListPost(ctx context.Context, req *APIAgentListPostReq) (APIAgentListPostRes, error)
	// APIAgentNewPost implements POST /api/agent/new operation.
	//
	// Создание нового агента.
	//
	// POST /api/agent/new
	APIAgentNewPost(ctx context.Context, req *APIAgentNewPostReq) (APIAgentNewPostRes, error)
	// APIAgentTaskExportPost implements POST /api/agent/task/export operation.
	//
	// Экспорт книг в другую систему.
	//
	// POST /api/agent/task/export
	APIAgentTaskExportPost(ctx context.Context, req *APIAgentTaskExportPostReq) (APIAgentTaskExportPostRes, error)
	// APIAgentUpdatePost implements POST /api/agent/update operation.
	//
	// Обновление данных агента.
	//
	// POST /api/agent/update
	APIAgentUpdatePost(ctx context.Context, req *Agent) (APIAgentUpdatePostRes, error)
	// APIAttributeColorCreatePost implements POST /api/attribute/color/create operation.
	//
	// Создание покраски аттрибута.
	//
	// POST /api/attribute/color/create
	APIAttributeColorCreatePost(ctx context.Context, req *APIAttributeColorCreatePostReq) (APIAttributeColorCreatePostRes, error)
	// APIAttributeColorDeletePost implements POST /api/attribute/color/delete operation.
	//
	// Удаление цвета атрибута.
	//
	// POST /api/attribute/color/delete
	APIAttributeColorDeletePost(ctx context.Context, req *APIAttributeColorDeletePostReq) (APIAttributeColorDeletePostRes, error)
	// APIAttributeColorGetPost implements POST /api/attribute/color/get operation.
	//
	// Цвет конкретного атрибута.
	//
	// POST /api/attribute/color/get
	APIAttributeColorGetPost(ctx context.Context, req *APIAttributeColorGetPostReq) (APIAttributeColorGetPostRes, error)
	// APIAttributeColorListGet implements GET /api/attribute/color/list operation.
	//
	// Цвета атрибутов.
	//
	// GET /api/attribute/color/list
	APIAttributeColorListGet(ctx context.Context) (APIAttributeColorListGetRes, error)
	// APIAttributeColorUpdatePost implements POST /api/attribute/color/update operation.
	//
	// Обновления покраски атрибута.
	//
	// POST /api/attribute/color/update
	APIAttributeColorUpdatePost(ctx context.Context, req *APIAttributeColorUpdatePostReq) (APIAttributeColorUpdatePostRes, error)
	// APIAttributeCountGet implements GET /api/attribute/count operation.
	//
	// Получение информации о количестве вариантов
	// атрибутов.
	//
	// GET /api/attribute/count
	APIAttributeCountGet(ctx context.Context) (APIAttributeCountGetRes, error)
	// APIBookArchiveIDGet implements GET /api/book/archive/{id} operation.
	//
	// Получение архива с книгой.
	//
	// GET /api/book/archive/{id}
	APIBookArchiveIDGet(ctx context.Context, params APIBookArchiveIDGetParams) (APIBookArchiveIDGetRes, error)
	// APIBookDeletePost implements POST /api/book/delete operation.
	//
	// Удаляет книгу и/или ее страницы.
	//
	// POST /api/book/delete
	APIBookDeletePost(ctx context.Context, req *APIBookDeletePostReq) (APIBookDeletePostRes, error)
	// APIBookDetailsPost implements POST /api/book/details operation.
	//
	// Информация о книге.
	//
	// POST /api/book/details
	APIBookDetailsPost(ctx context.Context, req *APIBookDetailsPostReq) (APIBookDetailsPostRes, error)
	// APIBookListPost implements POST /api/book/list operation.
	//
	// Список книг.
	//
	// POST /api/book/list
	APIBookListPost(ctx context.Context, req *BookFilter) (APIBookListPostRes, error)
	// APIBookPageBodyPost implements POST /api/book/page/body operation.
	//
	// Получение тела страницы (по оригинальному адресу или
	// данным книги).
	//
	// POST /api/book/page/body
	APIBookPageBodyPost(ctx context.Context, req *APIBookPageBodyPostReq) (APIBookPageBodyPostRes, error)
	// APIBookPageDeletePost implements POST /api/book/page/delete operation.
	//
	// Удаляет страницы из книг.
	//
	// POST /api/book/page/delete
	APIBookPageDeletePost(ctx context.Context, req *APIBookPageDeletePostReq) (APIBookPageDeletePostRes, error)
	// APIBookRawPost implements POST /api/book/raw operation.
	//
	// Информация о книге (или по ИД или по адресу).
	//
	// POST /api/book/raw
	APIBookRawPost(ctx context.Context, req *APIBookRawPostReq) (APIBookRawPostRes, error)
	// APIBookRebuildPost implements POST /api/book/rebuild operation.
	//
	// Может как создать новую книгу, так и добавить данные в
	// другую пересобранную.
	//
	// POST /api/book/rebuild
	APIBookRebuildPost(ctx context.Context, req *APIBookRebuildPostReq) (APIBookRebuildPostRes, error)
	// APIBookRestorePost implements POST /api/book/restore operation.
	//
	// Пытается восстановить книгу или ее страницы
	// (восстановление может быть не возможно если данные
	// уже были очищены).
	//
	// POST /api/book/restore
	APIBookRestorePost(ctx context.Context, req *APIBookRestorePostReq) (APIBookRestorePostRes, error)
	// APIBookStatusSetPost implements POST /api/book/status/set operation.
	//
	// Изменение статуса книги.
	//
	// POST /api/book/status/set
	APIBookStatusSetPost(ctx context.Context, req *APIBookStatusSetPostReq) (APIBookStatusSetPostRes, error)
	// APIBookUpdatePost implements POST /api/book/update operation.
	//
	// Изменяет часть данных книги, ряд полей не изменяется
	// (верификация, число страниц и т.д.).
	//
	// POST /api/book/update
	APIBookUpdatePost(ctx context.Context, req *BookRaw) (APIBookUpdatePostRes, error)
	// APIDeduplicateArchivePost implements POST /api/deduplicate/archive operation.
	//
	// Проверка наличия данных в системе из архива.
	//
	// POST /api/deduplicate/archive
	APIDeduplicateArchivePost(ctx context.Context, req APIDeduplicateArchivePostReq) (APIDeduplicateArchivePostRes, error)
	// APIDeduplicateBookByPageBodyPost implements POST /api/deduplicate/book-by-page-body operation.
	//
	// Поиск дубликатов книги по телу страницы.
	//
	// POST /api/deduplicate/book-by-page-body
	APIDeduplicateBookByPageBodyPost(ctx context.Context, req *APIDeduplicateBookByPageBodyPostReq) (APIDeduplicateBookByPageBodyPostRes, error)
	// APIDeduplicateBooksByPagePost implements POST /api/deduplicate/books-by-page operation.
	//
	// Поиск книг содержащих такую же страницу (тело).
	//
	// POST /api/deduplicate/books-by-page
	APIDeduplicateBooksByPagePost(ctx context.Context, req *APIDeduplicateBooksByPagePostReq) (APIDeduplicateBooksByPagePostRes, error)
	// APIDeduplicateComparePost implements POST /api/deduplicate/compare operation.
	//
	// Сравнение двух книг на дублируемые страницы.
	//
	// POST /api/deduplicate/compare
	APIDeduplicateComparePost(ctx context.Context, req *APIDeduplicateComparePostReq) (APIDeduplicateComparePostRes, error)
	// APIDeduplicateDeadHashSetPost implements POST /api/deduplicate/dead-hash/set operation.
	//
	// Устанавливает значение мертвых хешей для книги или
	// ее страницы.
	//
	// POST /api/deduplicate/dead-hash/set
	APIDeduplicateDeadHashSetPost(ctx context.Context, req *APIDeduplicateDeadHashSetPostReq) (APIDeduplicateDeadHashSetPostRes, error)
	// APIDeduplicateUniquePagesPost implements POST /api/deduplicate/unique-pages operation.
	//
	// Поиск уникальных страниц в книге.
	//
	// POST /api/deduplicate/unique-pages
	APIDeduplicateUniquePagesPost(ctx context.Context, req *APIDeduplicateUniquePagesPostReq) (APIDeduplicateUniquePagesPostRes, error)
	// APIFileIDGet implements GET /api/file/{id} operation.
	//
	// Получение тела файла (изображения страницы).
	//
	// GET /api/file/{id}
	APIFileIDGet(ctx context.Context, params APIFileIDGetParams) (APIFileIDGetRes, error)
	// APIFsCreatePost implements POST /api/fs/create operation.
	//
	// Создание файловой системы.
	//
	// POST /api/fs/create
	APIFsCreatePost(ctx context.Context, req *APIFsCreatePostReq) (APIFsCreatePostRes, error)
	// APIFsDeletePost implements POST /api/fs/delete operation.
	//
	// Удаление файловой системы.
	//
	// POST /api/fs/delete
	APIFsDeletePost(ctx context.Context, req *APIFsDeletePostReq) (APIFsDeletePostRes, error)
	// APIFsGetPost implements POST /api/fs/get operation.
	//
	// Данные настроек файловой системы.
	//
	// POST /api/fs/get
	APIFsGetPost(ctx context.Context, req *APIFsGetPostReq) (APIFsGetPostRes, error)
	// APIFsListPost implements POST /api/fs/list operation.
	//
	// Список файловых систем.
	//
	// POST /api/fs/list
	APIFsListPost(ctx context.Context, req *APIFsListPostReq) (APIFsListPostRes, error)
	// APIFsRemoveMismatchPost implements POST /api/fs/remove-mismatch operation.
	//
	// Запускает задачу удаления не совпавших файлов между
	// базой данных и файловым хранилищем.
	//
	// POST /api/fs/remove-mismatch
	APIFsRemoveMismatchPost(ctx context.Context, req *APIFsRemoveMismatchPostReq) (APIFsRemoveMismatchPostRes, error)
	// APIFsTransferBookPost implements POST /api/fs/transfer/book operation.
	//
	// Запускает перенос файлов между файловыми системами.
	//
	// POST /api/fs/transfer/book
	APIFsTransferBookPost(ctx context.Context, req *APIFsTransferBookPostReq) (APIFsTransferBookPostRes, error)
	// APIFsTransferPost implements POST /api/fs/transfer operation.
	//
	// Запускает перенос файлов между файловыми системами.
	//
	// POST /api/fs/transfer
	APIFsTransferPost(ctx context.Context, req *APIFsTransferPostReq) (APIFsTransferPostRes, error)
	// APIFsUpdatePost implements POST /api/fs/update operation.
	//
	// Изменение настроек файловой системы.
	//
	// POST /api/fs/update
	APIFsUpdatePost(ctx context.Context, req *APIFsUpdatePostReq) (APIFsUpdatePostRes, error)
	// APIFsValidatePost implements POST /api/fs/validate operation.
	//
	// Запускает валидацию файлов на файловой системе.
	//
	// POST /api/fs/validate
	APIFsValidatePost(ctx context.Context, req *APIFsValidatePostReq) (APIFsValidatePostRes, error)
	// APILabelDeletePost implements POST /api/label/delete operation.
	//
	// Удаление метки на книгу или страницу.
	//
	// POST /api/label/delete
	APILabelDeletePost(ctx context.Context, req *APILabelDeletePostReq) (APILabelDeletePostRes, error)
	// APILabelGetPost implements POST /api/label/get operation.
	//
	// Получение меток книги.
	//
	// POST /api/label/get
	APILabelGetPost(ctx context.Context, req *APILabelGetPostReq) (APILabelGetPostRes, error)
	// APILabelPresetCreatePost implements POST /api/label/preset/create operation.
	//
	// Создание пресета меток.
	//
	// POST /api/label/preset/create
	APILabelPresetCreatePost(ctx context.Context, req *APILabelPresetCreatePostReq) (APILabelPresetCreatePostRes, error)
	// APILabelPresetDeletePost implements POST /api/label/preset/delete operation.
	//
	// Удаление пресета меток.
	//
	// POST /api/label/preset/delete
	APILabelPresetDeletePost(ctx context.Context, req *APILabelPresetDeletePostReq) (APILabelPresetDeletePostRes, error)
	// APILabelPresetGetPost implements POST /api/label/preset/get operation.
	//
	// Пресеты меток.
	//
	// POST /api/label/preset/get
	APILabelPresetGetPost(ctx context.Context, req *APILabelPresetGetPostReq) (APILabelPresetGetPostRes, error)
	// APILabelPresetListGet implements GET /api/label/preset/list operation.
	//
	// Пресеты меток.
	//
	// GET /api/label/preset/list
	APILabelPresetListGet(ctx context.Context) (APILabelPresetListGetRes, error)
	// APILabelPresetUpdatePost implements POST /api/label/preset/update operation.
	//
	// Обновления пресета меток.
	//
	// POST /api/label/preset/update
	APILabelPresetUpdatePost(ctx context.Context, req *APILabelPresetUpdatePostReq) (APILabelPresetUpdatePostRes, error)
	// APILabelSetPost implements POST /api/label/set operation.
	//
	// Установка метки на книгу или страницу.
	//
	// POST /api/label/set
	APILabelSetPost(ctx context.Context, req *APILabelSetPostReq) (APILabelSetPostRes, error)
	// APIParsingHandlePost implements POST /api/parsing/handle operation.
	//
	// Обработка ссылок на новые книги.
	//
	// POST /api/parsing/handle
	APIParsingHandlePost(ctx context.Context, req *APIParsingHandlePostReq) (APIParsingHandlePostRes, error)
	// APIParsingMirrorCreatePost implements POST /api/parsing/mirror/create operation.
	//
	// Создание данных зеркала.
	//
	// POST /api/parsing/mirror/create
	APIParsingMirrorCreatePost(ctx context.Context, req *APIParsingMirrorCreatePostReq) (APIParsingMirrorCreatePostRes, error)
	// APIParsingMirrorDeletePost implements POST /api/parsing/mirror/delete operation.
	//
	// Удаление зеркала.
	//
	// POST /api/parsing/mirror/delete
	APIParsingMirrorDeletePost(ctx context.Context, req *APIParsingMirrorDeletePostReq) (APIParsingMirrorDeletePostRes, error)
	// APIParsingMirrorGetPost implements POST /api/parsing/mirror/get operation.
	//
	// Получение данных зеркала.
	//
	// POST /api/parsing/mirror/get
	APIParsingMirrorGetPost(ctx context.Context, req *APIParsingMirrorGetPostReq) (APIParsingMirrorGetPostRes, error)
	// APIParsingMirrorListGet implements GET /api/parsing/mirror/list operation.
	//
	// Зеркала.
	//
	// GET /api/parsing/mirror/list
	APIParsingMirrorListGet(ctx context.Context) (APIParsingMirrorListGetRes, error)
	// APIParsingMirrorUpdatePost implements POST /api/parsing/mirror/update operation.
	//
	// Обновления зеркала.
	//
	// POST /api/parsing/mirror/update
	APIParsingMirrorUpdatePost(ctx context.Context, req *APIParsingMirrorUpdatePostReq) (APIParsingMirrorUpdatePostRes, error)
	// APISystemImportArchivePost implements POST /api/system/import/archive operation.
	//
	// Импорт новой книги через архив.
	//
	// POST /api/system/import/archive
	APISystemImportArchivePost(ctx context.Context, req APISystemImportArchivePostReq) (APISystemImportArchivePostRes, error)
	// APISystemInfoSizeGet implements GET /api/system/info/size operation.
	//
	// Получение общей информации о системе.
	//
	// GET /api/system/info/size
	APISystemInfoSizeGet(ctx context.Context) (APISystemInfoSizeGetRes, error)
	// APISystemInfoWorkersGet implements GET /api/system/info/workers operation.
	//
	// Получение информации о воркерах в системе.
	//
	// GET /api/system/info/workers
	APISystemInfoWorkersGet(ctx context.Context) (APISystemInfoWorkersGetRes, error)
	// APISystemTaskCreatePost implements POST /api/system/task/create operation.
	//
	// Создание и фоновый запуск задачи.
	//
	// POST /api/system/task/create
	APISystemTaskCreatePost(ctx context.Context, req *APISystemTaskCreatePostReq) (APISystemTaskCreatePostRes, error)
	// APISystemTaskResultsGet implements GET /api/system/task/results operation.
	//
	// Получение результатов задач.
	//
	// GET /api/system/task/results
	APISystemTaskResultsGet(ctx context.Context) (APISystemTaskResultsGetRes, error)
	// APISystemWorkerConfigPost implements POST /api/system/worker/config operation.
	//
	// Динамическая конфигурация раннеров (воркеров),
	// сбрасывается при перезапуске системы.
	//
	// POST /api/system/worker/config
	APISystemWorkerConfigPost(ctx context.Context, req *APISystemWorkerConfigPostReq) (APISystemWorkerConfigPostRes, error)
	// APIUserLoginPost implements POST /api/user/login operation.
	//
	// Проставление токена в куки.
	//
	// POST /api/user/login
	APIUserLoginPost(ctx context.Context, req *APIUserLoginPostReq) (APIUserLoginPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
