package main

import (
	"fmt"
	"os"

	"github.com/axelyn/envx/internal/exporter"
	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export <project>",
	Short: "Export variables to a .env file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		store, err := storage.New()
		if err != nil {
			return err
		}
		manager := profile.New(store)

		if envFlag == "" {
			envFlag = "development"
		}

		variables, err := manager.ListVariables(projectName, envFlag)
		if err != nil {
			return err
		}
		if len(variables) == 0 {
			return fmt.Errorf("no variables to export")
		}

		if outputFlag == "" {
			outputFlag = ".env"
		}

		if _, err := os.Stat(outputFlag); err == nil && !overwriteFlag {
			color.Yellow("File '%s' already exists. Use --overwrite to replace it.", outputFlag)
			return fmt.Errorf("file already exists")
		}

		exp := exporter.New()
		if err := exp.ExportToDotenv(variables, outputFlag, withCommentsFlag); err != nil {
			return err
		}

		color.Green("✓ Exported %d variables to %s", len(variables), outputFlag)
		return nil
	},
}

func init() {
	exportCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
	exportCmd.Flags().StringVarP(&outputFlag, "output", "o", ".env", "Output file path")
	exportCmd.Flags().BoolVar(&withCommentsFlag, "with-comments", false, "Include descriptions as comments")
	exportCmd.Flags().BoolVar(&overwriteFlag, "overwrite", false, "Overwrite existing file")

	rootCmd.AddCommand(exportCmd)
}
