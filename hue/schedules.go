package hue

import "fmt"

type schedule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func listSchedules() map[string]schedule {
	schedules := make(map[string]schedule)
	getRequest(getSchedulesURL(), &schedules)
	return schedules
}

func getSchedulesURL() string {
	return fmt.Sprintf("%s/%s", getBaseURL(), "schedules")
}
