package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		sendRequest(args[0])
	},
}

type request struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Method string `json:"method"`
}

func sendRequest(name string) {
	fileData, err := ioutil.ReadFile("/home/filedescriptor/.req.json")
	if err != nil {
		panic(err)
	}

	var requests []request
	err = json.Unmarshal(fileData, &requests)
	if err != nil {
		panic(err)
	}

	for _, request := range requests {
		if request.Name == name {
			resp, err := http.Get(request.URL)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			fmt.Println("Response status:", resp.Status)
			return
		}
	}

	fmt.Println("Request not found")
}
