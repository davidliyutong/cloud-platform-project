package cmd

import (
    "github.com/spf13/cobra"
    "maintenance/internal/storageclass"
)

var getAllStorageClassesCmd = &cobra.Command{
    Use:   "get-all-sc",
    Short: "Get all storage classes",
    Run: func(cmd *cobra.Command, args []string) {
        if err := storageclass.GetAllStorageClasses(); err != nil {
            // Handle error
        }
    },
}

func init() {
    rootCmd.AddCommand(getAllStorageClassesCmd)
}
