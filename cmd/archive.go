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

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := commandbus.NewContext(context.Background())
		searchString := strings.Join(args, " ")
		_, err := commandbus.Dispatch(ctx, &services.ArchiveProjectAction{SearchString: searchString})
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	projectCmd.AddCommand(archiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// archiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// archiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
