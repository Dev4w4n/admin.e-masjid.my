/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits tenant in E-Masjid.My Saas using json file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
		filePath, _ := cmd.Flags().GetString("file")
		fmt.Println("edit called with file path:", filePath)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("file", "f", "", "Path to JSON file")
}
