package hue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type light struct {
	Name  string `json:"name"`
	State state  `json:"state"`
}

type state struct {
	On bool `json:"on"`
}

func listLights() map[string]light {

	resp, err := http.Get(getLightsURL())

	if err != nil {
		panic(err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	lights := make(map[string]light)
	//fmt.Println(string(data))

	err = json.Unmarshal(data, &lights)

	if err != nil {
		panic(err.Error())
	}

	return lights
}

func getLightsURL() string {
	return fmt.Sprintf("%s/%s", getBaseURL(), "lights")
}
