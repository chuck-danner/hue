package hue

import (
	"fmt"
)

type config struct {
	Name       string `json:"name"`
	ApivVrsion string `json:"apiversion"`
}

func listConfig() config {
	var configuration config
	getRequest(getConfigURL(), &configuration)
	return configuration
}

func getConfigURL() string {
	return fmt.Sprintf("%s/%s", getBaseURL(), "config")
}
