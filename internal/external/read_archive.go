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

func ReadArchive(
	ctx context.Context,
	zipReader *zip.Reader,
	handlePageBody func(ctx context.Context, pageNumber int, body io.Reader) error,
) (Info, error) {
	info := Info{}
	found := false

	for _, f := range zipReader.File {
		select {
		case <-ctx.Done():
			return Info{}, ctx.Err()
		default:
		}

		if f.Name == "info.json" {
			found = true

			r, err := f.Open()
			if err != nil {
				return Info{}, fmt.Errorf("open info file: %w", err)
			}

			err = json.NewDecoder(r).Decode(&info)
			if err != nil {
				return Info{}, fmt.Errorf("decode info file: %w", err)
			}

			err = r.Close()
			if err != nil {
				return Info{}, fmt.Errorf("close info file: %w", err)
			}

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

		err = handlePageBody(ctx, number, r)
		if err != nil {
			return Info{}, fmt.Errorf("page (%d) handle body: %w", number, err)
		}

		err = r.Close()
		if err != nil {
			return Info{}, fmt.Errorf("close page (%d) file: %w", number, err)
		}
	}

	if !found {
		return Info{}, ErrBookInfoNotFound
	}

	return info, nil
}
