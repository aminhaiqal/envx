package main

import (
	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <project>",
	Short: "Initialize a new project",
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

		if err := manager.InitProject(projectName, descFlag, envFlag); err != nil {
			return err
		}

		color.Green("✓ Initialized project '%s' with environment '%s'", projectName, envFlag)
		return nil
	},
}

func init() {
	initCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
	initCmd.Flags().StringVarP(&descFlag, "desc", "d", "", "Project description")

	rootCmd.PersistentFlags().Bool("version", false, "Show envx version")
	rootCmd.AddCommand(initCmd)
}
