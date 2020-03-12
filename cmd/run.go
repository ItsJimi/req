package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Method  string   `json:"method"`
	Headers []string `json:"headers"`
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
			client := &http.Client{}

			req, err := http.NewRequest(request.Method, request.URL, nil)

			for _, header := range request.Headers {
				splittedHeader := strings.Split(header, ":")
				req.Header.Add(splittedHeader[0], strings.TrimSpace(splittedHeader[1]))
			}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			scanner := bufio.NewScanner(resp.Body)
			for i := 0; scanner.Scan(); i++ {
				fmt.Println(scanner.Text())
			}
			return
		}
	}

	fmt.Println("Request not found")
}
