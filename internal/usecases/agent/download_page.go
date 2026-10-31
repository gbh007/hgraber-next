package agent

import (
	"context"
	"io"
	"net/url"
)

func (uc *UseCase) DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error) {
	return uc.loader.LoadImage(ctx, imageURL, bookURL) // FIXME: пока решение по быстрому и не учитывает что под капотом не буффер
}
