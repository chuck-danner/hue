package hue

import "fmt"

type schedule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h Hue) listSchedules() map[string]schedule {
	schedules := make(map[string]schedule)
	getRequest(h.getSchedulesURL(), &schedules)
	return schedules
}

func (h Hue) getSchedulesURL() string {
	return fmt.Sprintf("%s/%s", h.getBaseURL(), "schedules")
}
