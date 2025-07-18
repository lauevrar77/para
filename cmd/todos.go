/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"para.evrard.online/bcs/projects/services"
	"para.evrard.online/bcs/shared"
	"para.evrard.online/infrastructure/commandbus"
)

// todosCmd represents the todos command
var todosCmd = &cobra.Command{
	Use:   "todos",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := commandbus.NewContext(context.Background())
		searchString := strings.Join(args, " ")
		query := &services.ProjectTodosQuery{SearchString: searchString}
		_, err := commandbus.Query(ctx, query)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		for _, doc := range query.Documents {
			shared.PrintMdDocument(doc)
		}
	},
}

func init() {
	projectCmd.AddCommand(todosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
