package hue

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var client *http.Client = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func getRequest(url string, v interface{}) {

	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(data, &v)

	if err != nil {
		panic(err.Error())
	}
}
