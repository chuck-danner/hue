package hue

import (
	"fmt"
)

type group struct {
	Name   string   `json:"name"`
	Lights []string `json:"lights"`
	Type   string   `json:"type"`
}

func (h Hue) getGroups() map[string]group {

	groups := make(map[string]group)
	getRequest(h.getGroupsURL(), &groups)
	return groups
}

func (h Hue) getGroupsURL() string {
	return fmt.Sprintf("%s/%s", h.getBaseURL(), "groups")
}
