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
	"para.evrard.online/bcs/projects/domain"
	"para.evrard.online/bcs/projects/services"
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

		var lastTitle *domain.MdTitle = nil
		titleChanged := false
		for _, todo := range query.Todos {
			titleChanged = false
			if todo.ParentTitle != nil {
				if lastTitle == nil || lastTitle.LineNumber != todo.ParentTitle.LineNumber {
					titleChanged = true
					fmt.Println(todo.FilePath)
					for i := 0; i < todo.ParentTitle.Level; i++ {
						fmt.Printf("#")
					}
					fmt.Printf(" %s\n", todo.ParentTitle.Title)

				}
				lastTitle = todo.ParentTitle
			}
			if todo.Done {
				fmt.Printf("- [x]")
			} else {
				fmt.Printf("- [ ]")
			}
			fmt.Printf(" %s\n", todo.Todo)
			if titleChanged {
				fmt.Println("--------------------\n")
			}
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
