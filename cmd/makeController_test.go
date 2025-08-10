package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/sk1t0n/fiber-mvc-generator/lib"
)

func Test_makeController(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"file", "home_controller", false},
		{"file_with_dirs", "controllers/home_controller", false},
		{"file_with_dirs", "./controllers/category_controller.go", false},
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
			entityNameLower := lib.GetEntityName(tt.path, lib.FormatEntityNameLowerCase)
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
		lib.RemoveFilesAlongWithDir("internal/controller/http")
		lib.RemoveFilesAlongWithDir("internal/controller")
	})
}

func Test_getContentController(t *testing.T) {
	content := getContentController()
	want := `package http

import (
    "github.com/gofiber/fiber/v2"
    "github.com/open2b/scriggo/native"
)

type PostController struct {
}

func NewPostController() PostController {
    return PostController{}
}

func (PostController) Index(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Index | Project",
    }
    vars := map[string]any{}

    template := "web/templates/post/index.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func (PostController) Create(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Create | Project",
    }
    vars := map[string]any{}

    template := "web/templates/post/create.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func (PostController) Store(c *fiber.Ctx) error {
    return c.SendString("Store")
}

func (PostController) Show(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Show | Project",
    }
    vars := map[string]any{}

    template := "web/templates/post/show.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func (PostController) Edit(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Edit | Project",
    }
    vars := map[string]any{}

    template := "web/templates/post/edit.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func (PostController) Update(c *fiber.Ctx) error {
    return c.SendString("Update")
}

func (PostController) Destroy(c *fiber.Ctx) error {
    return c.SendString("Destroy")
}`

	entityName := lib.GetEntityName("post_controller", lib.FormatEntityNamePascalCase)
	entityNameLower := lib.GetEntityName("post_controller", lib.FormatEntityNameLowerCase)
	content = strings.ReplaceAll(content, "{{.EntityName}}", entityName)
	content = strings.ReplaceAll(content, "{{.EntityNameLower}}", entityNameLower)

	if content != want {
		t.Error("content != want")
	}
}
