package req

import (
	"encoding/json"
	"io/ioutil"
)

// List a requests with config file
func List(path string) ([]string, error) {
	fileData, err := ioutil.ReadFile(path)
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
		results = append(results, request.Name)
	}

	return results, nil
}
