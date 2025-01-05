package apiserver

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func optURL(u *url.URL) serverAPI.OptURI {
	if u == nil {
		return serverAPI.OptURI{}
	}

	return serverAPI.NewOptURI(*u)
}

func urlFromOpt(u serverAPI.OptURI) *url.URL {
	if !u.Set {
		return nil
	}

	return &u.Value
}

func optTime(t time.Time) serverAPI.OptDateTime {
	if t.IsZero() {
		return serverAPI.OptDateTime{}
	}

	return serverAPI.NewOptDateTime(t)
}

func optString(s string) serverAPI.OptString {
	if s == "" {
		return serverAPI.OptString{}
	}

	return serverAPI.NewOptString(s)
}

func (c *Controller) getFileURL(fileID uuid.UUID, ext string) url.URL {
	return url.URL{
		Scheme: c.externalServerScheme,
		Host:   c.externalServerHostWithPort,
		Path:   "/api/file/" + fileID.String() + ext,
	}
}

func (c *Controller) getPagePreview(p entities.Page) serverAPI.OptURI {
	previewURL := serverAPI.OptURI{}

	if p.Downloaded {
		previewURL = serverAPI.NewOptURI(c.getFileURL(
			p.FileID,
			p.Ext,
		))
	}

	return previewURL
}

func (c *Controller) convertSimpleBook(book entities.Book, previewPage entities.Page) serverAPI.BookSimple {
	return serverAPI.BookSimple{
		ID:         book.ID,
		CreateAt:   book.CreateAt,
		OriginURL:  optURL(book.OriginURL),
		Name:       book.Name,
		PageCount:  book.PageCount,
		PreviewURL: c.getPagePreview(previewPage),
		Flags: serverAPI.BookFlags{
			ParsedName: book.ParsedName(),
			ParsedPage: book.PageCount > 0, // FIXME: не самый надежный метод, мб стоит придумать что-то другое
			IsVerified: book.Verified,
			IsDeleted:  book.Deleted,
			IsRebuild:  book.IsRebuild,
		},
	}
}

func (c *Controller) convertSimplePageWithDeadHash(page entities.PageWithDeadHash) serverAPI.PageSimple {
	return serverAPI.PageSimple{
		PageNumber:  page.PageNumber,
		PreviewURL:  c.getPagePreview(page.Page),
		HasDeadHash: serverAPI.NewOptBool(page.HasDeadHash),
	}
}

func convertBookAttribute(a entities.AttributeToWeb) serverAPI.BookAttribute {
	return serverAPI.BookAttribute{
		Code:   a.Code,
		Name:   a.Name,
		Values: a.Values,
	}
}

func convertBookFullToBookRaw(book entities.BookFull) *serverAPI.BookRaw {
	return &serverAPI.BookRaw{
		ID:        book.Book.ID,
		CreateAt:  book.Book.CreateAt,
		OriginURL: optURL(book.Book.OriginURL),
		Name:      book.Book.Name,
		PageCount: book.Book.PageCount,
		Attributes: pkg.MapToSlice(book.Attributes, func(code string, values []string) serverAPI.BookRawAttributesItem {
			return serverAPI.BookRawAttributesItem{
				Code:   code,
				Values: values,
			}
		}),
		Pages: pkg.Map(book.Pages, func(p entities.Page) serverAPI.BookRawPagesItem {
			return serverAPI.BookRawPagesItem{
				PageNumber: p.PageNumber,
				OriginURL:  optURL(p.OriginURL),
				Ext:        p.Ext,
				CreateAt:   p.CreateAt,
				Downloaded: p.Downloaded,
				LoadAt:     optTime(p.LoadAt),
			}
		}),
		Labels: pkg.Map(book.Labels, func(l entities.BookLabel) serverAPI.BookRawLabelsItem {
			return serverAPI.BookRawLabelsItem{
				PageNumber: l.PageNumber,
				Name:       l.Name,
				Value:      l.Value,
				CreateAt:   l.CreateAt,
			}
		}),
	}
}

func convertBookRawToBookFull(book *serverAPI.BookRaw) entities.BookFull {
	if book == nil {
		return entities.BookFull{}
	}

	return entities.BookFull{
		Book: entities.Book{
			ID:        book.ID,
			Name:      book.Name,
			OriginURL: urlFromOpt(book.OriginURL),
			PageCount: book.PageCount,
			CreateAt:  book.CreateAt,
			// FIXME: нет ряд полей, возможно стоит расширить api
		},
		Pages: pkg.Map(book.Pages, func(raw serverAPI.BookRawPagesItem) entities.Page {
			return entities.Page{
				BookID:     book.ID,
				PageNumber: raw.PageNumber,
				Ext:        raw.Ext,
				OriginURL:  urlFromOpt(raw.OriginURL),
				CreateAt:   raw.CreateAt,
				Downloaded: raw.Downloaded,
				LoadAt:     raw.LoadAt.Value,
				// FIXME: нет ряд полей, возможно стоит расширить api
			}
		}),
		Attributes: pkg.SliceToMap(book.Attributes, func(raw serverAPI.BookRawAttributesItem) (string, []string) {
			return raw.Code, raw.Values
		}),
		Labels: pkg.Map(book.Labels, func(raw serverAPI.BookRawLabelsItem) entities.BookLabel {
			return entities.BookLabel{
				BookID:     book.ID,
				PageNumber: raw.PageNumber,
				Name:       raw.Name,
				Value:      raw.Value,
				CreateAt:   raw.CreateAt,
			}
		}),
	}
}
