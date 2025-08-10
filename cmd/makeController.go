package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/sk1t0n/fiber-mvc-generator/lib"
)

var makeControllerCmd = &cobra.Command{
	Use:   "make:controller arg",
	Short: "Make controller",
	Args:  cobra.ExactArgs(1),
	Example: `make:controller home_controller -> internal/controller/http/home_controller.go
make:controller controllers/home_controller -> controllers/home_controller.go
make:controller ./controllers/home_controller.go -> controllers/home_controller.go`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		err := makeController(path)
		if err != nil {
			fmt.Println("Failed to create file.")
		} else {
			fmt.Println("File created successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(makeControllerCmd)
}

func makeController(path string) error {
	if !strings.HasSuffix(path, ".go") {
		path += ".go"
	}

	if !strings.Contains(path, "/") && !strings.Contains(path, "\\") {
		path = "internal/controller/http/" + path
	}

	err := lib.MkdirAll(path)
	if err != nil {
		return err
	}

	file, err := lib.CreateFile(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return err
	}

	entityName := lib.GetEntityName(path, lib.FormatEntityNamePascalCase)
	entityNameLower := lib.GetEntityName(path, lib.FormatEntityNameLowerCase)
	content := getContentController()

	t := template.Must(template.New(entityName).Parse(content))
	data := struct {
		EntityName      string
		EntityNameLower string
	}{
		EntityName:      entityName,
		EntityNameLower: entityNameLower,
	}
	err = t.Execute(file, data)

	return err
}

func getContentController() string {
	return `package http

import (
    "github.com/gofiber/fiber/v2"
    "github.com/open2b/scriggo/native"
)

type {{.EntityName}}Controller struct {
}

func New{{.EntityName}}Controller() {{.EntityName}}Controller {
    return {{.EntityName}}Controller{}
}

func ({{.EntityName}}Controller) Index(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Index | Project",
    }
    vars := map[string]any{}

    template := "web/templates/{{.EntityNameLower}}/index.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func ({{.EntityName}}Controller) Create(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Create | Project",
    }
    vars := map[string]any{}

    template := "web/templates/{{.EntityNameLower}}/create.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func ({{.EntityName}}Controller) Store(c *fiber.Ctx) error {
    return c.SendString("Store")
}

func ({{.EntityName}}Controller) Show(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Show | Project",
    }
    vars := map[string]any{}

    template := "web/templates/{{.EntityNameLower}}/show.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func ({{.EntityName}}Controller) Edit(c *fiber.Ctx) error {
    globals := native.Declarations{
        "title": "Edit | Project",
    }
    vars := map[string]any{}

    template := "web/templates/{{.EntityNameLower}}/edit.html"
    err := templates.RenderTemplate(c, template, globals, vars)
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "failed to render template "+template,
        )
    }

    return nil
}

func ({{.EntityName}}Controller) Update(c *fiber.Ctx) error {
    return c.SendString("Update")
}

func ({{.EntityName}}Controller) Destroy(c *fiber.Ctx) error {
    return c.SendString("Destroy")
}`
}
