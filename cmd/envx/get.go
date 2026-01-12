package main

import (
	"fmt"

	"github.com/axelyn/envx/internal/profile"
	"github.com/axelyn/envx/internal/storage"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <project> <KEY>",
	Short: "Get an environment variable",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		key := args[1]

		store, err := storage.New()
		if err != nil {
			return err
		}
		manager := profile.New(store)

		if envFlag == "" {
			envFlag = "development"
		}

		variable, err := manager.GetVariable(projectName, envFlag, key)
		if err != nil {
			return err
		}

		fmt.Println(variable.Value)
		return nil
	},
}

func init() {
	getCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
	rootCmd.AddCommand(getCmd)
}
