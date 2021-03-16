package cmd

import (
	"fmt"
	"os/exec"

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
		results, err := req.Send(args, Path)
		if err != nil {
			panic(err)
		}

		if silent {
			return
		}

		for _, result := range results {
			if !result.IsCommand {
				fmt.Println(result.Output)
				continue
			}

			out, err := exec.Command("bash", "-c", result.Output).Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		}
	},
}
