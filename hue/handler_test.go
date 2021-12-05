package hue

import "testing"

func Test_getRequest(t *testing.T) {
	type args struct {
		url string
		v   interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getRequest(tt.args.url, tt.args.v)
		})
	}
}
