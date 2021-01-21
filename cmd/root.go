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

// Path expose config path
var Path string

// Execute cli
func Execute() {
	rootCmd.PersistentFlags().StringVarP(&Path, "config", "c", os.Getenv("HOME")+"/.req.json", "Path of config file (Must be a JSON file)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
