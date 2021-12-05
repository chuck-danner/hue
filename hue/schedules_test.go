package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHue_listSchedules(t *testing.T) {
	testResult := make(map[string]schedule)
	testResult["1"] = schedule{Name: "Test Schedule"}
	response, _ := json.Marshal(testResult)
	svr := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, string(response))
	}))
	tests := []struct {
		name string
		h    Hue
		want map[string]schedule
	}{
		{"Test Schedule Call", Hue{Key: "Test", URL: svr.URL}, testResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.listSchedules(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.listSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_getSchedulesURL(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want string
	}{
		{"Test Group URL", Hue{Key: "Test", URL: "http://test.com"}, "http://test.com/api/Test/schedules"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getSchedulesURL(); got != tt.want {
				t.Errorf("Hue.getSchedulesURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
