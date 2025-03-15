// Code generated by ogen, DO NOT EDIT.

package serverapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// APIAgentDeletePost implements POST /api/agent/delete operation.
//
// Удаление агента.
//
// POST /api/agent/delete
func (UnimplementedHandler) APIAgentDeletePost(ctx context.Context, req *APIAgentDeletePostReq) (r APIAgentDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentGetPost implements POST /api/agent/get operation.
//
// Получение данных агента.
//
// POST /api/agent/get
func (UnimplementedHandler) APIAgentGetPost(ctx context.Context, req *APIAgentGetPostReq) (r APIAgentGetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentListPost implements POST /api/agent/list operation.
//
// Список агентов.
//
// POST /api/agent/list
func (UnimplementedHandler) APIAgentListPost(ctx context.Context, req *APIAgentListPostReq) (r APIAgentListPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentNewPost implements POST /api/agent/new operation.
//
// Создание нового агента.
//
// POST /api/agent/new
func (UnimplementedHandler) APIAgentNewPost(ctx context.Context, req *APIAgentNewPostReq) (r APIAgentNewPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentTaskExportPost implements POST /api/agent/task/export operation.
//
// Экспорт книг в другую систему.
//
// POST /api/agent/task/export
func (UnimplementedHandler) APIAgentTaskExportPost(ctx context.Context, req *APIAgentTaskExportPostReq) (r APIAgentTaskExportPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAgentUpdatePost implements POST /api/agent/update operation.
//
// Обновление данных агента.
//
// POST /api/agent/update
func (UnimplementedHandler) APIAgentUpdatePost(ctx context.Context, req *APIAgentUpdatePostReq) (r APIAgentUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeColorCreatePost implements POST /api/attribute/color/create operation.
//
// Создание покраски аттрибута.
//
// POST /api/attribute/color/create
func (UnimplementedHandler) APIAttributeColorCreatePost(ctx context.Context, req *APIAttributeColorCreatePostReq) (r APIAttributeColorCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeColorDeletePost implements POST /api/attribute/color/delete operation.
//
// Удаление цвета атрибута.
//
// POST /api/attribute/color/delete
func (UnimplementedHandler) APIAttributeColorDeletePost(ctx context.Context, req *APIAttributeColorDeletePostReq) (r APIAttributeColorDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeColorGetPost implements POST /api/attribute/color/get operation.
//
// Цвет конкретного атрибута.
//
// POST /api/attribute/color/get
func (UnimplementedHandler) APIAttributeColorGetPost(ctx context.Context, req *APIAttributeColorGetPostReq) (r APIAttributeColorGetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeColorListGet implements GET /api/attribute/color/list operation.
//
// Цвета атрибутов.
//
// GET /api/attribute/color/list
func (UnimplementedHandler) APIAttributeColorListGet(ctx context.Context) (r APIAttributeColorListGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeColorUpdatePost implements POST /api/attribute/color/update operation.
//
// Обновления покраски атрибута.
//
// POST /api/attribute/color/update
func (UnimplementedHandler) APIAttributeColorUpdatePost(ctx context.Context, req *APIAttributeColorUpdatePostReq) (r APIAttributeColorUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeCountGet implements GET /api/attribute/count operation.
//
// Получение информации о количестве вариантов
// атрибутов.
//
// GET /api/attribute/count
func (UnimplementedHandler) APIAttributeCountGet(ctx context.Context) (r APIAttributeCountGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeOriginCountGet implements GET /api/attribute/origin/count operation.
//
// Получение информации о количестве вариантов
// оригинальных атрибутов.
//
// GET /api/attribute/origin/count
func (UnimplementedHandler) APIAttributeOriginCountGet(ctx context.Context) (r APIAttributeOriginCountGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeRemapCreatePost implements POST /api/attribute/remap/create operation.
//
// Создание ремапинга аттрибута.
//
// POST /api/attribute/remap/create
func (UnimplementedHandler) APIAttributeRemapCreatePost(ctx context.Context, req *APIAttributeRemapCreatePostReq) (r APIAttributeRemapCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeRemapDeletePost implements POST /api/attribute/remap/delete operation.
//
// Удаление ремапинга атрибута.
//
// POST /api/attribute/remap/delete
func (UnimplementedHandler) APIAttributeRemapDeletePost(ctx context.Context, req *APIAttributeRemapDeletePostReq) (r APIAttributeRemapDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeRemapGetPost implements POST /api/attribute/remap/get operation.
//
// Ремапинг конкретного атрибута.
//
// POST /api/attribute/remap/get
func (UnimplementedHandler) APIAttributeRemapGetPost(ctx context.Context, req *APIAttributeRemapGetPostReq) (r APIAttributeRemapGetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeRemapListGet implements GET /api/attribute/remap/list operation.
//
// Ремапинги атрибутов.
//
// GET /api/attribute/remap/list
func (UnimplementedHandler) APIAttributeRemapListGet(ctx context.Context) (r APIAttributeRemapListGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIAttributeRemapUpdatePost implements POST /api/attribute/remap/update operation.
//
// Обновления ремапинга атрибута.
//
// POST /api/attribute/remap/update
func (UnimplementedHandler) APIAttributeRemapUpdatePost(ctx context.Context, req *APIAttributeRemapUpdatePostReq) (r APIAttributeRemapUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookArchiveIDGet implements GET /api/book/archive/{id} operation.
//
// Получение архива с книгой.
//
// GET /api/book/archive/{id}
func (UnimplementedHandler) APIBookArchiveIDGet(ctx context.Context, params APIBookArchiveIDGetParams) (r APIBookArchiveIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookDeletePost implements POST /api/book/delete operation.
//
// Удаляет книгу и/или ее страницы.
//
// POST /api/book/delete
func (UnimplementedHandler) APIBookDeletePost(ctx context.Context, req *APIBookDeletePostReq) (r APIBookDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookDetailsPost implements POST /api/book/details operation.
//
// Информация о книге.
//
// POST /api/book/details
func (UnimplementedHandler) APIBookDetailsPost(ctx context.Context, req *APIBookDetailsPostReq) (r APIBookDetailsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookListPost implements POST /api/book/list operation.
//
// Список книг.
//
// POST /api/book/list
func (UnimplementedHandler) APIBookListPost(ctx context.Context, req *BookFilter) (r APIBookListPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookPageBodyPost implements POST /api/book/page/body operation.
//
// Получение тела страницы (по оригинальному адресу или
// данным книги).
//
// POST /api/book/page/body
func (UnimplementedHandler) APIBookPageBodyPost(ctx context.Context, req *APIBookPageBodyPostReq) (r APIBookPageBodyPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookPageDeletePost implements POST /api/book/page/delete operation.
//
// Удаляет страницы из книг.
//
// POST /api/book/page/delete
func (UnimplementedHandler) APIBookPageDeletePost(ctx context.Context, req *APIBookPageDeletePostReq) (r APIBookPageDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookRawPost implements POST /api/book/raw operation.
//
// Информация о книге (или по ИД или по адресу).
//
// POST /api/book/raw
func (UnimplementedHandler) APIBookRawPost(ctx context.Context, req *APIBookRawPostReq) (r APIBookRawPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookRebuildPost implements POST /api/book/rebuild operation.
//
// Может как создать новую книгу, так и добавить данные в
// другую пересобранную.
//
// POST /api/book/rebuild
func (UnimplementedHandler) APIBookRebuildPost(ctx context.Context, req *APIBookRebuildPostReq) (r APIBookRebuildPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookRestorePost implements POST /api/book/restore operation.
//
// Пытается восстановить книгу или ее страницы
// (восстановление может быть не возможно если данные
// уже были очищены).
//
// POST /api/book/restore
func (UnimplementedHandler) APIBookRestorePost(ctx context.Context, req *APIBookRestorePostReq) (r APIBookRestorePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookStatusSetPost implements POST /api/book/status/set operation.
//
// Изменение статуса книги.
//
// POST /api/book/status/set
func (UnimplementedHandler) APIBookStatusSetPost(ctx context.Context, req *APIBookStatusSetPostReq) (r APIBookStatusSetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookUpdatePost implements POST /api/book/update operation.
//
// Изменяет часть данных книги, ряд полей не изменяется
// (верификация, число страниц и т.д.).
//
// POST /api/book/update
func (UnimplementedHandler) APIBookUpdatePost(ctx context.Context, req *BookRaw) (r APIBookUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateArchivePost implements POST /api/deduplicate/archive operation.
//
// Проверка наличия данных в системе из архива.
//
// POST /api/deduplicate/archive
func (UnimplementedHandler) APIDeduplicateArchivePost(ctx context.Context, req APIDeduplicateArchivePostReq) (r APIDeduplicateArchivePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateBookByPageBodyPost implements POST /api/deduplicate/book-by-page-body operation.
//
// Поиск дубликатов книги по телу страницы.
//
// POST /api/deduplicate/book-by-page-body
func (UnimplementedHandler) APIDeduplicateBookByPageBodyPost(ctx context.Context, req *APIDeduplicateBookByPageBodyPostReq) (r APIDeduplicateBookByPageBodyPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateBooksByPagePost implements POST /api/deduplicate/books-by-page operation.
//
// Поиск книг содержащих такую же страницу (тело).
//
// POST /api/deduplicate/books-by-page
func (UnimplementedHandler) APIDeduplicateBooksByPagePost(ctx context.Context, req *APIDeduplicateBooksByPagePostReq) (r APIDeduplicateBooksByPagePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateComparePost implements POST /api/deduplicate/compare operation.
//
// Сравнение двух книг на дублируемые страницы.
//
// POST /api/deduplicate/compare
func (UnimplementedHandler) APIDeduplicateComparePost(ctx context.Context, req *APIDeduplicateComparePostReq) (r APIDeduplicateComparePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateDeadHashSetPost implements POST /api/deduplicate/dead-hash/set operation.
//
// Устанавливает значение мертвых хешей для книги или
// ее страницы.
//
// POST /api/deduplicate/dead-hash/set
func (UnimplementedHandler) APIDeduplicateDeadHashSetPost(ctx context.Context, req *APIDeduplicateDeadHashSetPostReq) (r APIDeduplicateDeadHashSetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateUniquePagesPost implements POST /api/deduplicate/unique-pages operation.
//
// Поиск уникальных страниц в книге.
//
// POST /api/deduplicate/unique-pages
func (UnimplementedHandler) APIDeduplicateUniquePagesPost(ctx context.Context, req *APIDeduplicateUniquePagesPostReq) (r APIDeduplicateUniquePagesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFileIDGet implements GET /api/file/{id} operation.
//
// Получение тела файла (изображения страницы).
//
// GET /api/file/{id}
func (UnimplementedHandler) APIFileIDGet(ctx context.Context, params APIFileIDGetParams) (r APIFileIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsCreatePost implements POST /api/fs/create operation.
//
// Создание файловой системы.
//
// POST /api/fs/create
func (UnimplementedHandler) APIFsCreatePost(ctx context.Context, req *APIFsCreatePostReq) (r APIFsCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsDeletePost implements POST /api/fs/delete operation.
//
// Удаление файловой системы.
//
// POST /api/fs/delete
func (UnimplementedHandler) APIFsDeletePost(ctx context.Context, req *APIFsDeletePostReq) (r APIFsDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsGetPost implements POST /api/fs/get operation.
//
// Данные настроек файловой системы.
//
// POST /api/fs/get
func (UnimplementedHandler) APIFsGetPost(ctx context.Context, req *APIFsGetPostReq) (r APIFsGetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsListPost implements POST /api/fs/list operation.
//
// Список файловых систем.
//
// POST /api/fs/list
func (UnimplementedHandler) APIFsListPost(ctx context.Context, req *APIFsListPostReq) (r APIFsListPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsRemoveMismatchPost implements POST /api/fs/remove-mismatch operation.
//
// Запускает задачу удаления не совпавших файлов между
// базой данных и файловым хранилищем.
//
// POST /api/fs/remove-mismatch
func (UnimplementedHandler) APIFsRemoveMismatchPost(ctx context.Context, req *APIFsRemoveMismatchPostReq) (r APIFsRemoveMismatchPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsTransferBookPost implements POST /api/fs/transfer/book operation.
//
// Запускает перенос файлов между файловыми системами.
//
// POST /api/fs/transfer/book
func (UnimplementedHandler) APIFsTransferBookPost(ctx context.Context, req *APIFsTransferBookPostReq) (r APIFsTransferBookPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsTransferPost implements POST /api/fs/transfer operation.
//
// Запускает перенос файлов между файловыми системами.
//
// POST /api/fs/transfer
func (UnimplementedHandler) APIFsTransferPost(ctx context.Context, req *APIFsTransferPostReq) (r APIFsTransferPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsUpdatePost implements POST /api/fs/update operation.
//
// Изменение настроек файловой системы.
//
// POST /api/fs/update
func (UnimplementedHandler) APIFsUpdatePost(ctx context.Context, req *APIFsUpdatePostReq) (r APIFsUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIFsValidatePost implements POST /api/fs/validate operation.
//
// Запускает валидацию файлов на файловой системе.
//
// POST /api/fs/validate
func (UnimplementedHandler) APIFsValidatePost(ctx context.Context, req *APIFsValidatePostReq) (r APIFsValidatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelDeletePost implements POST /api/label/delete operation.
//
// Удаление метки на книгу или страницу.
//
// POST /api/label/delete
func (UnimplementedHandler) APILabelDeletePost(ctx context.Context, req *APILabelDeletePostReq) (r APILabelDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelGetPost implements POST /api/label/get operation.
//
// Получение меток книги.
//
// POST /api/label/get
func (UnimplementedHandler) APILabelGetPost(ctx context.Context, req *APILabelGetPostReq) (r APILabelGetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelPresetCreatePost implements POST /api/label/preset/create operation.
//
// Создание пресета меток.
//
// POST /api/label/preset/create
func (UnimplementedHandler) APILabelPresetCreatePost(ctx context.Context, req *APILabelPresetCreatePostReq) (r APILabelPresetCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelPresetDeletePost implements POST /api/label/preset/delete operation.
//
// Удаление пресета меток.
//
// POST /api/label/preset/delete
func (UnimplementedHandler) APILabelPresetDeletePost(ctx context.Context, req *APILabelPresetDeletePostReq) (r APILabelPresetDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelPresetGetPost implements POST /api/label/preset/get operation.
//
// Пресеты меток.
//
// POST /api/label/preset/get
func (UnimplementedHandler) APILabelPresetGetPost(ctx context.Context, req *APILabelPresetGetPostReq) (r APILabelPresetGetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelPresetListGet implements GET /api/label/preset/list operation.
//
// Пресеты меток.
//
// GET /api/label/preset/list
func (UnimplementedHandler) APILabelPresetListGet(ctx context.Context) (r APILabelPresetListGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelPresetUpdatePost implements POST /api/label/preset/update operation.
//
// Обновления пресета меток.
//
// POST /api/label/preset/update
func (UnimplementedHandler) APILabelPresetUpdatePost(ctx context.Context, req *APILabelPresetUpdatePostReq) (r APILabelPresetUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APILabelSetPost implements POST /api/label/set operation.
//
// Установка метки на книгу или страницу.
//
// POST /api/label/set
func (UnimplementedHandler) APILabelSetPost(ctx context.Context, req *APILabelSetPostReq) (r APILabelSetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingHandlePost implements POST /api/parsing/handle operation.
//
// Обработка ссылок на новые книги.
//
// POST /api/parsing/handle
func (UnimplementedHandler) APIParsingHandlePost(ctx context.Context, req *APIParsingHandlePostReq) (r APIParsingHandlePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingMirrorCreatePost implements POST /api/parsing/mirror/create operation.
//
// Создание данных зеркала.
//
// POST /api/parsing/mirror/create
func (UnimplementedHandler) APIParsingMirrorCreatePost(ctx context.Context, req *APIParsingMirrorCreatePostReq) (r APIParsingMirrorCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingMirrorDeletePost implements POST /api/parsing/mirror/delete operation.
//
// Удаление зеркала.
//
// POST /api/parsing/mirror/delete
func (UnimplementedHandler) APIParsingMirrorDeletePost(ctx context.Context, req *APIParsingMirrorDeletePostReq) (r APIParsingMirrorDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingMirrorGetPost implements POST /api/parsing/mirror/get operation.
//
// Получение данных зеркала.
//
// POST /api/parsing/mirror/get
func (UnimplementedHandler) APIParsingMirrorGetPost(ctx context.Context, req *APIParsingMirrorGetPostReq) (r APIParsingMirrorGetPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingMirrorListGet implements GET /api/parsing/mirror/list operation.
//
// Зеркала.
//
// GET /api/parsing/mirror/list
func (UnimplementedHandler) APIParsingMirrorListGet(ctx context.Context) (r APIParsingMirrorListGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingMirrorUpdatePost implements POST /api/parsing/mirror/update operation.
//
// Обновления зеркала.
//
// POST /api/parsing/mirror/update
func (UnimplementedHandler) APIParsingMirrorUpdatePost(ctx context.Context, req *APIParsingMirrorUpdatePostReq) (r APIParsingMirrorUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemImportArchivePost implements POST /api/system/import/archive operation.
//
// Импорт новой книги через архив.
//
// POST /api/system/import/archive
func (UnimplementedHandler) APISystemImportArchivePost(ctx context.Context, req APISystemImportArchivePostReq) (r APISystemImportArchivePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemInfoSizeGet implements GET /api/system/info/size operation.
//
// Получение общей информации о системе.
//
// GET /api/system/info/size
func (UnimplementedHandler) APISystemInfoSizeGet(ctx context.Context) (r APISystemInfoSizeGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemInfoWorkersGet implements GET /api/system/info/workers operation.
//
// Получение информации о воркерах в системе.
//
// GET /api/system/info/workers
func (UnimplementedHandler) APISystemInfoWorkersGet(ctx context.Context) (r APISystemInfoWorkersGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemTaskCreatePost implements POST /api/system/task/create operation.
//
// Создание и фоновый запуск задачи.
//
// POST /api/system/task/create
func (UnimplementedHandler) APISystemTaskCreatePost(ctx context.Context, req *APISystemTaskCreatePostReq) (r APISystemTaskCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemTaskResultsGet implements GET /api/system/task/results operation.
//
// Получение результатов задач.
//
// GET /api/system/task/results
func (UnimplementedHandler) APISystemTaskResultsGet(ctx context.Context) (r APISystemTaskResultsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemWorkerConfigPost implements POST /api/system/worker/config operation.
//
// Динамическая конфигурация раннеров (воркеров),
// сбрасывается при перезапуске системы.
//
// POST /api/system/worker/config
func (UnimplementedHandler) APISystemWorkerConfigPost(ctx context.Context, req *APISystemWorkerConfigPostReq) (r APISystemWorkerConfigPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIUserLoginPost implements POST /api/user/login operation.
//
// Проставление токена в куки.
//
// POST /api/user/login
func (UnimplementedHandler) APIUserLoginPost(ctx context.Context, req *APIUserLoginPostReq) (r APIUserLoginPostRes, _ error) {
	return r, ht.ErrNotImplemented
}
