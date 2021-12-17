package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type lightResponseV2 struct {
	Errors []hueError `json:"errors"`
	Data   []light    `json:"data"`
}

type hueError struct {
	Description string `json:"description"`
}

type light struct {
	Id       string   `json:"id,omitempty"`
	Type     string   `json:"type,omitempty"`
	IDv1     string   `json:"id_v1,omitempty"`
	On       state    `json:"on,omitempty"`
	Metadata metadata `json:"metadata,omitempty"`
}

type metadata struct {
	Archetype string `json:"archetype,omitempty"`
	Name      string `json:"name,omitempty"`
}
type state struct {
	On bool `json:"on"`
}

type ResourceIdentifier struct {
	Rid   string `json:"rid,omitempty"`
	Rtype string `json:"rtype,omitempty"`
}

func (h Hue) getLights() lightResponseV2 {
	lights := lightResponseV2{}
	h.getRequest(h.getLightsURL(), &lights)
	return lights
}

func (h Hue) setLight(id string, light light) {
	response := ResourceIdentifier{}
	requestBody, _ := json.Marshal(light)

	url := fmt.Sprintf("%s/%s", h.getLightsURL(), id)
	h.putRequest(url, bytes.NewBuffer(requestBody), &response)
}

func (h Hue) getLightsURL() string {
	return fmt.Sprintf("%s/%s", h.getV2BaseURL(), "light")
}
