package cmd

import (
	"fmt"

	"github.com/ItsJimi/req/req"
	"github.com/spf13/cobra"
)

var silent bool

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Silent run command")
}

var runCmd = &cobra.Command{
	Use:   "run <request name>",
	Short: "Run a request",
	Long:  `Run a request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results, err := req.Send(args)
		if err != nil {
			panic(err)
		}

		if silent == true {
			return
		}

		for _, result := range results {
			fmt.Println(result)
		}
	},
}
