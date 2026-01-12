package main

import (
	"fmt"

	"github.com/axelyn/envx/internal/exporter"
	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template <project>",
	Short: "Generate a template file (without values)",
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

		exp := exporter.New()
		if err := exp.ExportTemplate(variables, outputFlag); err != nil {
			return err
		}

		if outputFlag != "" {
			color.Green("✓ Template exported to %s", outputFlag)
		}
		return nil
	},
}

func init() {
	templateCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
	templateCmd.Flags().StringVarP(&outputFlag, "output", "o", "", "Output file path (default: stdout)")

	rootCmd.AddCommand(templateCmd)
}
