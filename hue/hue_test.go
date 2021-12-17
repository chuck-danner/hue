package hue

import (
	"reflect"
	"testing"
)

func TestHue_Lights(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want map[string]light
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Lights(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.Lights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_Config(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Config(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.Config() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_Groups(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want map[string]group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Groups(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.Groups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_Schedules(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want map[string]schedule
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Schedules(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.Schedules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_Check(t *testing.T) {
	tests := []struct {
		name string
		h    Hue
		want check
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Check(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hue.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetStatus(t *testing.T) {
	type args struct {
		light  string
		status string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//SetStatus(tt.args.light, tt.args.status)
		})
	}
}

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Setup(); got != tt.want {
				t.Errorf("Setup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_getBaseURL(t *testing.T) {
	tests := []struct {
		name string
		hue  Hue
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hue.getBaseURL(); got != tt.want {
				t.Errorf("Hue.getBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHue_SetStatus(t *testing.T) {
	type args struct {
		id     string
		status string
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
			tt.h.SetStatus(tt.args.id, tt.args.status)
		})
	}
}

func TestHue_getV2BaseURL(t *testing.T) {
	tests := []struct {
		name string
		hue  Hue
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hue.getV2BaseURL(); got != tt.want {
				t.Errorf("Hue.getV2BaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
