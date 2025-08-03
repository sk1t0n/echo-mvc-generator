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

var makeModelCmd = &cobra.Command{
	Use:   "make:model arg",
	Short: "Make model",
	Args:  cobra.ExactArgs(1),
	Example: `make:model user -> internal/app/models/user.go
make:model User -> internal/app/models/User.go
make:model models/user -> models/user.go
make:model ./models/User.go -> models/User.go`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		err := makeModel(path)
		if err != nil {
			fmt.Println("Failed to create file.")
		} else {
			fmt.Println("File created successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(makeModelCmd)
}

func makeModel(path string) error {
	if !strings.HasSuffix(path, ".go") {
		path += ".go"
	}

	if !strings.Contains(path, "/") && !strings.Contains(path, "\\") {
		path = "internal/app/models/" + path
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
	content := `package models

import "gorm.io/gorm"

type {{.EntityName}} struct {
    gorm.Model
}`

	t := template.Must(template.New(entityName).Parse(content))
	data := struct{ EntityName string }{EntityName: entityName}
	err = t.Execute(file, data)

	return err
}
