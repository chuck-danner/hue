package hue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHue_getCheck(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "{\"name\": \"name\"}")
	}))
	defer svr.Close()
	tests := []struct {
		name string
		h    Hue
		want check
	}{
		{"base test", Hue{Key: "Test", URL: svr.URL}, check{Name: "name"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getCheck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.getCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_getCheckURL(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want string
	}{
		{"Check URL", Hue{Key: "Test", URL: "http://test.com"}, "http://test.com/api/0/config"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getCheckURL(); got != tt.want {
				t.Errorf("Hue.getCheckURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
