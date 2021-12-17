package hue

import (
	"io"
	"testing"
)

func TestHue_getRequest(t *testing.T) {
	type args struct {
		url string
		v   interface{}
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
			tt.h.getRequest(tt.args.url, tt.args.v)
		})
	}
}

func TestHue_putRequest(t *testing.T) {
	type args struct {
		url  string
		body io.Reader
		v    interface{}
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
			tt.h.putRequest(tt.args.url, tt.args.body, tt.args.v)
		})
	}
}

func TestHue_request(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
		v      interface{}
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
			tt.h.request(tt.args.method, tt.args.url, tt.args.body, tt.args.v)
		})
	}
}
