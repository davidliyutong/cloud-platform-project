package cmd

import (
    "github.com/spf13/cobra"
    "maintenance/internal/storageclass"
)

var setDefaultCmd = &cobra.Command{
    Use:   "set-default-sc",
    Short: "Set default storage class",
    Run: func(cmd *cobra.Command, args []string) {
        if err := storageclass.SetDefault(sc); err != nil {
            // Handle error
        }
    },
}

func init() {
    rootCmd.AddCommand(setDefaultCmd)
    setDefaultCmd.Flags().StringVarP(&sc, "sc", "s", "", "Storage class (required)")
    setDefaultCmd.MarkFlagRequired("sc")
}
