package bookusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) GetBookIDsForCalculation(ctx context.Context) ([]uuid.UUID, error) {
	return uc.storage.BookIDs(ctx, core.BookFilter{}) //nolint:wrapcheck // нет смысла в врапинге
}

//nolint:cyclop,funlen // будет исправлено позднее
func (uc *UseCase) CalculateBook(ctx context.Context, bookID uuid.UUID) error {
	book, err := uc.storage.GetBook(ctx, bookID)
	if err != nil {
		return fmt.Errorf("get book: %w", err)
	}

	bookPages, err := uc.storage.BookPagesWithHash(ctx, bookID)
	if err != nil {
		return fmt.Errorf("get pages: %w", err)
	}

	fileCounts := make(map[core.FileHash]struct{}, len(bookPages))

	md5Sums := make([]string, len(bookPages))
	for i, page := range bookPages {
		md5Sums[i] = page.Md5Sum
		fileCounts[page.FileHash] = struct{}{}
	}

	deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
	if err != nil {
		return fmt.Errorf("get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[core.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	var (
		calcPageCount     int64
		calcFileCount     int64
		calcDeadHashCount int64
		calcPageSize      int64
		calcFileSize      int64
		calcDeadHashSize  int64
		calcAvgPageSize   int64
	)

	for _, page := range bookPages {
		calcPageCount++
		calcPageSize += page.Size
	}

	for file := range fileCounts {
		calcFileCount++
		calcFileSize += file.Size

		if _, ok := existsDeadHashes[file]; ok {
			calcDeadHashCount++
			calcDeadHashSize += file.Size
		}
	}

	calcAvgPageSize = calcPageSize / calcPageCount

	needUpdate := false
	book.Calc.CalculatedAt = time.Now().UTC()

	if book.Calc.CalcPageCount == nil || *book.Calc.CalcPageCount != calcPageCount {
		book.Calc.CalcPageCount = &calcPageCount
		needUpdate = true
	}

	if book.Calc.CalcFileCount == nil || *book.Calc.CalcFileCount != calcFileCount {
		book.Calc.CalcFileCount = &calcFileCount
		needUpdate = true
	}

	if book.Calc.CalcDeadHashCount == nil || *book.Calc.CalcDeadHashCount != calcDeadHashCount {
		book.Calc.CalcDeadHashCount = &calcDeadHashCount
		needUpdate = true
	}

	if book.Calc.CalcPageSize == nil || *book.Calc.CalcPageSize != calcPageSize {
		book.Calc.CalcPageSize = &calcPageSize
		needUpdate = true
	}

	if book.Calc.CalcFileSize == nil || *book.Calc.CalcFileSize != calcFileSize {
		book.Calc.CalcFileSize = &calcFileSize
		needUpdate = true
	}

	if book.Calc.CalcDeadHashSize == nil || *book.Calc.CalcDeadHashSize != calcDeadHashSize {
		book.Calc.CalcDeadHashSize = &calcDeadHashSize
		needUpdate = true
	}

	if book.Calc.CalcAvgPageSize == nil || *book.Calc.CalcAvgPageSize != calcAvgPageSize {
		book.Calc.CalcAvgPageSize = &calcAvgPageSize
		needUpdate = true
	}

	if needUpdate {
		err = uc.storage.UpdateBookCalculation(ctx, bookID, book.Calc)
		if err != nil {
			return fmt.Errorf("update book calculation: %w", err)
		}
	}

	return nil
}
