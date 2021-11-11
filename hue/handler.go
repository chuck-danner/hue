package hue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getRequest(url string, v interface{}) {
	resp, err := http.Get(url)

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
