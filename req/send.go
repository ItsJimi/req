package req

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func contains(text string, elements []string) bool {
	for _, element := range elements {
		if element == text {
			return true
		}
	}
	return false
}

// Send a request with config file
func Send(names []string) ([]string, error) {
	fileData, err := ioutil.ReadFile(os.Getenv("HOME") + "/.req.json")
	if err != nil {
		return nil, err
	}

	var requests []Request
	err = json.Unmarshal(fileData, &requests)
	if err != nil {
		return nil, err
	}

	var results []string
	for _, request := range requests {
		if contains(request.Name, names) == true {
			client := &http.Client{}

			var body []byte
			if request.Body != nil {
				body, err = json.Marshal(request.Body)
				if err != nil {
					return nil, err
				}
			}

			req, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(body))
			if err != nil {
				return nil, err
			}

			for _, header := range request.Headers {
				splittedHeader := strings.Split(header, ":")
				req.Header.Add(splittedHeader[0], strings.TrimSpace(splittedHeader[1]))
			}
			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			scanner := bufio.NewScanner(resp.Body)
			for i := 0; scanner.Scan(); i++ {
				results = append(results, scanner.Text())
			}
		}
	}

	return results, nil
}
