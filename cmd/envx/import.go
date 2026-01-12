package main

import (
	"fmt"

	"github.com/axelyn/envx/pkg/envx"
	"github.com/axelyn/envx/internal/importer"
	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import <project> <file>",
	Short: "Import variables from a .env file",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		filePath := args[1]

		store, err := storage.New()
		if err != nil {
			return err
		}
		manager := profile.New(store)

		if envFlag == "" {
			envFlag = "development"
		}

		imp := importer.New()
		variables, err := imp.ImportFromDotenv(filePath)
		if err != nil {
			return err
		}
		if len(variables) == 0 {
			return fmt.Errorf("no variables found in file")
		}

		existing, _ := manager.ListVariables(projectName, envFlag)
		if existing == nil {
			existing = make(map[string]envx.Variable)
		}

		newVars, updatedVars, unchangedVars, err := imp.PreviewImport(filePath, existing)
		if err != nil {
			return err
		}

		if dryRunFlag {
			color.Cyan("\nImport Preview for %s (%s)\n", projectName, envFlag)
			if len(newVars) > 0 {
				color.Green("\n✓ New variables (%d):", len(newVars))
				for _, key := range newVars {
					fmt.Printf("  + %s\n", key)
				}
			}
			if len(updatedVars) > 0 {
				color.Yellow("\nUpdated variables (%d):", len(updatedVars))
				for _, key := range updatedVars {
					fmt.Printf("  ~ %s\n", key)
				}
			}
			if len(unchangedVars) > 0 {
				color.Blue("\n- Unchanged variables (%d):", len(unchangedVars))
				for _, key := range unchangedVars {
					fmt.Printf("  = %s\n", key)
				}
			}
			fmt.Println("\nRun without --dry-run to apply changes")
			return nil
		}

		imported := 0
		skipped := 0
		for key, variable := range variables {
			if mergeFlag {
				if _, exists := existing[key]; exists {
					skipped++
					continue
				}
			}
			if err := manager.SetVariable(projectName, envFlag, key, variable.Value, "", false); err != nil {
				color.Yellow("Failed to import %s: %v", key, err)
				continue
			}
			imported++
		}

		color.Green("✓ Imported %d variables", imported)
		if skipped > 0 {
			color.Yellow("⚠ Skipped %d existing variables (merge mode)", skipped)
		}

		return nil
	},
}

func init() {
	importCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
	importCmd.Flags().BoolVar(&mergeFlag, "merge", false, "Don't overwrite existing variables")
	importCmd.Flags().BoolVar(&dryRunFlag, "dry-run", false, "Preview changes without applying")

	rootCmd.AddCommand(importCmd)
}
