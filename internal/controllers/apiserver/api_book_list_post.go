package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIBookListPost(ctx context.Context, req *serverAPI.BookFilter) (serverAPI.APIBookListPostRes, error) {
	filter := convertAPIBookFilter(*req)

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

func convertAPIBookFilter(req serverAPI.BookFilter) entities.BookFilter {
	filter := entities.BookFilter{}

	if req.Count.IsSet() {
		filter.FillNewest(req.Page.Value, req.Count.Value)
	}

	if req.DeleteStatus.IsSet() {
		switch req.DeleteStatus.Value {
		case serverAPI.BookFilterDeleteStatusAll:
		case serverAPI.BookFilterDeleteStatusOnly:
			filter.ShowDeleted = entities.BookFilterShowTypeOnly
		case serverAPI.BookFilterDeleteStatusExcept:
			filter.ShowDeleted = entities.BookFilterShowTypeExcept
		}
	}

	if req.VerifyStatus.IsSet() {
		switch req.VerifyStatus.Value {
		case serverAPI.BookFilterVerifyStatusAll:
		case serverAPI.BookFilterVerifyStatusOnly:
			filter.ShowVerified = entities.BookFilterShowTypeOnly
		case serverAPI.BookFilterVerifyStatusExcept:
			filter.ShowVerified = entities.BookFilterShowTypeExcept
		}
	}

	filter.Fields.Name = req.Filter.Value.Name.Value
	filter.From = req.From.Value
	filter.To = req.To.Value

	return filter
}
