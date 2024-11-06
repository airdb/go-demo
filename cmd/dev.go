/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"demo/sandbox/sqlite"
	"fmt"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Development sandbox command",
	Long: `Development sandbox command for testing and experimenting with features.
Currently supports:
- SQLite database operations demo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dev called")
		// emailutil.RunSendEmail()
		sqlite.RunDemo()
	},
}

func init() {
	rootCmd.AddCommand(devCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
