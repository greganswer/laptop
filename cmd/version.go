package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of laptop",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("laptop version %s\n", version)
	},
}
