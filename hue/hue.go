package hue

import (
	"bytes"
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

func SetStatus(light string, status string) {
	statusURL := fmt.Sprintf("%s/lights/%s/state", getBaseURL(), light)
	fmt.Println(statusURL)

	state := false
	if status == "true" {
		state = true
	}
	requestBody, _ := json.Marshal(map[string]bool{
		"on": state})
	fmt.Println(string(requestBody))
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, statusURL, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp)

}

func getBaseURL() string {
	username := viper.GetString("hue.username")
	url := viper.GetString("hue.url")

	return fmt.Sprintf("http://%s/api/%s", url, username)
}
