package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHue_getGroups(t *testing.T) {
	testResult := make(map[string]group)
	testResult["1"] = group{Name: "Group1", Lights: make([]string, 3), Type: "Group"}
	response, _ := json.Marshal(testResult)
	svr := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, string(response))
	}))
	defer svr.Close()
	tests := []struct {
		name string
		h    Hue
		want map[string]group
	}{
		{"Test Get Groupe", Hue{Key: "Test", URL: svr.URL}, testResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getGroups(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.getGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_getGroupsURL(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want string
	}{
		{"Test Group URL", Hue{Key: "Test", URL: "http://test.com"}, "http://test.com/api/Test/groups"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getGroupsURL(); got != tt.want {
				t.Errorf("Hue.getGroupsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
