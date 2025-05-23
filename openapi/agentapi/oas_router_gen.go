// Code generated by ogen, DO NOT EDIT.

package agentapi

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [2]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/api/"
			origElem := elem
			if l := len("/api/"); len(elem) >= l && elem[0:l] == "/api/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "core/status"
				origElem := elem
				if l := len("core/status"); len(elem) >= l && elem[0:l] == "core/status" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleAPICoreStatusGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 'f': // Prefix: "fs/"
				origElem := elem
				if l := len("fs/"); len(elem) >= l && elem[0:l] == "fs/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					origElem := elem
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIFsCreatePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'd': // Prefix: "delete"
					origElem := elem
					if l := len("delete"); len(elem) >= l && elem[0:l] == "delete" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIFsDeletePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'g': // Prefix: "get"
					origElem := elem
					if l := len("get"); len(elem) >= l && elem[0:l] == "get" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAPIFsGetGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'i': // Prefix: "info"
					origElem := elem
					if l := len("info"); len(elem) >= l && elem[0:l] == "info" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIFsInfoPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "h"
				origElem := elem
				if l := len("h"); len(elem) >= l && elem[0:l] == "h" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "ighway/"
					origElem := elem
					if l := len("ighway/"); len(elem) >= l && elem[0:l] == "ighway/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'f': // Prefix: "file/"
						origElem := elem
						if l := len("file/"); len(elem) >= l && elem[0:l] == "file/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "."
						idx := strings.IndexByte(elem, '.')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '.': // Prefix: "."
							origElem := elem
							if l := len("."); len(elem) >= l && elem[0:l] == "." {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "ext"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPIHighwayFileIDExtGetRequest([2]string{
										args[0],
										args[1],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					case 't': // Prefix: "token/create"
						origElem := elem
						if l := len("token/create"); len(elem) >= l && elem[0:l] == "token/create" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIHighwayTokenCreatePostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "proxy/parse/"
					origElem := elem
					if l := len("proxy/parse/"); len(elem) >= l && elem[0:l] == "proxy/parse/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "book"
						origElem := elem
						if l := len("book"); len(elem) >= l && elem[0:l] == "book" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIHproxyParseBookPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 'l': // Prefix: "list"
						origElem := elem
						if l := len("list"); len(elem) >= l && elem[0:l] == "list" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIHproxyParseListPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'i': // Prefix: "import/archive"
				origElem := elem
				if l := len("import/archive"); len(elem) >= l && elem[0:l] == "import/archive" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleAPIImportArchivePostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

				elem = origElem
			case 'p': // Prefix: "parsing/"
				origElem := elem
				if l := len("parsing/"); len(elem) >= l && elem[0:l] == "parsing/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "book"
					origElem := elem
					if l := len("book"); len(elem) >= l && elem[0:l] == "book" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "POST":
							s.handleAPIParsingBookPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "check"
							origElem := elem
							if l := len("check"); len(elem) >= l && elem[0:l] == "check" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleAPIParsingBookCheckPostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						case 'm': // Prefix: "multi"
							origElem := elem
							if l := len("multi"); len(elem) >= l && elem[0:l] == "multi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleAPIParsingBookMultiPostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "page"
					origElem := elem
					if l := len("page"); len(elem) >= l && elem[0:l] == "page" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "POST":
							s.handleAPIParsingPagePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/check"
						origElem := elem
						if l := len("/check"); len(elem) >= l && elem[0:l] == "/check" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIParsingPageCheckPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [2]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/api/"
			origElem := elem
			if l := len("/api/"); len(elem) >= l && elem[0:l] == "/api/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "core/status"
				origElem := elem
				if l := len("core/status"); len(elem) >= l && elem[0:l] == "core/status" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = APICoreStatusGetOperation
						r.summary = "Получение данных о состоянии агента"
						r.operationID = ""
						r.pathPattern = "/api/core/status"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'f': // Prefix: "fs/"
				origElem := elem
				if l := len("fs/"); len(elem) >= l && elem[0:l] == "fs/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "create"
					origElem := elem
					if l := len("create"); len(elem) >= l && elem[0:l] == "create" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = APIFsCreatePostOperation
							r.summary = "Создание нового файла"
							r.operationID = ""
							r.pathPattern = "/api/fs/create"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'd': // Prefix: "delete"
					origElem := elem
					if l := len("delete"); len(elem) >= l && elem[0:l] == "delete" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = APIFsDeletePostOperation
							r.summary = "Удаление файла"
							r.operationID = ""
							r.pathPattern = "/api/fs/delete"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'g': // Prefix: "get"
					origElem := elem
					if l := len("get"); len(elem) >= l && elem[0:l] == "get" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = APIFsGetGetOperation
							r.summary = "Получение файла"
							r.operationID = ""
							r.pathPattern = "/api/fs/get"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'i': // Prefix: "info"
					origElem := elem
					if l := len("info"); len(elem) >= l && elem[0:l] == "info" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = APIFsInfoPostOperation
							r.summary = "Получение информации о состоянии файловой системы"
							r.operationID = ""
							r.pathPattern = "/api/fs/info"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "h"
				origElem := elem
				if l := len("h"); len(elem) >= l && elem[0:l] == "h" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "ighway/"
					origElem := elem
					if l := len("ighway/"); len(elem) >= l && elem[0:l] == "ighway/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'f': // Prefix: "file/"
						origElem := elem
						if l := len("file/"); len(elem) >= l && elem[0:l] == "file/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "id"
						// Match until "."
						idx := strings.IndexByte(elem, '.')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '.': // Prefix: "."
							origElem := elem
							if l := len("."); len(elem) >= l && elem[0:l] == "." {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "ext"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APIHighwayFileIDExtGetOperation
									r.summary = "Получение файла через highway"
									r.operationID = ""
									r.pathPattern = "/api/highway/file/{id}.{ext}"
									r.args = args
									r.count = 2
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					case 't': // Prefix: "token/create"
						origElem := elem
						if l := len("token/create"); len(elem) >= l && elem[0:l] == "token/create" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = APIHighwayTokenCreatePostOperation
								r.summary = "Создание нового токена для highway"
								r.operationID = ""
								r.pathPattern = "/api/highway/token/create"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "proxy/parse/"
					origElem := elem
					if l := len("proxy/parse/"); len(elem) >= l && elem[0:l] == "proxy/parse/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "book"
						origElem := elem
						if l := len("book"); len(elem) >= l && elem[0:l] == "book" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = APIHproxyParseBookPostOperation
								r.summary = "Парсинг данных книги по ссылке"
								r.operationID = ""
								r.pathPattern = "/api/hproxy/parse/book"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'l': // Prefix: "list"
						origElem := elem
						if l := len("list"); len(elem) >= l && elem[0:l] == "list" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = APIHproxyParseListPostOperation
								r.summary = "Парсинг списка данных по ссылке"
								r.operationID = ""
								r.pathPattern = "/api/hproxy/parse/list"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'i': // Prefix: "import/archive"
				origElem := elem
				if l := len("import/archive"); len(elem) >= l && elem[0:l] == "import/archive" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = APIImportArchivePostOperation
						r.summary = "Загрузка архива"
						r.operationID = ""
						r.pathPattern = "/api/import/archive"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'p': // Prefix: "parsing/"
				origElem := elem
				if l := len("parsing/"); len(elem) >= l && elem[0:l] == "parsing/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'b': // Prefix: "book"
					origElem := elem
					if l := len("book"); len(elem) >= l && elem[0:l] == "book" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							r.name = APIParsingBookPostOperation
							r.summary = "Обработка новой книги"
							r.operationID = ""
							r.pathPattern = "/api/parsing/book"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "check"
							origElem := elem
							if l := len("check"); len(elem) >= l && elem[0:l] == "check" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = APIParsingBookCheckPostOperation
									r.summary = "Предварительная проверка ссылок на новые книги"
									r.operationID = ""
									r.pathPattern = "/api/parsing/book/check"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'm': // Prefix: "multi"
							origElem := elem
							if l := len("multi"); len(elem) >= l && elem[0:l] == "multi" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = APIParsingBookMultiPostOperation
									r.summary = "Обработка ссылки с набором книг"
									r.operationID = ""
									r.pathPattern = "/api/parsing/book/multi"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				case 'p': // Prefix: "page"
					origElem := elem
					if l := len("page"); len(elem) >= l && elem[0:l] == "page" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							r.name = APIParsingPagePostOperation
							r.summary = "Загрузка изображения страницы"
							r.operationID = ""
							r.pathPattern = "/api/parsing/page"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/check"
						origElem := elem
						if l := len("/check"); len(elem) >= l && elem[0:l] == "/check" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = APIParsingPageCheckPostOperation
								r.summary = "Предварительная проверка ссылок для загрузки страниц"
								r.operationID = ""
								r.pathPattern = "/api/parsing/page/check"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
