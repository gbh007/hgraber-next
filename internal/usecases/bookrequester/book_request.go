package bookrequester

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type bookRequest struct {
	ID                      uuid.UUID
	IncludeAttributes       bool
	IncludeOriginAttributes bool
	IncludePages            bool
	IncludeLabels           bool
}

func (uc *UseCase) requestBook(ctx context.Context, req bookRequest) (entities.BookFull, error) {
	b, err := uc.storage.GetBook(ctx, req.ID)
	if err != nil {
		return entities.BookFull{}, fmt.Errorf("get book :%w", err)
	}

	out := entities.BookFull{
		Book: b,
	}

	switch {
	case req.IncludeOriginAttributes:
		attributes, err := uc.storage.BookOriginAttributes(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get attributes :%w", err)
		}

		out.Attributes = attributes

	case req.IncludeAttributes:
		attributes, err := uc.storage.BookAttributes(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get attributes :%w", err)
		}

		out.Attributes = attributes
	}

	if req.IncludePages {
		pages, err := uc.storage.BookPages(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get pages :%w", err)
		}

		out.Pages = pages
	}

	if req.IncludeLabels {
		labels, err := uc.storage.Labels(ctx, req.ID)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get labels :%w", err)
		}

		out.Labels = labels
	}

	return out, nil
}
