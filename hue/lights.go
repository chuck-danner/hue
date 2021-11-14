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

func (h Hue) getLights() map[string]light {
	lights := make(map[string]light)
	getRequest(h.getLightsURL(), &lights)
	return lights
}

func (h Hue) getLightsURL() string {
	return fmt.Sprintf("%s/%s", h.getBaseURL(), "lights")
}
