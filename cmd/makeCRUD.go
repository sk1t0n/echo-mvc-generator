package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/sk1t0n/fiber-mvc-generator/lib"
	"github.com/spf13/cobra"
)

var makeCRUDCmd = &cobra.Command{
	Use:   "make:crud arg",
	Short: "Make CRUD",
	Args:  cobra.ExactArgs(1),
	Example: `make:crud model_name -> internal/controller/http/model_name_controller.go ...
make:crud model_name -c controllers -> controllers/model_name_controller.go ...
make:crud model_name -m models -> models/model_name.go ...
make:crud model_name -v web/templates -> web/templates/modelname/index.html ...`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		err := makeCRUD(cmd, path)
		if err != nil {
			fmt.Println("Failed to create files.")
		} else {
			fmt.Println("Files created successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCRUDCmd)

	makeCRUDCmd.Flags().StringP(
		"dir_controller",
		"c",
		"internal/controller/http",
		"help message for dir_controller",
	)

	makeCRUDCmd.Flags().StringP(
		"dir_model",
		"m",
		"internal/entity",
		"help message for dir_model",
	)

	makeCRUDCmd.Flags().StringP(
		"dir_views",
		"v",
		"web/templates",
		"help message for dir_views",
	)
}

func makeCRUD(cmd *cobra.Command, modelName string) error {
	dirController, _ := cmd.Flags().GetString("dir_controller")
	dirController = lib.RemoveLastSlash(dirController)
	dirModel, _ := cmd.Flags().GetString("dir_model")
	dirModel = lib.RemoveLastSlash(dirModel)
	dirViews, _ := cmd.Flags().GetString("dir_views")
	dirViews = lib.RemoveLastSlash(dirViews)

	pathController := dirController + "/" + modelName + "_controller.go"
	pathModel := dirModel + "/" + modelName + ".go"
	entityNameLower := lib.GetEntityName(modelName, lib.FormatEntityNameLowerCase)
	pathViewIndex := dirViews + "/" + entityNameLower + "/index.html"
	pathViewShow := dirViews + "/" + entityNameLower + "/show.html"
	pathViewCreate := dirViews + "/" + entityNameLower + "/create.html"
	pathViewEdit := dirViews + "/" + entityNameLower + "/edit.html"

	err1 := makeController(pathController)
	err2 := makeModel(pathModel)
	err3 := makeView(pathViewIndex)
	err4 := makeView(pathViewShow)
	err5 := makeView(pathViewCreate)
	err6 := makeView(pathViewEdit)
	err7 := updateRoutes("internal/controller/http/router/router.go", modelName)

	return errors.Join(err1, err2, err3, err4, err5, err6, err7)
}

func updateRoutes(f string, modelName string) error {
	dataRoutes, err := os.ReadFile(f)
	if err != nil {
		return err
	}

	searchedText := "func (r *Router) registerRoutes() {"
	idx := bytes.Index(dataRoutes, []byte(searchedText))
	if idx == -1 {
		return fmt.Errorf(`not found "%s" in file`, searchedText)
	}

	var data string
	for i := idx; i < len(dataRoutes); i++ {
		if dataRoutes[i] == '}' {
			entityName := lib.GetEntityName(modelName, lib.FormatEntityNamePascalCase)
			entityNameLower := lib.GetEntityName(modelName, lib.FormatEntityNameLowerCase)
			callFunc := "\n    r.registerResource(\"" +
				entityNameLower +
				`", http.New` +
				entityName +
				"Controller())\n}"
			data = string(dataRoutes[:i]) + callFunc + string(dataRoutes[i+1:])
			break
		}
	}

	err = os.WriteFile(f, []byte(data), 0666)
	if err != nil {
		return err
	}

	return nil
}
