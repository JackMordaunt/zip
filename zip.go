package zip

import (
	"archive/zip"
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"
)

// Create a zip archive at `archive` with the contents of `source`.
func Create(archive, source string) error {
	wr, err := fs.Create(archive)
	if err != nil {
		return errors.Wrapf(err, "creating archive: %v", archive)
	}
	defer wr.Close()
	if err := create(wr, source); err != nil {
		return errors.Wrap(err, "writing archive")
	}
	return nil
}

// Extract the contents of a zip archive into `dest`.
func Extract(archive, dest string) error {
	_ = fs.MkdirAll(dest, 0755)
	info, err := fs.Stat(dest)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("destination '%s' is not a directory", dest)
	}
	r, err := fs.Open(archive)
	if err != nil {
		return errors.Wrap(err, "opening archive")
	}
	defer r.Close()
	rinfo, err := r.Stat()
	if err != nil {
		return errors.Wrap(err, "could not get size of archive")
	}
	zr, err := zip.NewReader(r, rinfo.Size())
	if err != nil {
		return errors.Wrap(err, "creating zip reader")
	}
	for _, file := range zr.File {
		if err := extract(file, filepath.Join(dest, file.Name)); err != nil {
			return errors.Wrap(err, "unzipping file")
		}
	}
	return nil
}
