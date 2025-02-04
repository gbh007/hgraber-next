package rebuilder

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) rebuildBookSave(
	ctx context.Context,
	isNewBook bool,
	bookToMerge core.Book,
	newAttributes map[string][]string,
	newPages []core.Page,
	newLabels []core.BookLabel,
) (returnErr error) {
	var err error

	if isNewBook {
		err = uc.storage.NewBook(ctx, bookToMerge)
		if err != nil {
			return fmt.Errorf("storage: create book: %w", err)
		}

		defer func() { // Удаляем книгу, если не получилось ее создать полностью
			if returnErr != nil {
				deleteErr := uc.storage.DeleteBook(ctx, bookToMerge.ID)
				if deleteErr != nil {
					uc.logger.ErrorContext(
						ctx, "delete new book after unsuccess rebuild",
						slog.Any("error", deleteErr),
						slog.String("book_id", bookToMerge.ID.String()),
					)
				}
			}
		}()
	} else {
		err = uc.storage.UpdateBook(ctx, bookToMerge)
		if err != nil {
			return fmt.Errorf("storage: create book: %w", err)
		}
	}

	if len(newPages) > 0 {
		err = uc.storage.NewBookPages(ctx, newPages)
		if err != nil {
			return fmt.Errorf("storage: new pages: %w", err)
		}
	}

	if len(newAttributes) > 0 {
		err = uc.storage.UpdateOriginAttributes(ctx, bookToMerge.ID, newAttributes)
		if err != nil {
			return fmt.Errorf("storage: set origin attributes: %w", err)
		}

		err = uc.storage.UpdateAttributes(ctx, bookToMerge.ID, newAttributes)
		if err != nil {
			return fmt.Errorf("storage: set attributes: %w", err)
		}
	}

	if len(newLabels) > 0 {
		err = uc.storage.SetLabels(ctx, newLabels)
		if err != nil {
			return fmt.Errorf("storage: set labels: %w", err)
		}
	}

	return nil
}
