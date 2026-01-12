package main

import (
	"fmt"
	"strings"

	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <project> <KEY=value>",
	Short: "Set an environment variable",
	Args:  cobra.MinimumNArgs(2),
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

		for _, pair := range args[1:] {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("invalid format: %s (expected KEY=value)", pair)
			}
			key, value := parts[0], parts[1]

			if err := manager.SetVariable(projectName, envFlag, key, value, descFlag, false); err != nil {
				return err
			}
			color.Green("✓ Set %s=%s", key, value)
		}

		return nil
	},
}

func init() {
	setCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
	setCmd.Flags().StringVarP(&descFlag, "desc", "d", "", "Variable description")

	rootCmd.AddCommand(setCmd)
}
