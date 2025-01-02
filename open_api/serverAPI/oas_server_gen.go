// Code generated by ogen, DO NOT EDIT.

package serverAPI

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
	// Удаление книги (без удаления метаинформации).
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
	// APIBookUpdatePost implements POST /api/book/update operation.
	//
	// Изменяет часть данных книги, ряд полей не изменяется
	// (верификация, число страниц и т.д.).
	//
	// POST /api/book/update
	APIBookUpdatePost(ctx context.Context, req *BookRaw) (APIBookUpdatePostRes, error)
	// APIBookVerifyPost implements POST /api/book/verify operation.
	//
	// Подтверждение (модерация) книги, нужна в случае
	// массовой обработки.
	//
	// POST /api/book/verify
	APIBookVerifyPost(ctx context.Context, req *APIBookVerifyPostReq) (APIBookVerifyPostRes, error)
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
	// APIPageBodyPost implements POST /api/page/body operation.
	//
	// Получение тела страницы (по оригинальному адресу или
	// данным книги).
	//
	// POST /api/page/body
	APIPageBodyPost(ctx context.Context, req *APIPageBodyPostReq) (APIPageBodyPostRes, error)
	// APIParsingBookExistsPost implements POST /api/parsing/book/exists operation.
	//
	// Проверка наличия ссылок на книги.
	//
	// POST /api/parsing/book/exists
	APIParsingBookExistsPost(ctx context.Context, req *APIParsingBookExistsPostReq) (APIParsingBookExistsPostRes, error)
	// APIParsingPageExistsPost implements POST /api/parsing/page/exists operation.
	//
	// Проверка наличия ссылок для страниц.
	//
	// POST /api/parsing/page/exists
	APIParsingPageExistsPost(ctx context.Context, req *APIParsingPageExistsPostReq) (APIParsingPageExistsPostRes, error)
	// APISystemDeduplicateArchivePost implements POST /api/system/deduplicate/archive operation.
	//
	// Проверка наличия данных в системе из архива.
	//
	// POST /api/system/deduplicate/archive
	APISystemDeduplicateArchivePost(ctx context.Context, req APISystemDeduplicateArchivePostReq) (APISystemDeduplicateArchivePostRes, error)
	// APISystemHandlePost implements POST /api/system/handle operation.
	//
	// Обработка ссылок на новые книги.
	//
	// POST /api/system/handle
	APISystemHandlePost(ctx context.Context, req *APISystemHandlePostReq) (APISystemHandlePostRes, error)
	// APISystemImportArchivePost implements POST /api/system/import/archive operation.
	//
	// Импорт новой книги через архив.
	//
	// POST /api/system/import/archive
	APISystemImportArchivePost(ctx context.Context, req APISystemImportArchivePostReq) (APISystemImportArchivePostRes, error)
	// APISystemInfoGet implements GET /api/system/info operation.
	//
	// Получение общей информации о системе.
	//
	// GET /api/system/info
	APISystemInfoGet(ctx context.Context) (APISystemInfoGetRes, error)
	// APISystemWorkerConfigPost implements POST /api/system/worker/config operation.
	//
	// Динамическая конфигурация раннеров (воркеров),
	// сбрасывается при перезапуске системы.
	//
	// POST /api/system/worker/config
	APISystemWorkerConfigPost(ctx context.Context, req *APISystemWorkerConfigPostReq) (APISystemWorkerConfigPostRes, error)
	// APITaskCreatePost implements POST /api/task/create operation.
	//
	// Создание и фоновый запуск задачи.
	//
	// POST /api/task/create
	APITaskCreatePost(ctx context.Context, req *APITaskCreatePostReq) (APITaskCreatePostRes, error)
	// APITaskResultsGet implements GET /api/task/results operation.
	//
	// Получение результатов задач.
	//
	// GET /api/task/results
	APITaskResultsGet(ctx context.Context) (APITaskResultsGetRes, error)
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
