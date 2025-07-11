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

	entityName := lib.GetEntityName(path)
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

	t := template.Must(template.New(entityName).Parse(content))
	data := struct{ ControllerName string }{ControllerName: entityName}
	err = t.Execute(file, data)

	return err
}
