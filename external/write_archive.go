package external

import (
	"archive/zip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrSkipPageBody = errors.New("skip page body")

func WriteArchive(
	ctx context.Context,
	zipWriter *zip.Writer,
	pageBodyGetter func(ctx context.Context, pageNumber int) (io.Reader, error),
	info Info,
) error {
	w, err := zipWriter.Create("info.json")
	if err != nil {
		return fmt.Errorf("create info: %w", err)
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	err = enc.Encode(info)
	if err != nil {
		return fmt.Errorf("encode info: %w", err)
	}

	for _, p := range info.Data.Pages {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		pageBody, err := pageBodyGetter(ctx, p.PageNumber)

		// Нет файла, пропускаем
		if errors.Is(err, ErrSkipPageBody) {
			continue
		}

		if err != nil {
			return fmt.Errorf("get page body: %w", err)
		}

		filename := fmt.Sprintf("%d%s", p.PageNumber, p.Ext)

		if !strings.HasPrefix(p.Ext, ".") {
			filename = fmt.Sprintf("%d.%s", p.PageNumber, p.Ext)
		}

		w, err := zipWriter.Create(filename)
		if err != nil {
			return fmt.Errorf("create page %d body: %w", p.PageNumber, err)
		}

		_, err = io.Copy(w, pageBody)
		if err != nil {
			return fmt.Errorf("copy page %d body: %w", p.PageNumber, err)
		}
	}

	return nil
}
