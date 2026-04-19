package main

type Event struct {
	EventType string `json:"event_type"`
	Service   string `json:"service"`
	Status    string `json:"status"`
}
