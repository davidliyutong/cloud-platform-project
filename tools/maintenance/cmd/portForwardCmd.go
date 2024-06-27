package cmd

import (
    "github.com/spf13/cobra"
    "maintenance/internal/common"
)

var serviceName string
var port string

var portForwardCmd = &cobra.Command{
    Use:   "port-forward",
    Short: "Port forward a service to local",
    Run: func(cmd *cobra.Command, args []string) {
        common.PortForward(namespace, serviceName, port)
    },
}

func init() {
    rootCmd.AddCommand(portForwardCmd)
    portForwardCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace (required)")
    portForwardCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Service name (required)")
    portForwardCmd.Flags().StringVarP(&port, "port", "p", "", "Port number (required): <local_port>:<remote_port> ")
    portForwardCmd.MarkFlagRequired("service")
    portForwardCmd.MarkFlagRequired("port")
}
