package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHue_getLights(t *testing.T) {
	testResult := make(map[string]light)
	testResult["1"] = light{Name: "Test Light", State: state{On: true}}
	response, _ := json.Marshal(testResult)
	svr := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, string(response))
	}))
	defer svr.Close()
	tests := []struct {
		name string
		h    Hue
		want map[string]light
	}{
		{"Test Lights Call", Hue{Key: "Test", URL: svr.URL}, testResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getLights(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.getLights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_getLightsURL(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want string
	}{
		{"Test Group URL", Hue{Key: "Test", URL: "http://test.com"}, "http://test.com/api/Test/lights"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getLightsURL(); got != tt.want {
				t.Errorf("Hue.getLightsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
