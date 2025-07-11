/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"para.evrard.online/bcs/projects/services"
	"para.evrard.online/infrastructure/commandbus"
)

// projectImportCmd represents the projectImport command
var projectImportCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := commandbus.NewContext(context.Background())
		searchString := strings.Join(args[1:], " ")
		_, err := commandbus.Dispatch(ctx, &services.ProjectImportAction{SearchString: searchString, Path: args[0], Delete: delete})
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if delete {
				fmt.Println("File imported in Project and deleted from source")

			} else {
				fmt.Println("File imported in Project")
			}
		}
	},
}

func init() {
	projectCmd.AddCommand(projectImportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectImportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectImportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectImportCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete source file ?")
}
