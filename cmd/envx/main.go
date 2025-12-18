package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/axelyn/envx/internal/profile"
    "github.com/axelyn/envx/internal/storage"
    "github.com/fatih/color"
    "github.com/spf13/cobra"
)

var (
    envFlag  string
    descFlag string
)

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

var rootCmd = &cobra.Command{
    Use:   "envx",
    Short: "⚡ Lightning-fast environment variable management",
    Long:  `envx - A blazingly fast, local-first CLI tool to manage environment variables across all your projects.`,
}

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

        // Parse KEY=value pairs
        for _, pair := range args[1:] {
            parts := strings.SplitN(pair, "=", 2)
            if len(parts) != 2 {
                return fmt.Errorf("invalid format: %s (expected KEY=value)", pair)
            }

            key := parts[0]
            value := parts[1]

            if err := manager.SetVariable(projectName, envFlag, key, value, descFlag, false); err != nil {
                return err
            }

            color.Green("✓ Set %s=%s", key, value)
        }

        return nil
    },
}

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
                // Mask secret values
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
    // Add flags
    initCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
    initCmd.Flags().StringVarP(&descFlag, "desc", "d", "", "Project description")

    setCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
    setCmd.Flags().StringVarP(&descFlag, "desc", "d", "", "Variable description")

    getCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")
    listCmd.Flags().StringVarP(&envFlag, "env", "e", "", "Environment name")

    // Add commands to root
    rootCmd.AddCommand(initCmd)
    rootCmd.AddCommand(setCmd)
    rootCmd.AddCommand(getCmd)
    rootCmd.AddCommand(listCmd)
}