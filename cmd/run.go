package cmd

import (
	"github.com/ItsJimi/req/req"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run <request name>",
	Short: "Run a request",
	Long:  `Run a request`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req.Send(args[0])
	},
}
