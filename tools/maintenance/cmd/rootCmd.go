package cmd

import (
    "github.com/spf13/cobra"
)

var namespace string
var sc string
var dryRun bool

var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "CLI tool for Kubernetes operations",
    Long: `A CLI tool built with Cobra that provides several operations
related to Kubernetes, such as PVC deletion, credential download,
port forwarding, storage class management, etc.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
    return rootCmd.Execute()
}
