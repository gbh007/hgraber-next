// Code generated by ogen, DO NOT EDIT.

package server

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
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
	APIBookListPost(ctx context.Context, req *APIBookListPostReq) (APIBookListPostRes, error)
	// APIFileIDGet implements GET /api/file/{id} operation.
	//
	// Получение тела файла (изображения страницы).
	//
	// GET /api/file/{id}
	APIFileIDGet(ctx context.Context, params APIFileIDGetParams) (APIFileIDGetRes, error)
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
