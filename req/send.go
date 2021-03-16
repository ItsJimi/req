package req

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

func contains(text string, elements []string) bool {
	for _, element := range elements {
		if element == text {
			return true
		}
	}
	return false
}

type Result struct {
	Output    string
	IsCommand bool
}

// Send a request with config file
func Send(names []string, path string) ([]Result, error) {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var requests []Request
	err = json.Unmarshal(fileData, &requests)
	if err != nil {
		return nil, err
	}

	var results []Result
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
			buffer, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			var payload interface{}
			json.Unmarshal(buffer, &payload)
			data := payload.(map[string]interface{})

			if request.Output == "" {
				results = append(results, Result{
					Output:    string(buffer),
					IsCommand: false,
				})
				continue
			}

			t, err := template.New("output").Parse(request.Output)
			if err != nil {
				return nil, err
			}
			buf := new(bytes.Buffer)
			err = t.Execute(buf, data)
			if err != nil {
				return nil, err
			}

			results = append(results, Result{
				Output:    buf.String(),
				IsCommand: true,
			})
		}
	}

	return results, nil
}
