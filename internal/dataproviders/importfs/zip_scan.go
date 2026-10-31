package importfs

import (
	"context"
	"os"
	"path"
)

func (s *Storage) AllZips(_ context.Context) ([]string, error) {
	return s.scanFiles("", true)
}

func (s *Storage) scanFiles(dirPath string, recursive bool) ([]string, error) {
	entries, err := os.ReadDir(path.Join(s.fsPath, dirPath))
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)
	for _, e := range entries {
		name := e.Name()
		relativePath := path.Join(dirPath, name)

		if e.IsDir() {
			if recursive {
				subData, err := s.scanFiles(relativePath, recursive)
				if err != nil {
					return nil, err
				}

				res = append(res, subData...)
			}

			continue
		}

		if path.Ext(name) != ".zip" {
			continue
		}

		res = append(res, relativePath)
	}

	return res, nil
}
