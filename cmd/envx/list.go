package main

import (
	"fmt"
	"strings"

	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list <project>",
	Short: "List all environment variables",
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

		color.Cyan("\n📦 %s (%s)\n", projectName, envFlag)
		fmt.Println()

		if len(variables) == 0 {
			color.Yellow("No variables set")
			return nil
		}

		for key, variable := range variables {
			value := variable.Value
			if variable.IsSecret {
				if len(value) > 8 {
					value = value[:4] + strings.Repeat("*", len(value)-8) + value[len(value)-4:]
				} else {
					value = strings.Repeat("*", len(value))
				}
			}
			fmt.Printf("%-20s %s\n", key, value)
		}

		fmt.Printf("\n%d variables\n\n", len(variables))
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
	rootCmd.AddCommand(listCmd)
}
