package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APIBookListPost(ctx context.Context, req *server.APIBookListPostReq) (server.APIBookListPostRes, error) {
	filter := entities.BookFilter{}
	filter.FillNewest(req.Page.Value, req.Count)

	books, pageListRaw, err := c.webAPIUseCases.BookList(ctx, filter)
	if err != nil {
		return &server.APIBookListPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIBookListPostOK{
		Books: pkg.Map(books, func(b entities.BookToWeb) server.BookShortInfo {

			previewURL := server.OptURI{}

			if b.PreviewPage.Downloaded {
				previewURL = server.NewOptURI(c.getFileURL(
					b.PreviewPage.FileID,
					b.PreviewPage.Ext,
				))
			}

			return server.BookShortInfo{
				ID:                b.Book.ID,
				Created:           b.Book.CreateAt,
				PreviewURL:        previewURL,
				ParsedName:        b.ParsedName(),
				Name:              b.Book.Name,
				ParsedPage:        b.ParsedPages,
				PageCount:         b.Book.PageCount,
				PageLoadedPercent: b.PageDownloadPercent(),
				Tags:              b.Tags,
				HasMoreTags:       b.HasMoreTags,
			}
		}),
		Pages: pkg.Map(pageListRaw, func(v int) server.APIBookListPostOKPagesItem {
			return server.APIBookListPostOKPagesItem{
				Value:       v,
				IsCurrent:   v == req.Page.Value,
				IsSeparator: v == -1,
			}
		}),
	}, nil
}
