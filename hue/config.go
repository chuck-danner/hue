package hue

import (
	"fmt"
)

type config struct {
	Name       string `json:"name"`
	ApivVrsion string `json:"apiversion"`
}

func (h Hue) getConfig() config {
	var configuration config
	getRequest(h.getConfigURL(), &configuration)
	return configuration
}

func (h Hue) getConfigURL() string {
	return fmt.Sprintf("%s/%s", h.getBaseURL(), "config")
}
