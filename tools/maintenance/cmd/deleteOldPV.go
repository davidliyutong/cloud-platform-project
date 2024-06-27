package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "log"
    "maintenance/internal/pv"
)

// deleteOldPVCmd represents the deleteOldPVs command
var deleteOldPVCmd = &cobra.Command{
    Use:   "rm-old-pvs",
    Short: "Delete old PVs",
    Long:  `Deletes all PVs marked as 'Released' and older than 24 hours.`,
    Run: func(cmd *cobra.Command, args []string) {
        err := pv.DeleteOldPVs(dryRun)
        if err != nil {
            log.Fatalf("Failed to delete old PVs: %v", err)
        }

        fmt.Println("Old PVs deleted successfully")
    },
}

func init() {
    rootCmd.AddCommand(deleteOldPVCmd)

    // add dry run flag
    deleteOldPVCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Dry run")
}
