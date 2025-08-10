package cmd

import (
	"testing"

	"github.com/sk1t0n/fiber-mvc-generator/lib"
)

func Test_makeView(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"file", "index", false},
		{"file_with_dirs", "templates/post/index", false},
		{"file_with_dirs", "./templates/index.html", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := makeView(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("makeView(%s) failed: %v", tt.path, gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatalf("makeView(%s) succeeded unexpectedly", tt.path)
			}
		})

		t.Cleanup(func() {
			lib.RemoveFilesAlongWithDir("templates/post")
			lib.RemoveFilesAlongWithDir("templates")
			lib.RemoveFilesAlongWithDir("web/templates")
			lib.RemoveFilesAlongWithDir("web")
		})
	}
}
