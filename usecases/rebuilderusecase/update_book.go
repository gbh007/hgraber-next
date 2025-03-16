package rebuilderusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) UpdateBook(ctx context.Context, book core.BookContainer) error {
	originBook, err := uc.storage.GetBook(ctx, book.Book.ID)
	if err != nil {
		return fmt.Errorf("storage: get book: %w", err)
	}

	// Переносим только часть данных
	originBook.Name = book.Book.Name
	originBook.OriginURL = book.Book.OriginURL
	originBook.AttributesParsed = true

	err = uc.storage.UpdateBook(ctx, originBook)
	if err != nil {
		return fmt.Errorf("storage: update book: %w", err)
	}

	if len(book.Attributes) > 0 {
		attributes := book.Attributes

		if uc.autoRemap {
			remaps, err := uc.storage.AttributeRemaps(ctx)
			if err != nil {
				return fmt.Errorf("storage: get attributes remaps: %w", err)
			}

			remaper := core.NewAttributeRemaper(remaps, uc.remapToLower)
			attributes = remaper.Remap(attributes)
		}

		err = uc.storage.UpdateOriginAttributes(ctx, book.Book.ID, book.Attributes)
		if err != nil {
			return fmt.Errorf("storage: update origin attributes: %w", err)
		}

		if len(attributes) > 0 {
			err = uc.storage.UpdateAttributes(ctx, book.Book.ID, attributes)
			if err != nil {
				return fmt.Errorf("storage: update attributes: %w", err)
			}
		} else {
			err = uc.storage.DeleteBookAttributes(ctx, book.Book.ID)
			if err != nil {
				return fmt.Errorf("storage: delete book origin attributes: %w", err)
			}
		}
	} else {
		err = uc.storage.DeleteBookAttributes(ctx, book.Book.ID)
		if err != nil {
			return fmt.Errorf("storage: delete book attributes: %w", err)
		}

		err = uc.storage.DeleteBookOriginAttributes(ctx, book.Book.ID)
		if err != nil {
			return fmt.Errorf("storage: delete book origin attributes: %w", err)
		}
	}

	if len(book.Labels) > 0 {
		err = uc.storage.ReplaceLabels(ctx, book.Book.ID, book.Labels)
		if err != nil {
			return fmt.Errorf("storage: replace labels: %w", err)
		}
	} else {
		err = uc.storage.DeleteBookLabels(ctx, book.Book.ID)
		if err != nil {
			return fmt.Errorf("storage: delete book labels: %w", err)
		}
	}

	return nil
}
