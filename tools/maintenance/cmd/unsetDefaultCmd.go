package cmd

import (
    "github.com/spf13/cobra"
    "maintenance/internal/storageclass"
)

var unsetDefaultCmd = &cobra.Command{
    Use:   "unset-default-sc",
    Short: "Unset default storage class",
    Run: func(cmd *cobra.Command, args []string) {
        if err := storageclass.UnsetDefault(sc); err != nil {
            // Handle error
        }
    },
}

func init() {
    rootCmd.AddCommand(unsetDefaultCmd)
    unsetDefaultCmd.Flags().StringVarP(&sc, "sc", "s", "", "Storage class (required)")
    unsetDefaultCmd.MarkFlagRequired("sc")
}
