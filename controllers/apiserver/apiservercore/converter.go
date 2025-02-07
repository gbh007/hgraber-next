package apiservercore

import (
	"context"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) GetFileURL(fileID uuid.UUID, ext string, fsID uuid.UUID) url.URL {
	if c.fsUseCases != nil {
		// FIXME: подумать над местом получше,
		// или более явным пробросом контекста,
		// или автообновлением токенов, чтобы не было надобности в ошибках.
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		u, ok, err := c.fsUseCases.HighwayFileURL(ctx, fileID, ext, fsID)
		if err != nil {
			c.logger.ErrorContext(
				ctx, "get highway file url",
				slog.Any("error", err),
			)
		}

		if ok {
			return u
		}
	}

	u := url.URL{
		Scheme: c.externalServerScheme,
		Host:   c.externalServerHostWithPort,
		Path:   "/api/file/" + fileID.String() + ext,
	}

	v := url.Values{}
	v.Add("fsid", fsID.String())
	u.RawQuery = v.Encode()

	return u
}

func (c *Controller) ConvertPreviewPageUrl(p bff.PreviewPage) serverapi.OptURI {
	previewURL := serverapi.OptURI{}

	if p.Downloaded {
		previewURL = serverapi.NewOptURI(c.GetFileURL(
			p.FileID,
			p.Ext,
			p.FSID,
		))
	}

	return previewURL
}

func (c *Controller) ConvertSimpleBook(book core.Book, previewPage bff.PreviewPage) serverapi.BookSimple {
	return serverapi.BookSimple{
		ID:         book.ID,
		CreatedAt:  book.CreateAt,
		OriginURL:  OptURL(book.OriginURL),
		Name:       book.Name,
		PageCount:  book.PageCount,
		PreviewURL: c.ConvertPreviewPageUrl(previewPage),
		Flags: serverapi.BookFlags{
			ParsedName: book.ParsedName(),
			ParsedPage: book.PageCount > 0, // FIXME: не самый надежный метод, мб стоит придумать что-то другое
			IsVerified: book.Verified,
			IsDeleted:  book.Deleted,
			IsRebuild:  book.IsRebuild,
		},
	}
}

func (c *Controller) ConvertPreviewPage(page bff.PreviewPage) serverapi.PageSimple {
	return serverapi.PageSimple{
		PageNumber:  page.PageNumber,
		PreviewURL:  c.ConvertPreviewPageUrl(page),
		HasDeadHash: ConvertStatusFlagToAPI(page.HasDeadHash),
	}
}

func ConvertBookAttribute(a bff.AttributeToWeb) serverapi.BookAttribute {
	return serverapi.BookAttribute{
		Code:   a.Code,
		Name:   a.Name,
		Values: a.Values,
	}
}

func ConvertBookFullToBookRaw(book core.BookContainer) *serverapi.BookRaw {
	return &serverapi.BookRaw{
		ID:        book.Book.ID,
		CreateAt:  book.Book.CreateAt,
		OriginURL: OptURL(book.Book.OriginURL),
		Name:      book.Book.Name,
		PageCount: book.Book.PageCount,
		Attributes: pkg.MapToSlice(book.Attributes, func(code string, values []string) serverapi.BookRawAttributesItem {
			return serverapi.BookRawAttributesItem{
				Code:   code,
				Values: values,
			}
		}),
		Pages: pkg.Map(book.Pages, func(p core.Page) serverapi.BookRawPagesItem {
			return serverapi.BookRawPagesItem{
				PageNumber: p.PageNumber,
				OriginURL:  OptURL(p.OriginURL),
				Ext:        p.Ext,
				CreateAt:   p.CreateAt,
				Downloaded: p.Downloaded,
				LoadAt:     OptTime(p.LoadAt),
			}
		}),
		Labels: pkg.Map(book.Labels, func(l core.BookLabel) serverapi.BookRawLabelsItem {
			return serverapi.BookRawLabelsItem{
				PageNumber: l.PageNumber,
				Name:       l.Name,
				Value:      l.Value,
				CreateAt:   l.CreateAt,
			}
		}),
	}
}

func ConvertBookRawToBookFull(book *serverapi.BookRaw) core.BookContainer {
	if book == nil {
		return core.BookContainer{}
	}

	return core.BookContainer{
		Book: core.Book{
			ID:        book.ID,
			Name:      book.Name,
			OriginURL: UrlFromOpt(book.OriginURL),
			PageCount: book.PageCount,
			CreateAt:  book.CreateAt,
			// FIXME: нет ряд полей, возможно стоит расширить api
		},
		Pages: pkg.Map(book.Pages, func(raw serverapi.BookRawPagesItem) core.Page {
			return core.Page{
				BookID:     book.ID,
				PageNumber: raw.PageNumber,
				Ext:        raw.Ext,
				OriginURL:  UrlFromOpt(raw.OriginURL),
				CreateAt:   raw.CreateAt,
				Downloaded: raw.Downloaded,
				LoadAt:     raw.LoadAt.Value,
				// FIXME: нет ряд полей, возможно стоит расширить api
			}
		}),
		Attributes: pkg.SliceToMap(book.Attributes, func(raw serverapi.BookRawAttributesItem) (string, []string) {
			return raw.Code, raw.Values
		}),
		Labels: pkg.Map(book.Labels, func(raw serverapi.BookRawLabelsItem) core.BookLabel {
			return core.BookLabel{
				BookID:     book.ID,
				PageNumber: raw.PageNumber,
				Name:       raw.Name,
				Value:      raw.Value,
				CreateAt:   raw.CreateAt,
			}
		}),
	}
}

func ConvertAgentToAPI(raw core.Agent) serverapi.Agent {
	return serverapi.Agent{
		ID:            raw.ID,
		Name:          raw.Name,
		Addr:          raw.Addr,
		Token:         raw.Token,
		CanParse:      raw.CanParse,
		CanParseMulti: raw.CanParseMulti,
		CanExport:     raw.CanExport,
		HasFs:         raw.HasFS,
		Priority:      raw.Priority,
		CreatedAt:     raw.CreateAt,
	}
}

func ConvertFileSystemInfoFromAPI(raw *serverapi.FileSystemInfo) fsmodel.FileStorageSystem {
	return fsmodel.FileStorageSystem{
		ID:                  raw.ID,
		Name:                raw.Name,
		Description:         raw.Description.Value,
		AgentID:             raw.AgentID.Value,
		Path:                raw.Path.Value,
		DownloadPriority:    raw.DownloadPriority,
		DeduplicatePriority: raw.DeduplicatePriority,
		HighwayEnabled:      raw.HighwayEnabled,
		HighwayAddr:         UrlFromOpt(raw.HighwayAddr),
		CreatedAt:           raw.CreatedAt,
	}
}

func ConvertFileSystemInfoToAPI(raw fsmodel.FileStorageSystem) serverapi.FileSystemInfo {
	return serverapi.FileSystemInfo{
		ID:                  raw.ID,
		Name:                raw.Name,
		Description:         OptString(raw.Description),
		AgentID:             OptUUID(raw.AgentID),
		Path:                OptString(raw.Path),
		DownloadPriority:    raw.DownloadPriority,
		DeduplicatePriority: raw.DeduplicatePriority,
		HighwayEnabled:      raw.HighwayEnabled,
		HighwayAddr:         OptURL(raw.HighwayAddr),
		CreatedAt:           raw.CreatedAt,
	}
}

func ConvertStatusFlagToAPI(f bff.StatusFlag) serverapi.OptBool {
	return serverapi.OptBool{
		Value: f == bff.TrueStatusFlag,
		Set:   f != bff.UnknownStatusFlag,
	}
}

func ConvertFSDBFilesInfoToAPI(raw *fsmodel.FSFilesInfo) serverapi.OptFSDBFilesInfo {
	if raw == nil {
		return serverapi.OptFSDBFilesInfo{}
	}

	return serverapi.NewOptFSDBFilesInfo(serverapi.FSDBFilesInfo{
		Count:         raw.Count,
		Size:          raw.Size,
		SizeFormatted: core.PrettySize(raw.Size),
	})
}
