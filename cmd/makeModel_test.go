package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/sk1t0n/echo-mvc-generator/lib"
)

func Test_makeModel(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"file:lower_case", "user", false},
		{"file:pascal_case", "User", false},
		{"file_with_dirs:lower_case", "models/user", false},
		{"file_with_dirs:pascal_case", "./models/User.go", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := makeModel(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("makeModel(%s) failed: %v", tt.path, gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatalf("makeModel(%s) succeeded unexpectedly", tt.path)
			}

			entityName := lib.GetEntityName(tt.path, lib.FormatEntityNamePascalCase)
			content := `package models

import "gorm.io/gorm"

type {{.EntityName}} struct {
    gorm.Model
}`

			content = strings.ReplaceAll(content, "{{.EntityName}}", entityName)

			data, err := os.ReadFile(tt.path)
			if err == nil && string(data) != content {
				t.Fatalf("makeModel(%s), content is invalid", tt.path)
			}
		})

		t.Cleanup(func() {
			lib.RemoveFilesAlongWithDir("models")
			lib.RemoveFilesAlongWithDir("internal/app/models")
		})
	}
}
