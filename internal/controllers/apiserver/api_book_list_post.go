package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookListPost(ctx context.Context, req *serverAPI.APIBookListPostReq) (serverAPI.APIBookListPostRes, error) {
	filter := entities.BookFilter{}
	filter.FillNewest(req.Page.Value, req.Count)

	books, pageListRaw, err := c.webAPIUseCases.BookList(ctx, filter)
	if err != nil {
		return &serverAPI.APIBookListPostInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIBookListPostOK{
		Books: pkg.Map(books, func(b entities.BookToWeb) serverAPI.BookShortInfo {

			previewURL := serverAPI.OptURI{}

			if b.PreviewPage.Downloaded {
				previewURL = serverAPI.NewOptURI(c.getFileURL(
					b.PreviewPage.FileID,
					b.PreviewPage.Ext,
				))
			}

			return serverAPI.BookShortInfo{
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
		Pages: pkg.Map(pageListRaw, func(v int) serverAPI.APIBookListPostOKPagesItem {
			return serverAPI.APIBookListPostOKPagesItem{
				Value:       v,
				IsCurrent:   v == req.Page.Value,
				IsSeparator: v == -1,
			}
		}),
	}, nil
}
