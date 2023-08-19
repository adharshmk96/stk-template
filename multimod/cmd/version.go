package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display the version of stktemplate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("stktemplate version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
