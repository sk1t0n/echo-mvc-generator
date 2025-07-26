package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sk1t0n/echo-mvc-generator/lib"
	"github.com/spf13/cobra"
)

var makeCRUDCmd = &cobra.Command{
	Use:   "make:crud arg",
	Short: "Make CRUD",
	Args:  cobra.ExactArgs(1),
	Example: `make:crud model_name -> internal/app/http/controllers/model_name_controller.go ...
make:crud ModelName -> internal/app/http/controllers/ModelNameController.go ...
make:crud model_name -c controllers -> controllers/model_name_controller.go ...
make:crud model_name -m models -> models/model_name.go ...
make:crud model_name -v templates -> templates/index.html ...`,
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
		"internal/app/http/controllers",
		"help message for dir_controller",
	)

	makeCRUDCmd.Flags().StringP(
		"dir_model",
		"m",
		"internal/app/models",
		"help message for dir_model",
	)

	makeCRUDCmd.Flags().StringP(
		"dir_views",
		"v",
		"internal/templates",
		"help message for dir_views",
	)
}

func makeCRUD(cmd *cobra.Command, modelName string) error {
	hasSnakeCase := strings.Contains(modelName, "_")

	dirController, _ := cmd.Flags().GetString("dir_controller")
	dirController = lib.RemoveLastSlash(dirController)
	dirModel, _ := cmd.Flags().GetString("dir_model")
	dirModel = lib.RemoveLastSlash(dirModel)
	dirViews, _ := cmd.Flags().GetString("dir_views")
	dirViews = lib.RemoveLastSlash(dirViews)

	var pathController string
	var pathModel string = dirModel + "/" + modelName + ".go"
	var pathViewIndex string = dirViews + "/" + modelName + "/index.html"

	if hasSnakeCase || lib.IsLower(modelName) {
		pathController = dirController + "/" + modelName + "_controller.go"
	} else {
		pathController = dirController + "/" + modelName + "Controller.go"
	}

	err1 := makeController(pathController)
	err2 := makeModel(pathModel)
	err3 := makeView(pathViewIndex)

	return errors.Join(err1, err2, err3)
}
