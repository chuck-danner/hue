package hue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type config struct {
	Name       string `json:"name"`
	ApivVrsion string `json:"apiversion"`
}

func listConfig() config {

	resp, err := http.Get(getConfigURL())

	if err != nil {
		panic(err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var configuration config
	err = json.Unmarshal(data, &configuration)

	if err != nil {
		panic(err.Error())
	}

	return configuration
}

func getConfigURL() string {
	return fmt.Sprintf("%s/%s", getBaseURL(), "config")
}
