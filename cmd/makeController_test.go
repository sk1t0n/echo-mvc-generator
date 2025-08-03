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
		{"file:snake_case", "home_controller", false},
		{"file:pascal_case", "HomeController", false},
		{"file_with_dirs:snake_case", "controllers/home_controller", false},
		{"file_with_dirs:pascal_case", "./controllers/HomeController.go", false},
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

			entityName := lib.GetEntityName(tt.path, lib.FormatEntityNamePascalCase)
			entityNameLower := lib.GetEntityName(tt.path, lib.FormatEntityNameSnakeCase)
			content := getContentController()
			content = strings.ReplaceAll(content, "{{.EntityName}}", entityName)
			content = strings.ReplaceAll(content, "{{.EntityNameLower}}", entityNameLower)

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

func Test_getContentController(t *testing.T) {
	content := getContentController()
	want := `package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/open2b/scriggo/native"
)

type PostController struct {
}

func NewPostController() PostController {
    return PostController{}
}

func (PostController) Index(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Index | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/post/index.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func (PostController) Create(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Create | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/post/create.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func (PostController) Store(c echo.Context) error {
    return c.String(http.StatusOK, "Store")
}

func (PostController) Show(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Show | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/post/show.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func (PostController) Edit(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Edit | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/post/edit.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func (PostController) Update(c echo.Context) error {
    return c.String(http.StatusOK, "Update")
}

func (PostController) Destroy(c echo.Context) error {
    return c.String(http.StatusOK, "Destroy")
}`

	entityName := lib.GetEntityName("post_controller", lib.FormatEntityNamePascalCase)
	entityNameLower := lib.GetEntityName("post_controller", lib.FormatEntityNameSnakeCase)
	content = strings.ReplaceAll(content, "{{.EntityName}}", entityName)
	content = strings.ReplaceAll(content, "{{.EntityNameLower}}", entityNameLower)

	if content != want {
		t.Error("content != want")
	}
}
