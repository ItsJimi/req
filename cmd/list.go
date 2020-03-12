package cmd

import (
	"github.com/ItsJimi/req/req"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all requests names",
	Long:  `List all requests names`,
	Run: func(cmd *cobra.Command, args []string) {
		req.List()
	},
}