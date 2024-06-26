/*
Copyright © 2024 Rohaizan Roosley rohaizanr@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "masjid",
	Short: "E-Masjid.My Saas Command Line Interface",
	Long: `
masjid is a CLI for managing E-Masjid.My Saas.
This tool communicates with the E-Masjid.My Saas database through gRPC.`,
}

func Execute() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
