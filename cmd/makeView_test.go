package cmd

import (
	"testing"

	"github.com/sk1t0n/echo-mvc-generator/lib"
)

func Test_makeView(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"file: lower case", "index", false},
		{"file with dirs: lower case", "templates/index", false},
		{"file with dirs: lower case", "./templates/index.html", false},
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
			lib.RemoveFilesAlongWithDir("templates")
			lib.RemoveFilesAlongWithDir("internal/templates")
			lib.RemoveFilesAlongWithDir("internal")
		})
	}
}
