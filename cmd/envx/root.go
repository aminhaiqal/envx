package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	commit  = "none"
	date    = "13-01-2026"
)

var rootCmd = &cobra.Command{
	Use:   "envx",
	Short: "⚡ Lightning-fast environment variable management",
	Long:  `envx - A blazingly fast, local-first CLI tool to manage environment variables across all your projects.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Printf("envx %s\ncommit: %s\nbuilt at: %s\n", version, commit, date)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, check version too
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Printf("envx %s\ncommit: %s\nbuilt at: %s\n", version, commit, date)
			os.Exit(0)
		}
	},
}
