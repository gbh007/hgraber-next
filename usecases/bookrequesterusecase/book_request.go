package bookrequesterusecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type bookRequest struct {
	ID                      uuid.UUID
	IncludeAttributes       bool
	IncludeOriginAttributes bool
	IncludePages            bool
	IncludeLabels           bool
}

// FIXME: избавится от этого непотребства, вынести в отдельные методы получение списка книг и получение детальных данных по 1 книге.
func (uc *UseCase) requestBook(ctx context.Context, req bookRequest) (core.BookContainer, error) {
	b, err := uc.storage.GetBook(ctx, req.ID)
	if err != nil {
		return core.BookContainer{}, fmt.Errorf("get book: %w", err)
	}

	out := core.BookContainer{
		Book: b,
	}

	switch {
	case req.IncludeOriginAttributes:
		attributes, err := uc.storage.BookOriginAttributes(ctx, req.ID)
		if err != nil {
			return core.BookContainer{}, fmt.Errorf("get attributes: %w", err)
		}

		out.Attributes = attributes

	case req.IncludeAttributes:
		attributes, err := uc.storage.BookAttributes(ctx, req.ID)
		if err != nil {
			return core.BookContainer{}, fmt.Errorf("get attributes: %w", err)
		}

		out.Attributes = attributes
	}

	if req.IncludePages {
		pages, err := uc.storage.BookPages(ctx, req.ID)
		if err != nil {
			return core.BookContainer{}, fmt.Errorf("get pages: %w", err)
		}

		out.Pages = pages
	}

	if req.IncludeLabels {
		labels, err := uc.storage.Labels(ctx, req.ID)
		if err != nil {
			return core.BookContainer{}, fmt.Errorf("get labels: %w", err)
		}

		out.Labels = labels
	}

	return out, nil
}
