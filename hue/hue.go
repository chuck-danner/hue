package hue

import (
	"fmt"
)

type Hue struct {
	Key string
	URL string
	ID  string
}

func (h Hue) Lights() lightResponseV2 {
	return h.getLights()
}

func (h Hue) Config() config {
	return h.getConfig()
}

func (h Hue) Groups() map[string]group {
	return h.getGroups()
}

func (h Hue) Schedules() map[string]schedule {
	return h.listSchedules()
}

func (h Hue) Check() check {
	return h.getCheck()
}

func (h Hue) SetStatus(id string, status string) {
	set := false
	if status == "on" {
		set = true
	}
	light := light{On: state{On: set}}

	h.setLight(id, light)

}

func Setup() string {
	return discoverEndpoint()
}

func (hue Hue) getBaseURL() string {
	return fmt.Sprintf("%s/api/%s", hue.URL, hue.Key)
}

func (hue Hue) getV2BaseURL() string {
	return fmt.Sprintf("%s/clip/v2/resource", hue.URL)
}
