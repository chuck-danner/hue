package hue

import (
	"testing"
)

func TestHue_getLightsURL(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want string
	}{
		{"Test Group URL", Hue{Key: "Test", URL: "http://test.com"}, "http://test.com/clip/v2/resource/light"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getLightsURL(); got != tt.want {
				t.Errorf("Hue.getLightsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_setLight(t *testing.T) {
	type args struct {
		id    string
		light light
	}
	tests := []struct {
		name string
		h    Hue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.setLight(tt.args.id, tt.args.light)
		})
	}
}
