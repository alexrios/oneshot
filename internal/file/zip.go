package file

import (
	z "archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func zip(path string, w io.Writer) error {
	zw := z.NewWriter(w)
	defer zw.Close()

	return filepath.Walk(path, func(fp string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(fp, filepath.Dir(path))
		zFile, err := zw.Create(relPath)
		if err != nil {
			return err
		}
		currFile, err := os.Open(fp)
		if err != nil {
			return err
		}
		_, err = io.Copy(zFile, currFile)
		currFile.Close()
		if err != nil {
			return err
		}
		return nil
	})
}
