// Code generated by ogen, DO NOT EDIT.

package serverAPI

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
	args := [1]string{}

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
			case 'a': // Prefix: "agent/"
				origElem := elem
				if l := len("agent/"); len(elem) >= l && elem[0:l] == "agent/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
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
							s.handleAPIAgentDeletePostRequest([0]string{}, elemIsEscaped, w, r)
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
							s.handleAPIAgentListPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'n': // Prefix: "new"
					origElem := elem
					if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIAgentNewPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 't': // Prefix: "task/export"
					origElem := elem
					if l := len("task/export"); len(elem) >= l && elem[0:l] == "task/export" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIAgentTaskExportPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'b': // Prefix: "book/"
				origElem := elem
				if l := len("book/"); len(elem) >= l && elem[0:l] == "book/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "archive/"
					origElem := elem
					if l := len("archive/"); len(elem) >= l && elem[0:l] == "archive/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAPIBookArchiveIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'd': // Prefix: "de"
					origElem := elem
					if l := len("de"); len(elem) >= l && elem[0:l] == "de" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "lete"
						origElem := elem
						if l := len("lete"); len(elem) >= l && elem[0:l] == "lete" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIBookDeletePostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 't': // Prefix: "tails"
						origElem := elem
						if l := len("tails"); len(elem) >= l && elem[0:l] == "tails" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIBookDetailsPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
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
							s.handleAPIBookListPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'r': // Prefix: "raw"
					origElem := elem
					if l := len("raw"); len(elem) >= l && elem[0:l] == "raw" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIBookRawPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'v': // Prefix: "verify"
					origElem := elem
					if l := len("verify"); len(elem) >= l && elem[0:l] == "verify" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIBookVerifyPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'f': // Prefix: "file/"
				origElem := elem
				if l := len("file/"); len(elem) >= l && elem[0:l] == "file/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleAPIFileIDGetRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 'p': // Prefix: "pa"
				origElem := elem
				if l := len("pa"); len(elem) >= l && elem[0:l] == "pa" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'g': // Prefix: "ge/body"
					origElem := elem
					if l := len("ge/body"); len(elem) >= l && elem[0:l] == "ge/body" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPIPageBodyPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'r': // Prefix: "rsing/"
					origElem := elem
					if l := len("rsing/"); len(elem) >= l && elem[0:l] == "rsing/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "book/exists"
						origElem := elem
						if l := len("book/exists"); len(elem) >= l && elem[0:l] == "book/exists" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIParsingBookExistsPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 'p': // Prefix: "page/exists"
						origElem := elem
						if l := len("page/exists"); len(elem) >= l && elem[0:l] == "page/exists" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPIParsingPageExistsPostRequest([0]string{}, elemIsEscaped, w, r)
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
			case 's': // Prefix: "system/"
				origElem := elem
				if l := len("system/"); len(elem) >= l && elem[0:l] == "system/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'd': // Prefix: "deduplicate/archive"
					origElem := elem
					if l := len("deduplicate/archive"); len(elem) >= l && elem[0:l] == "deduplicate/archive" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPISystemDeduplicateArchivePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'h': // Prefix: "handle"
					origElem := elem
					if l := len("handle"); len(elem) >= l && elem[0:l] == "handle" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPISystemHandlePostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				case 'i': // Prefix: "i"
					origElem := elem
					if l := len("i"); len(elem) >= l && elem[0:l] == "i" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'm': // Prefix: "mport/archive"
						origElem := elem
						if l := len("mport/archive"); len(elem) >= l && elem[0:l] == "mport/archive" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPISystemImportArchivePostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 'n': // Prefix: "nfo"
						origElem := elem
						if l := len("nfo"); len(elem) >= l && elem[0:l] == "nfo" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAPISystemInfoGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'r': // Prefix: "rpc/"
					origElem := elem
					if l := len("rpc/"); len(elem) >= l && elem[0:l] == "rpc/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'd': // Prefix: "deduplicate/files"
						origElem := elem
						if l := len("deduplicate/files"); len(elem) >= l && elem[0:l] == "deduplicate/files" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAPISystemRPCDeduplicateFilesPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

						elem = origElem
					case 'r': // Prefix: "remove/"
						origElem := elem
						if l := len("remove/"); len(elem) >= l && elem[0:l] == "remove/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'd': // Prefix: "detached-files"
							origElem := elem
							if l := len("detached-files"); len(elem) >= l && elem[0:l] == "detached-files" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleAPISystemRPCRemoveDetachedFilesPostRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}

							elem = origElem
						case 'm': // Prefix: "mismatch-files"
							origElem := elem
							if l := len("mismatch-files"); len(elem) >= l && elem[0:l] == "mismatch-files" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleAPISystemRPCRemoveMismatchFilesPostRequest([0]string{}, elemIsEscaped, w, r)
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
				case 'w': // Prefix: "worker/config"
					origElem := elem
					if l := len("worker/config"); len(elem) >= l && elem[0:l] == "worker/config" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleAPISystemWorkerConfigPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'u': // Prefix: "user/login"
				origElem := elem
				if l := len("user/login"); len(elem) >= l && elem[0:l] == "user/login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleAPIUserLoginPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
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
	args        [1]string
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
			case 'a': // Prefix: "agent/"
				origElem := elem
				if l := len("agent/"); len(elem) >= l && elem[0:l] == "agent/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
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
							r.name = "APIAgentDeletePost"
							r.summary = "Удаление агента"
							r.operationID = ""
							r.pathPattern = "/api/agent/delete"
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
							r.name = "APIAgentListPost"
							r.summary = "Список агентов"
							r.operationID = ""
							r.pathPattern = "/api/agent/list"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'n': // Prefix: "new"
					origElem := elem
					if l := len("new"); len(elem) >= l && elem[0:l] == "new" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APIAgentNewPost"
							r.summary = "Создание нового агента"
							r.operationID = ""
							r.pathPattern = "/api/agent/new"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 't': // Prefix: "task/export"
					origElem := elem
					if l := len("task/export"); len(elem) >= l && elem[0:l] == "task/export" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APIAgentTaskExportPost"
							r.summary = "Экспорт книг в другую систему"
							r.operationID = ""
							r.pathPattern = "/api/agent/task/export"
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
			case 'b': // Prefix: "book/"
				origElem := elem
				if l := len("book/"); len(elem) >= l && elem[0:l] == "book/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "archive/"
					origElem := elem
					if l := len("archive/"); len(elem) >= l && elem[0:l] == "archive/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = "APIBookArchiveIDGet"
							r.summary = "Получение архива с книгой"
							r.operationID = ""
							r.pathPattern = "/api/book/archive/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'd': // Prefix: "de"
					origElem := elem
					if l := len("de"); len(elem) >= l && elem[0:l] == "de" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "lete"
						origElem := elem
						if l := len("lete"); len(elem) >= l && elem[0:l] == "lete" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "APIBookDeletePost"
								r.summary = "Удаление книги"
								r.operationID = ""
								r.pathPattern = "/api/book/delete"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 't': // Prefix: "tails"
						origElem := elem
						if l := len("tails"); len(elem) >= l && elem[0:l] == "tails" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "APIBookDetailsPost"
								r.summary = "Информация о книге"
								r.operationID = ""
								r.pathPattern = "/api/book/details"
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
							r.name = "APIBookListPost"
							r.summary = "Список книг"
							r.operationID = ""
							r.pathPattern = "/api/book/list"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'r': // Prefix: "raw"
					origElem := elem
					if l := len("raw"); len(elem) >= l && elem[0:l] == "raw" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APIBookRawPost"
							r.summary = "Информация о книге"
							r.operationID = ""
							r.pathPattern = "/api/book/raw"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'v': // Prefix: "verify"
					origElem := elem
					if l := len("verify"); len(elem) >= l && elem[0:l] == "verify" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APIBookVerifyPost"
							r.summary = "Подтверждение (модерация) книги"
							r.operationID = ""
							r.pathPattern = "/api/book/verify"
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
			case 'f': // Prefix: "file/"
				origElem := elem
				if l := len("file/"); len(elem) >= l && elem[0:l] == "file/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = "APIFileIDGet"
						r.summary = "Получение тела файла"
						r.operationID = ""
						r.pathPattern = "/api/file/{id}"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'p': // Prefix: "pa"
				origElem := elem
				if l := len("pa"); len(elem) >= l && elem[0:l] == "pa" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'g': // Prefix: "ge/body"
					origElem := elem
					if l := len("ge/body"); len(elem) >= l && elem[0:l] == "ge/body" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APIPageBodyPost"
							r.summary = "Получение тела страницы"
							r.operationID = ""
							r.pathPattern = "/api/page/body"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'r': // Prefix: "rsing/"
					origElem := elem
					if l := len("rsing/"); len(elem) >= l && elem[0:l] == "rsing/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "book/exists"
						origElem := elem
						if l := len("book/exists"); len(elem) >= l && elem[0:l] == "book/exists" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "APIParsingBookExistsPost"
								r.summary = "Проверка наличия ссылок на книги"
								r.operationID = ""
								r.pathPattern = "/api/parsing/book/exists"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'p': // Prefix: "page/exists"
						origElem := elem
						if l := len("page/exists"); len(elem) >= l && elem[0:l] == "page/exists" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "APIParsingPageExistsPost"
								r.summary = "Проверка наличия ссылок для страниц"
								r.operationID = ""
								r.pathPattern = "/api/parsing/page/exists"
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
			case 's': // Prefix: "system/"
				origElem := elem
				if l := len("system/"); len(elem) >= l && elem[0:l] == "system/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'd': // Prefix: "deduplicate/archive"
					origElem := elem
					if l := len("deduplicate/archive"); len(elem) >= l && elem[0:l] == "deduplicate/archive" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APISystemDeduplicateArchivePost"
							r.summary = "Проверка наличия данных в системе из архива"
							r.operationID = ""
							r.pathPattern = "/api/system/deduplicate/archive"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'h': // Prefix: "handle"
					origElem := elem
					if l := len("handle"); len(elem) >= l && elem[0:l] == "handle" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APISystemHandlePost"
							r.summary = "Обработка ссылок на новые книги"
							r.operationID = ""
							r.pathPattern = "/api/system/handle"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'i': // Prefix: "i"
					origElem := elem
					if l := len("i"); len(elem) >= l && elem[0:l] == "i" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'm': // Prefix: "mport/archive"
						origElem := elem
						if l := len("mport/archive"); len(elem) >= l && elem[0:l] == "mport/archive" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "APISystemImportArchivePost"
								r.summary = "Импорт новой книги"
								r.operationID = ""
								r.pathPattern = "/api/system/import/archive"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'n': // Prefix: "nfo"
						origElem := elem
						if l := len("nfo"); len(elem) >= l && elem[0:l] == "nfo" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = "APISystemInfoGet"
								r.summary = "Текущее состояние системы"
								r.operationID = ""
								r.pathPattern = "/api/system/info"
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
				case 'r': // Prefix: "rpc/"
					origElem := elem
					if l := len("rpc/"); len(elem) >= l && elem[0:l] == "rpc/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'd': // Prefix: "deduplicate/files"
						origElem := elem
						if l := len("deduplicate/files"); len(elem) >= l && elem[0:l] == "deduplicate/files" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = "APISystemRPCDeduplicateFilesPost"
								r.summary = "Дедупликация файлов"
								r.operationID = ""
								r.pathPattern = "/api/system/rpc/deduplicate/files"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'r': // Prefix: "remove/"
						origElem := elem
						if l := len("remove/"); len(elem) >= l && elem[0:l] == "remove/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'd': // Prefix: "detached-files"
							origElem := elem
							if l := len("detached-files"); len(elem) >= l && elem[0:l] == "detached-files" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = "APISystemRPCRemoveDetachedFilesPost"
									r.summary = "Удаление несвязанных файлов"
									r.operationID = ""
									r.pathPattern = "/api/system/rpc/remove/detached-files"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'm': // Prefix: "mismatch-files"
							origElem := elem
							if l := len("mismatch-files"); len(elem) >= l && elem[0:l] == "mismatch-files" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "POST":
									r.name = "APISystemRPCRemoveMismatchFilesPost"
									r.summary = "Удаление рассинхронизированных файлов"
									r.operationID = ""
									r.pathPattern = "/api/system/rpc/remove/mismatch-files"
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
				case 'w': // Prefix: "worker/config"
					origElem := elem
					if l := len("worker/config"); len(elem) >= l && elem[0:l] == "worker/config" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = "APISystemWorkerConfigPost"
							r.summary = "Динамическая конфигурация раннеров (воркеров)"
							r.operationID = ""
							r.pathPattern = "/api/system/worker/config"
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
			case 'u': // Prefix: "user/login"
				origElem := elem
				if l := len("user/login"); len(elem) >= l && elem[0:l] == "user/login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = "APIUserLoginPost"
						r.summary = "Проставление токена в куки"
						r.operationID = ""
						r.pathPattern = "/api/user/login"
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
	}
	return r, false
}
