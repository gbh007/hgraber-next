package apiserver

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func optURL(u *url.URL) serverAPI.OptURI {
	if u == nil {
		return serverAPI.OptURI{}
	}

	return serverAPI.NewOptURI(*u)
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
	}
}

func (c *Controller) convertSimplePage(page entities.Page) serverAPI.PageSimple {
	return serverAPI.PageSimple{
		PageNumber: page.PageNumber,
		PreviewURL: c.getPagePreview(page),
	}
}
