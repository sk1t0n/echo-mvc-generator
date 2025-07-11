package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/sk1t0n/echo-mvc-generator/lib"
)

func Test_makeController(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"file: snake case", "home_controller", false},
		{"file: pascal case", "HomeController", false},
		{"file with dirs: snake case", "controllers/home_controller", false},
		{"file with dirs: pascal case", "./controllers/HomeController.go", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := makeController(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("makeController(%s) failed: %v", tt.path, gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatalf("makeController(%s) succeeded unexpectedly", tt.path)
			}

			entityName := lib.GetEntityName(tt.path)
			content := `package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type {{.ControllerName}} struct {
}

func New{{.ControllerName}}() {{.ControllerName}} {
    return {{.ControllerName}}{}
}

func ({{.ControllerName}}) Index(c echo.Context) error {
    return c.String(http.StatusOK, "Index")
}`
			content = strings.ReplaceAll(content, "{{.ControllerName}}", entityName)

			data, err := os.ReadFile(tt.path)
			if err == nil && string(data) != content {
				t.Fatalf("makeController(%s), content is invalid", tt.path)
			}
		})
	}

	t.Cleanup(func() {
		lib.RemoveFilesAlongWithDir("controllers")
		lib.RemoveFilesAlongWithDir("internal/app/http/controllers")
		lib.RemoveFilesAlongWithDir("internal/app/http")
	})
}
