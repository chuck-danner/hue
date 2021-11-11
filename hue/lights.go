package hue

import (
	"fmt"
)

type light struct {
	Name  string `json:"name"`
	State state  `json:"state"`
}

type state struct {
	On bool `json:"on"`
}

func listLights() map[string]light {
	lights := make(map[string]light)
	getRequest(getLightsURL(), &lights)
	return lights
}

func getLightsURL() string {
	return fmt.Sprintf("%s/%s", getBaseURL(), "lights")
}
