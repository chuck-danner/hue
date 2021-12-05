package hue

import (
	"fmt"
)

type check struct {
	Name       string `json:"name"`
	ApiVersion string `json:"apiversion"`
	Swversion  string `json:"swversion"`
	Mac        string `json:"mac"`
	BridgeId   string `json:"bridgeid"`
	ModelId    string `json:"modelid"`
}

func (h Hue) getCheck() check {
	var checkConfig check
	getRequest(h.getCheckURL(), &checkConfig)
	return checkConfig
}

func (h Hue) getCheckURL() string {
	return fmt.Sprintf("%s/%s", h.URL, "api/0/config")
}
