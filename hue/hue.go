package hue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type light struct {
	Name string `json:"name"`
}

func List() map[string]light {

	listURL := getBaseURL() + "/lights"
	println(listURL)
	resp, err := http.Get(listURL)

	if err != nil {
		panic(err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	lights := make(map[string]light)

	err = json.Unmarshal(data, &lights)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(lights)

	return lights
}

func getBaseURL() string {
	username := viper.GetString("hue.username")
	url := viper.GetString("hue.url")

	return fmt.Sprintf("http://%s/api/%s", url, username)
}
