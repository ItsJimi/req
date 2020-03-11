package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "req",
	Short: "Req send requests from config file",
	Long:  `Req send requests from config file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Type `req help` to view commands available")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
