package external

import (
	"archive/zip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var ErrBookInfoNotFound = errors.New("book info not found")

type ReadArchiveOptions struct {
	HandlePageBody func(ctx context.Context, pageNumber int, filename string, body io.Reader) error
	ReadInfoTXT    bool
	ReadDataJSON   bool
	SkipInfo       bool
}

//nolint:gocognit,cyclop,funlen // будет исправлено позднее
func ReadArchive(
	ctx context.Context,
	zipReader *zip.Reader,
	opt ReadArchiveOptions,
) (Info, error) {
	var (
		infoJSON, dataJSON, infoTXT                Info
		foundInfoJSON, foundDataJSON, foundInfoTXT bool
	)

	for _, f := range zipReader.File {
		select {
		case <-ctx.Done():
			return Info{}, ctx.Err() //nolint:wrapcheck // оставляем оригинальные данные
		default:
		}

		switch f.Name {
		case "info.json":
			if opt.SkipInfo {
				continue
			}

			foundInfoJSON = true

			r, err := f.Open()
			if err != nil {
				return Info{}, fmt.Errorf("open info.json file: %w", err)
			}

			err = json.NewDecoder(r).Decode(&infoJSON)
			if err != nil {
				return Info{}, fmt.Errorf("decode info.json file: %w", err)
			}

			err = r.Close()
			if err != nil {
				return Info{}, fmt.Errorf("close info.json file: %w", err)
			}

		case "data.json":
			if opt.SkipInfo || !opt.ReadDataJSON {
				continue
			}

			r, err := f.Open()
			if err != nil {
				return Info{}, fmt.Errorf("open data.json file: %w", err)
			}

			dataJSON, foundDataJSON, err = HG4ParseDataJSON(r)
			if err != nil {
				return Info{}, fmt.Errorf("decode data.json file: %w", err)
			}

			err = r.Close()
			if err != nil {
				return Info{}, fmt.Errorf("close data.json file: %w", err)
			}

		case "info.txt":
			if opt.SkipInfo || !opt.ReadInfoTXT {
				continue
			}

			r, err := f.Open()
			if err != nil {
				return Info{}, fmt.Errorf("open info.txt file: %w", err)
			}

			infoTXT, foundInfoTXT, err = HG4ParseInfoTXT(r)
			if err != nil {
				return Info{}, fmt.Errorf("decode info.txt file: %w", err)
			}

			err = r.Close()
			if err != nil {
				return Info{}, fmt.Errorf("close info.txt file: %w", err)
			}

		default:
			if opt.HandlePageBody == nil {
				continue
			}

			number, _ := strconv.Atoi(strings.Split(f.Name, ".")[0])
			if number < 1 {
				continue
			}

			r, err := f.Open()
			if err != nil {
				return Info{}, fmt.Errorf("open page (%d) file: %w", number, err)
			}

			err = opt.HandlePageBody(ctx, number, f.Name, r)
			if err != nil {
				return Info{}, fmt.Errorf("page (%d) handle body: %w", number, err)
			}

			err = r.Close()
			if err != nil {
				return Info{}, fmt.Errorf("close page (%d) file: %w", number, err)
			}
		}
	}

	switch {
	case foundInfoJSON:
		return infoJSON, nil

	case foundDataJSON:
		return dataJSON, nil

	case foundInfoTXT:
		return infoTXT, nil

	case opt.SkipInfo:
		return Info{}, nil

	default:
		return Info{}, ErrBookInfoNotFound
	}
}
