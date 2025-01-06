// Code generated by ogen, DO NOT EDIT.

package serverAPI

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

// APIAttributeCountGet implements GET /api/attribute/count operation.
//
// Получение информации о количестве вариантов
// атрибутов.
//
// GET /api/attribute/count
func (UnimplementedHandler) APIAttributeCountGet(ctx context.Context) (r APIAttributeCountGetRes, _ error) {
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
// Удаление книги (без удаления метаинформации).
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

// APIBookUpdatePost implements POST /api/book/update operation.
//
// Изменяет часть данных книги, ряд полей не изменяется
// (верификация, число страниц и т.д.).
//
// POST /api/book/update
func (UnimplementedHandler) APIBookUpdatePost(ctx context.Context, req *BookRaw) (r APIBookUpdatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIBookVerifyPost implements POST /api/book/verify operation.
//
// Подтверждение (модерация) книги, нужна в случае
// массовой обработки.
//
// POST /api/book/verify
func (UnimplementedHandler) APIBookVerifyPost(ctx context.Context, req *APIBookVerifyPostReq) (r APIBookVerifyPostRes, _ error) {
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

// APIDeduplicateDeadHashByBookPagesCreatePost implements POST /api/deduplicate/dead-hash-by-book-pages/create operation.
//
// Создает запись о мертвом хеше по страницам книги.
//
// POST /api/deduplicate/dead-hash-by-book-pages/create
func (UnimplementedHandler) APIDeduplicateDeadHashByBookPagesCreatePost(ctx context.Context, req *APIDeduplicateDeadHashByBookPagesCreatePostReq) (r APIDeduplicateDeadHashByBookPagesCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateDeadHashByBookPagesDeletePost implements POST /api/deduplicate/dead-hash-by-book-pages/delete operation.
//
// Удаляет запись о мертвом хеше по страницам книги.
//
// POST /api/deduplicate/dead-hash-by-book-pages/delete
func (UnimplementedHandler) APIDeduplicateDeadHashByBookPagesDeletePost(ctx context.Context, req *APIDeduplicateDeadHashByBookPagesDeletePostReq) (r APIDeduplicateDeadHashByBookPagesDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateDeadHashByPageCreatePost implements POST /api/deduplicate/dead-hash-by-page/create operation.
//
// Создает запись о мертвом хеше по странице.
//
// POST /api/deduplicate/dead-hash-by-page/create
func (UnimplementedHandler) APIDeduplicateDeadHashByPageCreatePost(ctx context.Context, req *APIDeduplicateDeadHashByPageCreatePostReq) (r APIDeduplicateDeadHashByPageCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateDeadHashByPageDeletePost implements POST /api/deduplicate/dead-hash-by-page/delete operation.
//
// Удаляет запись о мертвом хеше по странице.
//
// POST /api/deduplicate/dead-hash-by-page/delete
func (UnimplementedHandler) APIDeduplicateDeadHashByPageDeletePost(ctx context.Context, req *APIDeduplicateDeadHashByPageDeletePostReq) (r APIDeduplicateDeadHashByPageDeletePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIDeduplicateDeleteAllPagesByHashPost implements POST /api/deduplicate/delete-all-pages-by-hash operation.
//
// Удаляет страницы с таким же хешом как у указанной.
//
// POST /api/deduplicate/delete-all-pages-by-hash
func (UnimplementedHandler) APIDeduplicateDeleteAllPagesByHashPost(ctx context.Context, req *APIDeduplicateDeleteAllPagesByHashPostReq) (r APIDeduplicateDeleteAllPagesByHashPostRes, _ error) {
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

// APIPageBodyPost implements POST /api/page/body operation.
//
// Получение тела страницы (по оригинальному адресу или
// данным книги).
//
// POST /api/page/body
func (UnimplementedHandler) APIPageBodyPost(ctx context.Context, req *APIPageBodyPostReq) (r APIPageBodyPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingBookExistsPost implements POST /api/parsing/book/exists operation.
//
// Проверка наличия ссылок на книги.
//
// POST /api/parsing/book/exists
func (UnimplementedHandler) APIParsingBookExistsPost(ctx context.Context, req *APIParsingBookExistsPostReq) (r APIParsingBookExistsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APIParsingPageExistsPost implements POST /api/parsing/page/exists operation.
//
// Проверка наличия ссылок для страниц.
//
// POST /api/parsing/page/exists
func (UnimplementedHandler) APIParsingPageExistsPost(ctx context.Context, req *APIParsingPageExistsPostReq) (r APIParsingPageExistsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemDeduplicateArchivePost implements POST /api/system/deduplicate/archive operation.
//
// Проверка наличия данных в системе из архива.
//
// POST /api/system/deduplicate/archive
func (UnimplementedHandler) APISystemDeduplicateArchivePost(ctx context.Context, req APISystemDeduplicateArchivePostReq) (r APISystemDeduplicateArchivePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APISystemHandlePost implements POST /api/system/handle operation.
//
// Обработка ссылок на новые книги.
//
// POST /api/system/handle
func (UnimplementedHandler) APISystemHandlePost(ctx context.Context, req *APISystemHandlePostReq) (r APISystemHandlePostRes, _ error) {
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

// APISystemInfoGet implements GET /api/system/info operation.
//
// Получение общей информации о системе.
//
// GET /api/system/info
func (UnimplementedHandler) APISystemInfoGet(ctx context.Context) (r APISystemInfoGetRes, _ error) {
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

// APITaskCreatePost implements POST /api/task/create operation.
//
// Создание и фоновый запуск задачи.
//
// POST /api/task/create
func (UnimplementedHandler) APITaskCreatePost(ctx context.Context, req *APITaskCreatePostReq) (r APITaskCreatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// APITaskResultsGet implements GET /api/task/results operation.
//
// Получение результатов задач.
//
// GET /api/task/results
func (UnimplementedHandler) APITaskResultsGet(ctx context.Context) (r APITaskResultsGetRes, _ error) {
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
