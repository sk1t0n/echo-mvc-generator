package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sk1t0n/fiber-mvc-generator/lib"
)

var makeViewCmd = &cobra.Command{
	Use:   "make:view arg",
	Short: "Make view",
	Args:  cobra.ExactArgs(1),
	Example: `make:view index -> web/templates/index.html
make:view templates/post/index -> templates/post/index.html
make:view ./templates/index.html -> templates/index.html`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		err := makeView(path)
		if err != nil {
			fmt.Println("Failed to create file.")
		} else {
			fmt.Println("File created successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(makeViewCmd)
}

func makeView(path string) error {
	if !strings.HasSuffix(path, ".html") {
		path += ".html"
	}

	if !strings.Contains(path, "/") && !strings.Contains(path, "\\") {
		path = "web/templates/" + path
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
	content := `{% extends "/web/templates/layouts/base.html" %}

{% macro Body %}
  <h1>{{ title }}</h1>
{% end %}`
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return err
	}

	return err
}
