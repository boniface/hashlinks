package domain

import "time"

type LogEvent struct {
	SiteId    string    `json:"site_id"`
	Id        string    `json:"id"`
	EventName string    `json:"event_name"`
	EventType string    `json:"event_type"`
	Message   string    `json:"message"`
	Date      time.Time `json:"date"`
}

type LogEvents [] LogEvent
