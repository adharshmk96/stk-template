package cmd

import (
	"github.com/spf13/cobra"
	"github.com/adharshmk96/stk-template/server"
)

var startingPort string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		startAddr := "0.0.0.0:"
		server.StartServer(startAddr + startingPort)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&startingPort, "port", "p", "8080", "Port to start the server on")

	rootCmd.AddCommand(serveCmd)
}
