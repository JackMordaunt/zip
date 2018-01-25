package zip

import (
	"archive/zip"
	"io"
	"path/filepath"

	"github.com/pkg/errors"
)

func read(file *zip.File, dest string) error {
	if file.FileInfo().IsDir() {
		path := filepath.Join(dest, file.Name)
		return fs.MkdirAll(path, file.FileInfo().Mode())
	}
	zippedF, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "opening zipped file")
	}
	defer zippedF.Close()
	destF, err := fs.Create(dest)
	if err != nil {
		return errors.Wrap(err, "opening destination file")
	}
	defer destF.Close()
	if _, err := io.Copy(destF, zippedF); err != nil {
		return errors.Wrap(err, "copying data")
	}
	return nil
}
