package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func GetLights() map[string]light {
	return listLights()
}

func GetConfig() config {
	return listConfig()
}

func GetGroups() map[string]group {
	return listGroups()
}

func GetSchedules() map[string]schedule {
	return listSchedules()
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
