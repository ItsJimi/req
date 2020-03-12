package req

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// List a requests with config file
func List() {
	fileData, err := ioutil.ReadFile(os.Getenv("HOME") + "/.req.json")
	if err != nil {
		panic(err)
	}

	var requests []Request
	err = json.Unmarshal(fileData, &requests)
	if err != nil {
		panic(err)
	}

	for _, request := range requests {
		fmt.Println(request.Name)
	}
}
