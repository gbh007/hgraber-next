package rebuilder

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) RestoreBook(ctx context.Context, bookID uuid.UUID, onlyPages bool) error {
	deletedPages, err := uc.storage.DeletedPages(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get deleted pages: %w", err)
	}

	md5Sums := make([]string, 0, len(deletedPages))

	for _, page := range deletedPages {
		md5Sums = append(md5Sums, page.Md5Sum)
	}

	fileIDs := make(map[entities.FileHash]uuid.UUID, len(md5Sums))

	if len(md5Sums) > 0 {
		files, err := uc.storage.FilesByMD5Sums(ctx, md5Sums)
		if err != nil {
			return fmt.Errorf("storage: get files: %w", err)
		}

		for _, file := range files {
			fileIDs[file.Hash()] = file.ID
		}
	}

	pageToRestore := make([]entities.Page, 0, min(len(deletedPages), len(fileIDs)))
	pageNumbersToRestore := make([]int, 0, min(len(deletedPages), len(fileIDs)))

	for _, page := range deletedPages {
		fileID, ok := fileIDs[page.FileHash]
		if !ok { // Восстанавливаем только страницы с файлами
			continue
		}

		page.FileID = fileID

		pageToRestore = append(pageToRestore, page.Page)
		pageNumbersToRestore = append(pageNumbersToRestore, page.PageNumber)
	}

	if len(pageToRestore) > 0 {
		err = uc.storage.NewBookPages(ctx, pageToRestore)
		if err != nil {
			return fmt.Errorf("storage: new pages: %w", err)
		}

		err = uc.storage.RemoveDeletedPages(ctx, bookID, pageNumbersToRestore)
		if err != nil {
			return fmt.Errorf("storage: remove restored pages from deleted: %w", err)
		}
	}

	if onlyPages {
		return nil
	}

	attributes, err := uc.storage.BookOriginAttributes(ctx, bookID)
	if err != nil {
		return fmt.Errorf("storage: get book origin attributes: %w", err)
	}

	if len(attributes) > 0 {
		err = uc.storage.UpdateAttributes(ctx, bookID, attributes)
		if err != nil {
			return fmt.Errorf("storage: update book attributes: %w", err)
		}
	}

	err = uc.storage.UpdateBookDeletion(ctx, entities.Book{
		ID:      bookID,
		Deleted: false,
	})
	if err != nil {
		return fmt.Errorf("storage: update book: %w", err)
	}

	return nil
}
