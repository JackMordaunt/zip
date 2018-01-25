package zip

import (
	"github.com/spf13/afero"
)

var fs afero.Fs

func init() {
	fs = afero.NewOsFs()
}

// SetFs to the provided implementation.
func SetFs(filesystem afero.Fs) {
	fs = filesystem
}
