package zip

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"

	fb "github.com/jackmordaunt/filebuilder"
)

func TestZip(t *testing.T) {
	tests := []struct {
		desc    string
		archive string
		source  string
		dest    string
		input   fb.Entry
		wantErr bool
	}{
		{
			desc:    "archive a folder",
			archive: "foo.zip",
			source:  "source",
			dest:    "extracted",
			input: fb.Entries([]fb.Entry{
				fb.Directory{Path: "foo"},
				fb.File{Path: "foo/bar.exe"},
				fb.File{Path: "foo/bar/baz.exe"},
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		SetFs(afero.NewMemMapFs())
		if _, err := fb.Build(fs, "source", tt.input); err != nil {
			t.Fatalf("[%s] error during file setup: %v", tt.desc, err)
		}
		var err error
		err = Create(tt.archive, tt.source)
		if tt.wantErr {
			if err == nil {
				t.Errorf("[%s] want error during archiving, got nil",
					tt.desc)
			}
			continue
		}
		if err != nil {
			t.Fatalf("[%s] unexpected error while creating archive: %v",
				tt.desc, err)
		}
		err = Extract(tt.archive, tt.dest)
		if tt.wantErr {
			if err == nil {
				t.Errorf("[%s] want error during extraction, got nil",
					tt.desc)
			}
			continue
		} else {
			if err != nil {
				t.Errorf("[%s] unexpected error while extracting archive: %v",
					tt.desc, err)
			}
		}
		dest := filepath.Join(tt.dest, tt.source)
		if diff, ok, err := fb.CompareDirectories(fs, tt.source, dest); err != nil || !ok {
			t.Errorf("[%s] source != extracted, err: %v, diff: \n%v", tt.desc, err, diff)
		}
	}
}
