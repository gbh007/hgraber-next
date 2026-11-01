package agentparseruc

import (
	"context"
	"io"
	"net/url"
)

func (uc *UseCase) Image(ctx context.Context, bookURL, imageURL url.URL) (io.ReadCloser, error) {
	return uc.loader.Image(ctx, imageURL, bookURL)
}
