package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type Hue struct {
	Username string
	URL      string
}

func DefaultHue() *Hue {
	return &Hue{
		Username: viper.GetString("hue.username"),
		URL:      viper.GetString("hue.url")}
}

func (h Hue) Lights() map[string]light {
	return h.getLights()
}

func (h Hue) Config() config {
	return h.getConfig()
}

func (h Hue) Groups() map[string]group {
	return h.getGroups()
}

func (h Hue) Schedules() map[string]schedule {
	return h.listSchedules()
}

func (h Hue) Check() check {
	return h.getCheck()
}

//TODO: Needs fixed and moved
func SetStatus(light string, status string) {
	statusURL := fmt.Sprintf("%s/lights/%s/state", "", light)
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

func Setup() string {
	return discoverEndpoint()
}

func (hue Hue) getBaseURL() string {
	username := hue.Username
	url := hue.URL

	return fmt.Sprintf("http://%s/api/%s", url, username)
}
