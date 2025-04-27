package apiservercore

import (
	"context"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"
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

func (c *Controller) GetHProxyFileURL(bookURL, imageURL url.URL) url.URL {
	u := url.URL{
		Scheme: c.externalServerScheme,
		Host:   c.externalServerHostWithPort,
		Path:   "/api/hproxy/file",
	}

	v := url.Values{}
	v.Add("book_url", bookURL.String())
	v.Add("image_url", imageURL.String())
	u.RawQuery = v.Encode()

	return u
}
