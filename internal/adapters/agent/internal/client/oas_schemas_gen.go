// Code generated by ogen, DO NOT EDIT.

package client

import (
	"io"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

type APICoreStatusGetBadRequest ErrorResponse

func (*APICoreStatusGetBadRequest) aPICoreStatusGetRes() {}

type APICoreStatusGetForbidden ErrorResponse

func (*APICoreStatusGetForbidden) aPICoreStatusGetRes() {}

type APICoreStatusGetInternalServerError ErrorResponse

func (*APICoreStatusGetInternalServerError) aPICoreStatusGetRes() {}

type APICoreStatusGetOK struct {
	// Время запуска агента.
	StartAt time.Time `json:"start_at"`
	// Текущее состояние агента.
	Status APICoreStatusGetOKStatus `json:"status"`
	// Список проблем.
	Problems []APICoreStatusGetOKProblemsItem `json:"problems"`
}

// GetStartAt returns the value of StartAt.
func (s *APICoreStatusGetOK) GetStartAt() time.Time {
	return s.StartAt
}

// GetStatus returns the value of Status.
func (s *APICoreStatusGetOK) GetStatus() APICoreStatusGetOKStatus {
	return s.Status
}

// GetProblems returns the value of Problems.
func (s *APICoreStatusGetOK) GetProblems() []APICoreStatusGetOKProblemsItem {
	return s.Problems
}

// SetStartAt sets the value of StartAt.
func (s *APICoreStatusGetOK) SetStartAt(val time.Time) {
	s.StartAt = val
}

// SetStatus sets the value of Status.
func (s *APICoreStatusGetOK) SetStatus(val APICoreStatusGetOKStatus) {
	s.Status = val
}

// SetProblems sets the value of Problems.
func (s *APICoreStatusGetOK) SetProblems(val []APICoreStatusGetOKProblemsItem) {
	s.Problems = val
}

func (*APICoreStatusGetOK) aPICoreStatusGetRes() {}

type APICoreStatusGetOKProblemsItem struct {
	// Тип проблемы.
	Type APICoreStatusGetOKProblemsItemType `json:"type"`
	// Описание проблемы.
	Details string `json:"details"`
}

// GetType returns the value of Type.
func (s *APICoreStatusGetOKProblemsItem) GetType() APICoreStatusGetOKProblemsItemType {
	return s.Type
}

// GetDetails returns the value of Details.
func (s *APICoreStatusGetOKProblemsItem) GetDetails() string {
	return s.Details
}

// SetType sets the value of Type.
func (s *APICoreStatusGetOKProblemsItem) SetType(val APICoreStatusGetOKProblemsItemType) {
	s.Type = val
}

// SetDetails sets the value of Details.
func (s *APICoreStatusGetOKProblemsItem) SetDetails(val string) {
	s.Details = val
}

// Тип проблемы.
type APICoreStatusGetOKProblemsItemType string

const (
	APICoreStatusGetOKProblemsItemTypeInfo    APICoreStatusGetOKProblemsItemType = "info"
	APICoreStatusGetOKProblemsItemTypeWarning APICoreStatusGetOKProblemsItemType = "warning"
	APICoreStatusGetOKProblemsItemTypeError   APICoreStatusGetOKProblemsItemType = "error"
)

// AllValues returns all APICoreStatusGetOKProblemsItemType values.
func (APICoreStatusGetOKProblemsItemType) AllValues() []APICoreStatusGetOKProblemsItemType {
	return []APICoreStatusGetOKProblemsItemType{
		APICoreStatusGetOKProblemsItemTypeInfo,
		APICoreStatusGetOKProblemsItemTypeWarning,
		APICoreStatusGetOKProblemsItemTypeError,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s APICoreStatusGetOKProblemsItemType) MarshalText() ([]byte, error) {
	switch s {
	case APICoreStatusGetOKProblemsItemTypeInfo:
		return []byte(s), nil
	case APICoreStatusGetOKProblemsItemTypeWarning:
		return []byte(s), nil
	case APICoreStatusGetOKProblemsItemTypeError:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *APICoreStatusGetOKProblemsItemType) UnmarshalText(data []byte) error {
	switch APICoreStatusGetOKProblemsItemType(data) {
	case APICoreStatusGetOKProblemsItemTypeInfo:
		*s = APICoreStatusGetOKProblemsItemTypeInfo
		return nil
	case APICoreStatusGetOKProblemsItemTypeWarning:
		*s = APICoreStatusGetOKProblemsItemTypeWarning
		return nil
	case APICoreStatusGetOKProblemsItemTypeError:
		*s = APICoreStatusGetOKProblemsItemTypeError
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Текущее состояние агента.
type APICoreStatusGetOKStatus string

const (
	APICoreStatusGetOKStatusOk      APICoreStatusGetOKStatus = "ok"
	APICoreStatusGetOKStatusWarning APICoreStatusGetOKStatus = "warning"
	APICoreStatusGetOKStatusError   APICoreStatusGetOKStatus = "error"
)

// AllValues returns all APICoreStatusGetOKStatus values.
func (APICoreStatusGetOKStatus) AllValues() []APICoreStatusGetOKStatus {
	return []APICoreStatusGetOKStatus{
		APICoreStatusGetOKStatusOk,
		APICoreStatusGetOKStatusWarning,
		APICoreStatusGetOKStatusError,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s APICoreStatusGetOKStatus) MarshalText() ([]byte, error) {
	switch s {
	case APICoreStatusGetOKStatusOk:
		return []byte(s), nil
	case APICoreStatusGetOKStatusWarning:
		return []byte(s), nil
	case APICoreStatusGetOKStatusError:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *APICoreStatusGetOKStatus) UnmarshalText(data []byte) error {
	switch APICoreStatusGetOKStatus(data) {
	case APICoreStatusGetOKStatusOk:
		*s = APICoreStatusGetOKStatusOk
		return nil
	case APICoreStatusGetOKStatusWarning:
		*s = APICoreStatusGetOKStatusWarning
		return nil
	case APICoreStatusGetOKStatusError:
		*s = APICoreStatusGetOKStatusError
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type APICoreStatusGetUnauthorized ErrorResponse

func (*APICoreStatusGetUnauthorized) aPICoreStatusGetRes() {}

type APIExportArchivePostBadRequest ErrorResponse

func (*APIExportArchivePostBadRequest) aPIExportArchivePostRes() {}

type APIExportArchivePostForbidden ErrorResponse

func (*APIExportArchivePostForbidden) aPIExportArchivePostRes() {}

type APIExportArchivePostInternalServerError ErrorResponse

func (*APIExportArchivePostInternalServerError) aPIExportArchivePostRes() {}

// APIExportArchivePostNoContent is response for APIExportArchivePost operation.
type APIExportArchivePostNoContent struct{}

func (*APIExportArchivePostNoContent) aPIExportArchivePostRes() {}

type APIExportArchivePostReq struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s APIExportArchivePostReq) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

type APIExportArchivePostUnauthorized ErrorResponse

func (*APIExportArchivePostUnauthorized) aPIExportArchivePostRes() {}

type APIFsCreatePostBadRequest ErrorResponse

func (*APIFsCreatePostBadRequest) aPIFsCreatePostRes() {}

type APIFsCreatePostConflict ErrorResponse

func (*APIFsCreatePostConflict) aPIFsCreatePostRes() {}

type APIFsCreatePostForbidden ErrorResponse

func (*APIFsCreatePostForbidden) aPIFsCreatePostRes() {}

type APIFsCreatePostInternalServerError ErrorResponse

func (*APIFsCreatePostInternalServerError) aPIFsCreatePostRes() {}

// APIFsCreatePostNoContent is response for APIFsCreatePost operation.
type APIFsCreatePostNoContent struct{}

func (*APIFsCreatePostNoContent) aPIFsCreatePostRes() {}

type APIFsCreatePostReq struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s APIFsCreatePostReq) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

type APIFsCreatePostUnauthorized ErrorResponse

func (*APIFsCreatePostUnauthorized) aPIFsCreatePostRes() {}

type APIFsDeletePostBadRequest ErrorResponse

func (*APIFsDeletePostBadRequest) aPIFsDeletePostRes() {}

type APIFsDeletePostForbidden ErrorResponse

func (*APIFsDeletePostForbidden) aPIFsDeletePostRes() {}

type APIFsDeletePostInternalServerError ErrorResponse

func (*APIFsDeletePostInternalServerError) aPIFsDeletePostRes() {}

// APIFsDeletePostNoContent is response for APIFsDeletePost operation.
type APIFsDeletePostNoContent struct{}

func (*APIFsDeletePostNoContent) aPIFsDeletePostRes() {}

type APIFsDeletePostNotFound ErrorResponse

func (*APIFsDeletePostNotFound) aPIFsDeletePostRes() {}

type APIFsDeletePostReq struct {
	// ID файла для удаления.
	FileID uuid.UUID `json:"file_id"`
}

// GetFileID returns the value of FileID.
func (s *APIFsDeletePostReq) GetFileID() uuid.UUID {
	return s.FileID
}

// SetFileID sets the value of FileID.
func (s *APIFsDeletePostReq) SetFileID(val uuid.UUID) {
	s.FileID = val
}

type APIFsDeletePostUnauthorized ErrorResponse

func (*APIFsDeletePostUnauthorized) aPIFsDeletePostRes() {}

type APIFsGetGetBadRequest ErrorResponse

func (*APIFsGetGetBadRequest) aPIFsGetGetRes() {}

type APIFsGetGetForbidden ErrorResponse

func (*APIFsGetGetForbidden) aPIFsGetGetRes() {}

type APIFsGetGetInternalServerError ErrorResponse

func (*APIFsGetGetInternalServerError) aPIFsGetGetRes() {}

type APIFsGetGetNotFound ErrorResponse

func (*APIFsGetGetNotFound) aPIFsGetGetRes() {}

type APIFsGetGetOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s APIFsGetGetOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*APIFsGetGetOK) aPIFsGetGetRes() {}

type APIFsGetGetUnauthorized ErrorResponse

func (*APIFsGetGetUnauthorized) aPIFsGetGetRes() {}

type APIFsIdsGetBadRequest ErrorResponse

func (*APIFsIdsGetBadRequest) aPIFsIdsGetRes() {}

type APIFsIdsGetForbidden ErrorResponse

func (*APIFsIdsGetForbidden) aPIFsIdsGetRes() {}

type APIFsIdsGetInternalServerError ErrorResponse

func (*APIFsIdsGetInternalServerError) aPIFsIdsGetRes() {}

type APIFsIdsGetOKApplicationJSON []uuid.UUID

func (*APIFsIdsGetOKApplicationJSON) aPIFsIdsGetRes() {}

type APIFsIdsGetUnauthorized ErrorResponse

func (*APIFsIdsGetUnauthorized) aPIFsIdsGetRes() {}

type APIParsingBookCheckPostBadRequest ErrorResponse

func (*APIParsingBookCheckPostBadRequest) aPIParsingBookCheckPostRes() {}

type APIParsingBookCheckPostForbidden ErrorResponse

func (*APIParsingBookCheckPostForbidden) aPIParsingBookCheckPostRes() {}

type APIParsingBookCheckPostInternalServerError ErrorResponse

func (*APIParsingBookCheckPostInternalServerError) aPIParsingBookCheckPostRes() {}

type APIParsingBookCheckPostOK struct {
	// Результат обработки.
	Result []APIParsingBookCheckPostOKResultItem `json:"result"`
}

// GetResult returns the value of Result.
func (s *APIParsingBookCheckPostOK) GetResult() []APIParsingBookCheckPostOKResultItem {
	return s.Result
}

// SetResult sets the value of Result.
func (s *APIParsingBookCheckPostOK) SetResult(val []APIParsingBookCheckPostOKResultItem) {
	s.Result = val
}

func (*APIParsingBookCheckPostOK) aPIParsingBookCheckPostRes() {}

type APIParsingBookCheckPostOKResultItem struct {
	// Ссылка на внешнюю систему.
	URL url.URL `json:"url"`
	// Результат проверки.
	Result APIParsingBookCheckPostOKResultItemResult `json:"result"`
	// Данные об ошибке во время обработки.
	ErrorDetails OptString `json:"error_details"`
	// Список возможных ссылок дубликатов (зеркала внешних
	// систем и т.п.).
	PossibleDuplicates []url.URL `json:"possible_duplicates"`
}

// GetURL returns the value of URL.
func (s *APIParsingBookCheckPostOKResultItem) GetURL() url.URL {
	return s.URL
}

// GetResult returns the value of Result.
func (s *APIParsingBookCheckPostOKResultItem) GetResult() APIParsingBookCheckPostOKResultItemResult {
	return s.Result
}

// GetErrorDetails returns the value of ErrorDetails.
func (s *APIParsingBookCheckPostOKResultItem) GetErrorDetails() OptString {
	return s.ErrorDetails
}

// GetPossibleDuplicates returns the value of PossibleDuplicates.
func (s *APIParsingBookCheckPostOKResultItem) GetPossibleDuplicates() []url.URL {
	return s.PossibleDuplicates
}

// SetURL sets the value of URL.
func (s *APIParsingBookCheckPostOKResultItem) SetURL(val url.URL) {
	s.URL = val
}

// SetResult sets the value of Result.
func (s *APIParsingBookCheckPostOKResultItem) SetResult(val APIParsingBookCheckPostOKResultItemResult) {
	s.Result = val
}

// SetErrorDetails sets the value of ErrorDetails.
func (s *APIParsingBookCheckPostOKResultItem) SetErrorDetails(val OptString) {
	s.ErrorDetails = val
}

// SetPossibleDuplicates sets the value of PossibleDuplicates.
func (s *APIParsingBookCheckPostOKResultItem) SetPossibleDuplicates(val []url.URL) {
	s.PossibleDuplicates = val
}

// Результат проверки.
type APIParsingBookCheckPostOKResultItemResult string

const (
	APIParsingBookCheckPostOKResultItemResultOk          APIParsingBookCheckPostOKResultItemResult = "ok"
	APIParsingBookCheckPostOKResultItemResultUnsupported APIParsingBookCheckPostOKResultItemResult = "unsupported"
	APIParsingBookCheckPostOKResultItemResultError       APIParsingBookCheckPostOKResultItemResult = "error"
)

// AllValues returns all APIParsingBookCheckPostOKResultItemResult values.
func (APIParsingBookCheckPostOKResultItemResult) AllValues() []APIParsingBookCheckPostOKResultItemResult {
	return []APIParsingBookCheckPostOKResultItemResult{
		APIParsingBookCheckPostOKResultItemResultOk,
		APIParsingBookCheckPostOKResultItemResultUnsupported,
		APIParsingBookCheckPostOKResultItemResultError,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s APIParsingBookCheckPostOKResultItemResult) MarshalText() ([]byte, error) {
	switch s {
	case APIParsingBookCheckPostOKResultItemResultOk:
		return []byte(s), nil
	case APIParsingBookCheckPostOKResultItemResultUnsupported:
		return []byte(s), nil
	case APIParsingBookCheckPostOKResultItemResultError:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *APIParsingBookCheckPostOKResultItemResult) UnmarshalText(data []byte) error {
	switch APIParsingBookCheckPostOKResultItemResult(data) {
	case APIParsingBookCheckPostOKResultItemResultOk:
		*s = APIParsingBookCheckPostOKResultItemResultOk
		return nil
	case APIParsingBookCheckPostOKResultItemResultUnsupported:
		*s = APIParsingBookCheckPostOKResultItemResultUnsupported
		return nil
	case APIParsingBookCheckPostOKResultItemResultError:
		*s = APIParsingBookCheckPostOKResultItemResultError
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type APIParsingBookCheckPostReq struct {
	// Ссылки на внешние системы.
	Urls []url.URL `json:"urls"`
}

// GetUrls returns the value of Urls.
func (s *APIParsingBookCheckPostReq) GetUrls() []url.URL {
	return s.Urls
}

// SetUrls sets the value of Urls.
func (s *APIParsingBookCheckPostReq) SetUrls(val []url.URL) {
	s.Urls = val
}

type APIParsingBookCheckPostUnauthorized ErrorResponse

func (*APIParsingBookCheckPostUnauthorized) aPIParsingBookCheckPostRes() {}

type APIParsingBookPostBadRequest ErrorResponse

func (*APIParsingBookPostBadRequest) aPIParsingBookPostRes() {}

type APIParsingBookPostForbidden ErrorResponse

func (*APIParsingBookPostForbidden) aPIParsingBookPostRes() {}

type APIParsingBookPostInternalServerError ErrorResponse

func (*APIParsingBookPostInternalServerError) aPIParsingBookPostRes() {}

type APIParsingBookPostReq struct {
	// Ссылка на внешнюю систему.
	URL url.URL `json:"url"`
}

// GetURL returns the value of URL.
func (s *APIParsingBookPostReq) GetURL() url.URL {
	return s.URL
}

// SetURL sets the value of URL.
func (s *APIParsingBookPostReq) SetURL(val url.URL) {
	s.URL = val
}

type APIParsingBookPostUnauthorized ErrorResponse

func (*APIParsingBookPostUnauthorized) aPIParsingBookPostRes() {}

type APIParsingPageCheckPostBadRequest ErrorResponse

func (*APIParsingPageCheckPostBadRequest) aPIParsingPageCheckPostRes() {}

type APIParsingPageCheckPostForbidden ErrorResponse

func (*APIParsingPageCheckPostForbidden) aPIParsingPageCheckPostRes() {}

type APIParsingPageCheckPostInternalServerError ErrorResponse

func (*APIParsingPageCheckPostInternalServerError) aPIParsingPageCheckPostRes() {}

type APIParsingPageCheckPostOK struct {
	// Результат обработки.
	Result []APIParsingPageCheckPostOKResultItem `json:"result"`
}

// GetResult returns the value of Result.
func (s *APIParsingPageCheckPostOK) GetResult() []APIParsingPageCheckPostOKResultItem {
	return s.Result
}

// SetResult sets the value of Result.
func (s *APIParsingPageCheckPostOK) SetResult(val []APIParsingPageCheckPostOKResultItem) {
	s.Result = val
}

func (*APIParsingPageCheckPostOK) aPIParsingPageCheckPostRes() {}

type APIParsingPageCheckPostOKResultItem struct {
	// Ссылка на книгу во внешней системе.
	BookURL url.URL `json:"book_url"`
	// Ссылка на изображение во внешней системе.
	ImageURL url.URL `json:"image_url"`
	// Результат проверки.
	Result APIParsingPageCheckPostOKResultItemResult `json:"result"`
	// Данные об ошибке во время обработки.
	ErrorDetails OptString `json:"error_details"`
}

// GetBookURL returns the value of BookURL.
func (s *APIParsingPageCheckPostOKResultItem) GetBookURL() url.URL {
	return s.BookURL
}

// GetImageURL returns the value of ImageURL.
func (s *APIParsingPageCheckPostOKResultItem) GetImageURL() url.URL {
	return s.ImageURL
}

// GetResult returns the value of Result.
func (s *APIParsingPageCheckPostOKResultItem) GetResult() APIParsingPageCheckPostOKResultItemResult {
	return s.Result
}

// GetErrorDetails returns the value of ErrorDetails.
func (s *APIParsingPageCheckPostOKResultItem) GetErrorDetails() OptString {
	return s.ErrorDetails
}

// SetBookURL sets the value of BookURL.
func (s *APIParsingPageCheckPostOKResultItem) SetBookURL(val url.URL) {
	s.BookURL = val
}

// SetImageURL sets the value of ImageURL.
func (s *APIParsingPageCheckPostOKResultItem) SetImageURL(val url.URL) {
	s.ImageURL = val
}

// SetResult sets the value of Result.
func (s *APIParsingPageCheckPostOKResultItem) SetResult(val APIParsingPageCheckPostOKResultItemResult) {
	s.Result = val
}

// SetErrorDetails sets the value of ErrorDetails.
func (s *APIParsingPageCheckPostOKResultItem) SetErrorDetails(val OptString) {
	s.ErrorDetails = val
}

// Результат проверки.
type APIParsingPageCheckPostOKResultItemResult string

const (
	APIParsingPageCheckPostOKResultItemResultOk          APIParsingPageCheckPostOKResultItemResult = "ok"
	APIParsingPageCheckPostOKResultItemResultUnsupported APIParsingPageCheckPostOKResultItemResult = "unsupported"
	APIParsingPageCheckPostOKResultItemResultError       APIParsingPageCheckPostOKResultItemResult = "error"
)

// AllValues returns all APIParsingPageCheckPostOKResultItemResult values.
func (APIParsingPageCheckPostOKResultItemResult) AllValues() []APIParsingPageCheckPostOKResultItemResult {
	return []APIParsingPageCheckPostOKResultItemResult{
		APIParsingPageCheckPostOKResultItemResultOk,
		APIParsingPageCheckPostOKResultItemResultUnsupported,
		APIParsingPageCheckPostOKResultItemResultError,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s APIParsingPageCheckPostOKResultItemResult) MarshalText() ([]byte, error) {
	switch s {
	case APIParsingPageCheckPostOKResultItemResultOk:
		return []byte(s), nil
	case APIParsingPageCheckPostOKResultItemResultUnsupported:
		return []byte(s), nil
	case APIParsingPageCheckPostOKResultItemResultError:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *APIParsingPageCheckPostOKResultItemResult) UnmarshalText(data []byte) error {
	switch APIParsingPageCheckPostOKResultItemResult(data) {
	case APIParsingPageCheckPostOKResultItemResultOk:
		*s = APIParsingPageCheckPostOKResultItemResultOk
		return nil
	case APIParsingPageCheckPostOKResultItemResultUnsupported:
		*s = APIParsingPageCheckPostOKResultItemResultUnsupported
		return nil
	case APIParsingPageCheckPostOKResultItemResultError:
		*s = APIParsingPageCheckPostOKResultItemResultError
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type APIParsingPageCheckPostReq struct {
	// Ссылки на внешние системы.
	Urls []APIParsingPageCheckPostReqUrlsItem `json:"urls"`
}

// GetUrls returns the value of Urls.
func (s *APIParsingPageCheckPostReq) GetUrls() []APIParsingPageCheckPostReqUrlsItem {
	return s.Urls
}

// SetUrls sets the value of Urls.
func (s *APIParsingPageCheckPostReq) SetUrls(val []APIParsingPageCheckPostReqUrlsItem) {
	s.Urls = val
}

type APIParsingPageCheckPostReqUrlsItem struct {
	// Ссылка на книгу во внешней системе.
	BookURL url.URL `json:"book_url"`
	// Ссылка на изображение во внешней системе.
	ImageURL url.URL `json:"image_url"`
}

// GetBookURL returns the value of BookURL.
func (s *APIParsingPageCheckPostReqUrlsItem) GetBookURL() url.URL {
	return s.BookURL
}

// GetImageURL returns the value of ImageURL.
func (s *APIParsingPageCheckPostReqUrlsItem) GetImageURL() url.URL {
	return s.ImageURL
}

// SetBookURL sets the value of BookURL.
func (s *APIParsingPageCheckPostReqUrlsItem) SetBookURL(val url.URL) {
	s.BookURL = val
}

// SetImageURL sets the value of ImageURL.
func (s *APIParsingPageCheckPostReqUrlsItem) SetImageURL(val url.URL) {
	s.ImageURL = val
}

type APIParsingPageCheckPostUnauthorized ErrorResponse

func (*APIParsingPageCheckPostUnauthorized) aPIParsingPageCheckPostRes() {}

type APIParsingPagePostBadRequest ErrorResponse

func (*APIParsingPagePostBadRequest) aPIParsingPagePostRes() {}

type APIParsingPagePostForbidden ErrorResponse

func (*APIParsingPagePostForbidden) aPIParsingPagePostRes() {}

type APIParsingPagePostInternalServerError ErrorResponse

func (*APIParsingPagePostInternalServerError) aPIParsingPagePostRes() {}

type APIParsingPagePostOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s APIParsingPagePostOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*APIParsingPagePostOK) aPIParsingPagePostRes() {}

type APIParsingPagePostReq struct {
	// Ссылка на книгу во внешней системе.
	BookURL url.URL `json:"book_url"`
	// Ссылка на изображение во внешней системе.
	ImageURL url.URL `json:"image_url"`
}

// GetBookURL returns the value of BookURL.
func (s *APIParsingPagePostReq) GetBookURL() url.URL {
	return s.BookURL
}

// GetImageURL returns the value of ImageURL.
func (s *APIParsingPagePostReq) GetImageURL() url.URL {
	return s.ImageURL
}

// SetBookURL sets the value of BookURL.
func (s *APIParsingPagePostReq) SetBookURL(val url.URL) {
	s.BookURL = val
}

// SetImageURL sets the value of ImageURL.
func (s *APIParsingPagePostReq) SetImageURL(val url.URL) {
	s.ImageURL = val
}

type APIParsingPagePostUnauthorized ErrorResponse

func (*APIParsingPagePostUnauthorized) aPIParsingPagePostRes() {}

// Данные книги.
// Ref: #/components/schemas/BookDetails
type BookDetails struct {
	// Ссылка на внешнюю систему.
	URL url.URL `json:"url"`
	// Название книги.
	Name string `json:"name"`
	// Количество страниц.
	PageCount int `json:"page_count"`
	// Данные атрибутов книги.
	Attributes []BookDetailsAttributesItem `json:"attributes"`
	// Данные страниц книги.
	Pages []BookDetailsPagesItem `json:"pages"`
}

// GetURL returns the value of URL.
func (s *BookDetails) GetURL() url.URL {
	return s.URL
}

// GetName returns the value of Name.
func (s *BookDetails) GetName() string {
	return s.Name
}

// GetPageCount returns the value of PageCount.
func (s *BookDetails) GetPageCount() int {
	return s.PageCount
}

// GetAttributes returns the value of Attributes.
func (s *BookDetails) GetAttributes() []BookDetailsAttributesItem {
	return s.Attributes
}

// GetPages returns the value of Pages.
func (s *BookDetails) GetPages() []BookDetailsPagesItem {
	return s.Pages
}

// SetURL sets the value of URL.
func (s *BookDetails) SetURL(val url.URL) {
	s.URL = val
}

// SetName sets the value of Name.
func (s *BookDetails) SetName(val string) {
	s.Name = val
}

// SetPageCount sets the value of PageCount.
func (s *BookDetails) SetPageCount(val int) {
	s.PageCount = val
}

// SetAttributes sets the value of Attributes.
func (s *BookDetails) SetAttributes(val []BookDetailsAttributesItem) {
	s.Attributes = val
}

// SetPages sets the value of Pages.
func (s *BookDetails) SetPages(val []BookDetailsPagesItem) {
	s.Pages = val
}

func (*BookDetails) aPIParsingBookPostRes() {}

type BookDetailsAttributesItem struct {
	// Код атрибута.
	Code BookDetailsAttributesItemCode `json:"code"`
	// Значения атрибута.
	Values []string `json:"values"`
}

// GetCode returns the value of Code.
func (s *BookDetailsAttributesItem) GetCode() BookDetailsAttributesItemCode {
	return s.Code
}

// GetValues returns the value of Values.
func (s *BookDetailsAttributesItem) GetValues() []string {
	return s.Values
}

// SetCode sets the value of Code.
func (s *BookDetailsAttributesItem) SetCode(val BookDetailsAttributesItemCode) {
	s.Code = val
}

// SetValues sets the value of Values.
func (s *BookDetailsAttributesItem) SetValues(val []string) {
	s.Values = val
}

// Код атрибута.
type BookDetailsAttributesItemCode string

const (
	BookDetailsAttributesItemCodeAuthor    BookDetailsAttributesItemCode = "author"
	BookDetailsAttributesItemCodeCategory  BookDetailsAttributesItemCode = "category"
	BookDetailsAttributesItemCodeCharacter BookDetailsAttributesItemCode = "character"
	BookDetailsAttributesItemCodeGroup     BookDetailsAttributesItemCode = "group"
	BookDetailsAttributesItemCodeLanguage  BookDetailsAttributesItemCode = "language"
	BookDetailsAttributesItemCodeParody    BookDetailsAttributesItemCode = "parody"
	BookDetailsAttributesItemCodeTag       BookDetailsAttributesItemCode = "tag"
)

// AllValues returns all BookDetailsAttributesItemCode values.
func (BookDetailsAttributesItemCode) AllValues() []BookDetailsAttributesItemCode {
	return []BookDetailsAttributesItemCode{
		BookDetailsAttributesItemCodeAuthor,
		BookDetailsAttributesItemCodeCategory,
		BookDetailsAttributesItemCodeCharacter,
		BookDetailsAttributesItemCodeGroup,
		BookDetailsAttributesItemCodeLanguage,
		BookDetailsAttributesItemCodeParody,
		BookDetailsAttributesItemCodeTag,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s BookDetailsAttributesItemCode) MarshalText() ([]byte, error) {
	switch s {
	case BookDetailsAttributesItemCodeAuthor:
		return []byte(s), nil
	case BookDetailsAttributesItemCodeCategory:
		return []byte(s), nil
	case BookDetailsAttributesItemCodeCharacter:
		return []byte(s), nil
	case BookDetailsAttributesItemCodeGroup:
		return []byte(s), nil
	case BookDetailsAttributesItemCodeLanguage:
		return []byte(s), nil
	case BookDetailsAttributesItemCodeParody:
		return []byte(s), nil
	case BookDetailsAttributesItemCodeTag:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *BookDetailsAttributesItemCode) UnmarshalText(data []byte) error {
	switch BookDetailsAttributesItemCode(data) {
	case BookDetailsAttributesItemCodeAuthor:
		*s = BookDetailsAttributesItemCodeAuthor
		return nil
	case BookDetailsAttributesItemCodeCategory:
		*s = BookDetailsAttributesItemCodeCategory
		return nil
	case BookDetailsAttributesItemCodeCharacter:
		*s = BookDetailsAttributesItemCodeCharacter
		return nil
	case BookDetailsAttributesItemCodeGroup:
		*s = BookDetailsAttributesItemCodeGroup
		return nil
	case BookDetailsAttributesItemCodeLanguage:
		*s = BookDetailsAttributesItemCodeLanguage
		return nil
	case BookDetailsAttributesItemCodeParody:
		*s = BookDetailsAttributesItemCodeParody
		return nil
	case BookDetailsAttributesItemCodeTag:
		*s = BookDetailsAttributesItemCodeTag
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type BookDetailsPagesItem struct {
	// Номер страницы в книге.
	PageNumber int `json:"page_number"`
	// Ссылка на изображение во внешней системе.
	URL url.URL `json:"url"`
	// Название файла с расширением.
	Filename string `json:"filename"`
}

// GetPageNumber returns the value of PageNumber.
func (s *BookDetailsPagesItem) GetPageNumber() int {
	return s.PageNumber
}

// GetURL returns the value of URL.
func (s *BookDetailsPagesItem) GetURL() url.URL {
	return s.URL
}

// GetFilename returns the value of Filename.
func (s *BookDetailsPagesItem) GetFilename() string {
	return s.Filename
}

// SetPageNumber sets the value of PageNumber.
func (s *BookDetailsPagesItem) SetPageNumber(val int) {
	s.PageNumber = val
}

// SetURL sets the value of URL.
func (s *BookDetailsPagesItem) SetURL(val url.URL) {
	s.URL = val
}

// SetFilename sets the value of Filename.
func (s *BookDetailsPagesItem) SetFilename(val string) {
	s.Filename = val
}

// Данные ошибки.
// Ref: #/components/schemas/ErrorResponse
type ErrorResponse struct {
	// Внутренний код ошибки.
	InnerCode string `json:"inner_code"`
	// Детальные данные ошибки.
	Details OptString `json:"details"`
}

// GetInnerCode returns the value of InnerCode.
func (s *ErrorResponse) GetInnerCode() string {
	return s.InnerCode
}

// GetDetails returns the value of Details.
func (s *ErrorResponse) GetDetails() OptString {
	return s.Details
}

// SetInnerCode sets the value of InnerCode.
func (s *ErrorResponse) SetInnerCode(val string) {
	s.InnerCode = val
}

// SetDetails sets the value of Details.
func (s *ErrorResponse) SetDetails(val OptString) {
	s.Details = val
}

type HeaderAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *HeaderAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *HeaderAuth) SetAPIKey(val string) {
	s.APIKey = val
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}
