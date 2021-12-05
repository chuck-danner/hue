package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHue_getConfig(t *testing.T) {
	testConfig := config{Name: "Test Config", ApivVrsion: "0.0.0.0"}
	response, _ := json.Marshal(testConfig)
	svr := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, string(response))
	}))
	defer svr.Close()
	tests := []struct {
		name string
		h    Hue
		want config
	}{
		{"Test Config 1", Hue{Key: "Test", URL: svr.URL}, testConfig},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.getConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_getConfigURL(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want string
	}{
		{"Test ConfigURL", Hue{Key: "Test", URL: "http://test.com"}, "http://test.com/api/Test/config"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getConfigURL(); got != tt.want {
				t.Errorf("Hue.getConfigURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
