package status

import "time"

type Health struct {
	Status      string      `json:"status"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	Uptime      Uptime      `json:"uptime"`
	Details     []Component `json:"details"`
}

type Uptime struct {
	From time.Time `json:"date"`
}

type Component struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Status string  `json:"status"`
	Links  []Links `json:"links"`
}

type Links struct {
	Self string `json:"self"`
}
