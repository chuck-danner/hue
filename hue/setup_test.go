package hue

import "testing"

func Test_discoverMdns(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discoverMdns()
		})
	}
}

func Test_discoverEndpoint(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := discoverEndpoint(); got != tt.want {
				t.Errorf("discoverEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
