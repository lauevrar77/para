/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"para.evrard.online/bcs/inbox/services"
	"para.evrard.online/infrastructure/commandbus"
)

var delete bool

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := commandbus.NewContext(context.Background())
		_, err := commandbus.Dispatch(ctx, &services.InboxImportAction{Path: args[0], Delete: delete})
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if delete {
				fmt.Println("File imported in Inbox and deleted from source")

			} else {
				fmt.Println("File imported in Inbox")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	importCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete source file ?")
}
