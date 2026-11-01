package apiservercore

import (
	"context"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type ExternalAddrCtxKey struct{}

type ExternalAddr struct {
	Scheme string
	Host   string
}

func (c *Controller) GetFileURL(ctx context.Context, fileID uuid.UUID, ext string, fsID uuid.UUID) url.URL {
	if c.fsUseCases != nil {
		// FIXME: подумать над местом получше,
		// или более явным пробросом контекста,
		// или автообновлением токенов, чтобы не было надобности в ошибках.
		ctx, cancel := context.WithTimeout(ctx, time.Minute)
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

	addr, ok := ctx.Value(ExternalAddrCtxKey{}).(ExternalAddr)
	if !c.useHeaderExternalAddr || !ok {
		addr = ExternalAddr{
			Scheme: c.externalServerScheme,
			Host:   c.externalServerHostWithPort,
		}
	}

	u := url.URL{
		Scheme: addr.Scheme,
		Host:   addr.Host,
		Path:   "/api/file/" + fileID.String() + ext,
	}

	v := url.Values{}
	v.Add("fsid", fsID.String())
	u.RawQuery = v.Encode()

	return u
}

func (c *Controller) GetHProxyFileURL(ctx context.Context, bookURL, imageURL url.URL) url.URL {
	addr, ok := ctx.Value(ExternalAddrCtxKey{}).(ExternalAddr)
	if !c.useHeaderExternalAddr || !ok {
		addr = ExternalAddr{
			Scheme: c.externalServerScheme,
			Host:   c.externalServerHostWithPort,
		}
	}

	u := url.URL{
		Scheme: addr.Scheme,
		Host:   addr.Host,
		Path:   "/api/hproxy/file",
	}

	v := url.Values{}
	v.Add("book_url", bookURL.String())
	v.Add("image_url", imageURL.String())
	u.RawQuery = v.Encode()

	return u
}
