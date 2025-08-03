package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/sk1t0n/echo-mvc-generator/lib"
)

var makeControllerCmd = &cobra.Command{
	Use:   "make:controller arg",
	Short: "Make controller",
	Args:  cobra.ExactArgs(1),
	Example: `make:controller home_controller -> internal/app/http/controllers/home_controller.go
make:controller HomeController -> internal/app/http/controllers/HomeController.go
make:controller controllers/home_controller -> controllers/home_controller.go
make:controller ./controllers/HomeController.go -> controllers/HomeController.go`,
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
		path = "internal/app/http/controllers/" + path
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
	entityNameLower := lib.GetEntityName(path, lib.FormatEntityNameSnakeCase)
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
	return `package controllers

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/open2b/scriggo/native"
)

type {{.EntityName}}Controller struct {
}

func New{{.EntityName}}Controller() {{.EntityName}}Controller {
    return {{.EntityName}}Controller{}
}

func ({{.EntityName}}Controller) Index(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Index | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/{{.EntityNameLower}}/index.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func ({{.EntityName}}Controller) Create(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Create | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/{{.EntityNameLower}}/create.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func ({{.EntityName}}Controller) Store(c echo.Context) error {
    return c.String(http.StatusOK, "Store")
}

func ({{.EntityName}}Controller) Show(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Show | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/{{.EntityNameLower}}/show.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func ({{.EntityName}}Controller) Edit(c echo.Context) error {
    w := c.Response().Writer
    globals := native.Declarations{
        "title": "Edit | Project",
    }
    vars := map[string]any{}

    err := templates.RenderTemplate(w, "internal/templates/{{.EntityNameLower}}/edit.html", globals, vars)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
    }

    return nil
}

func ({{.EntityName}}Controller) Update(c echo.Context) error {
    return c.String(http.StatusOK, "Update")
}

func ({{.EntityName}}Controller) Destroy(c echo.Context) error {
    return c.String(http.StatusOK, "Destroy")
}`
}
