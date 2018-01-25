package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

func write(wr io.Writer, source string) error {
	info, err := fs.Stat(source)
	if err != nil {
		return errors.Wrap(err, "could not open source directory")
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", source)
	}
	zwr := zip.NewWriter(wr)
	defer zwr.Close()
	if err := _write(zwr, source); err != nil {
		return errors.Wrap(err, "writing archive")
	}
	return nil
}

func _write(zwr *zip.Writer, source string) error {
	info, err := fs.Stat(source)
	if err != nil {
		return err
	}
	if info.IsDir() {
		entries, err := afero.ReadDir(fs, source)
		if err != nil {
			return errors.Wrap(err, "walking filesystem")
		}
		for _, e := range entries {
			path := filepath.Join(source, e.Name())
			if err := _write(zwr, path); err != nil {
				return err
			}
		}
		return nil
	}
	zipF, err := zwr.Create(source)
	if err != nil {
		return err
	}
	sourceF, err := fs.Open(source)
	if err != nil {
		return err
	}
	defer sourceF.Close()
	if _, err := io.Copy(zipF, sourceF); err != nil {
		return err
	}
	return nil
}
