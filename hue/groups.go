package hue

import (
	"fmt"
)

type group struct {
	Name   string   `json:"name"`
	Lights []string `json:"lights"`
	Type   string   `json:"type"`
}

func listGroups() map[string]group {

	groups := make(map[string]group)
	getRequest(getGroupsURL(), &groups)
	return groups
}

func getGroupsURL() string {
	return fmt.Sprintf("%s/%s", getBaseURL(), "groups")
}
