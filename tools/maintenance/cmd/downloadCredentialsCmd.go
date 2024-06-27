package cmd

import (
    "github.com/spf13/cobra"
    "maintenance/internal/sacredential"
)

var serviceAccount string
var credsDir string

var downloadCredentialsCmd = &cobra.Command{
    Use:   "download-creds",
    Short: "Download credentials of a serviceAccount",
    Run: func(cmd *cobra.Command, args []string) {
        if err := sacredential.DownloadCredentials(namespace, serviceAccount, credsDir); err != nil {
            // Handle error
        }
    },
}

func init() {
    rootCmd.AddCommand(downloadCredentialsCmd)
    downloadCredentialsCmd.Flags().StringVarP(&credsDir, "credsDir", "d", "", "Credentials directory (required)")
    downloadCredentialsCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace (required)")
    downloadCredentialsCmd.Flags().StringVarP(&serviceAccount, "serviceAccount", "s", "", "Service account (required)")

    downloadCredentialsCmd.MarkFlagRequired("credsDir")
}
