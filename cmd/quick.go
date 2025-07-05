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

// quickCmd represents the quick command
var quickCmd = &cobra.Command{
	Use:   "quick",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := commandbus.NewContext(context.Background())
		_, err := commandbus.Dispatch(ctx, &services.InboxQuickNoteAction{})
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(quickCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quickCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quickCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
