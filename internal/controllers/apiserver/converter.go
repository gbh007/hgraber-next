package apiserver

import (
	"context"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
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

func optUUID(u uuid.UUID) serverAPI.OptUUID {
	if u == uuid.Nil {
		return serverAPI.OptUUID{}
	}

	return serverAPI.NewOptUUID(u)
}

func (c *Controller) getFileURL(fileID uuid.UUID, ext string, fsID uuid.UUID) url.URL {
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

func (c *Controller) getPagePreview(p entities.BFFPreviewPage) serverAPI.OptURI {
	previewURL := serverAPI.OptURI{}

	if p.Downloaded {
		previewURL = serverAPI.NewOptURI(c.getFileURL(
			p.FileID,
			p.Ext,
			p.FSID,
		))
	}

	return previewURL
}

func (c *Controller) convertSimpleBook(book entities.Book, previewPage entities.BFFPreviewPage) serverAPI.BookSimple {
	return serverAPI.BookSimple{
		ID:         book.ID,
		CreatedAt:  book.CreateAt,
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

func (c *Controller) convertPreviewPage(page entities.BFFPreviewPage) serverAPI.PageSimple {
	return serverAPI.PageSimple{
		PageNumber:  page.PageNumber,
		PreviewURL:  c.getPagePreview(page),
		HasDeadHash: convertStatusFlagToAPI(page.HasDeadHash),
	}
}

func convertBookAttribute(a entities.AttributeToWeb) serverAPI.BookAttribute {
	return serverAPI.BookAttribute{
		Code:   a.Code,
		Name:   a.Name,
		Values: a.Values,
	}
}

func convertBookFullToBookRaw(book entities.BookContainer) *serverAPI.BookRaw {
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

func convertBookRawToBookFull(book *serverAPI.BookRaw) entities.BookContainer {
	if book == nil {
		return entities.BookContainer{}
	}

	return entities.BookContainer{
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

func convertAgentToAPI(raw entities.Agent) serverAPI.Agent {
	return serverAPI.Agent{
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

func convertFileSystemInfoFromAPI(raw *serverAPI.FileSystemInfo) entities.FileStorageSystem {
	return entities.FileStorageSystem{
		ID:                  raw.ID,
		Name:                raw.Name,
		Description:         raw.Description.Value,
		AgentID:             raw.AgentID.Value,
		Path:                raw.Path.Value,
		DownloadPriority:    raw.DownloadPriority,
		DeduplicatePriority: raw.DeduplicatePriority,
		HighwayEnabled:      raw.HighwayEnabled,
		HighwayAddr:         urlFromOpt(raw.HighwayAddr),
		CreatedAt:           raw.CreatedAt,
	}
}

func convertFileSystemInfoToAPI(raw entities.FileStorageSystem) serverAPI.FileSystemInfo {
	return serverAPI.FileSystemInfo{
		ID:                  raw.ID,
		Name:                raw.Name,
		Description:         optString(raw.Description),
		AgentID:             optUUID(raw.AgentID),
		Path:                optString(raw.Path),
		DownloadPriority:    raw.DownloadPriority,
		DeduplicatePriority: raw.DeduplicatePriority,
		HighwayEnabled:      raw.HighwayEnabled,
		HighwayAddr:         optURL(raw.HighwayAddr),
		CreatedAt:           raw.CreatedAt,
	}
}

func convertStatusFlagToAPI(f entities.StatusFlag) serverAPI.OptBool {
	return serverAPI.OptBool{
		Value: f == entities.TrueStatusFlag,
		Set:   f != entities.UnknownStatusFlag,
	}
}

func convertFSDBFilesInfoToAPI(raw *entities.FSFilesInfo) serverAPI.OptFSDBFilesInfo {
	if raw == nil {
		return serverAPI.OptFSDBFilesInfo{}
	}

	return serverAPI.NewOptFSDBFilesInfo(serverAPI.FSDBFilesInfo{
		Count:         raw.Count,
		Size:          raw.Size,
		SizeFormatted: entities.PrettySize(raw.Size),
	})
}
