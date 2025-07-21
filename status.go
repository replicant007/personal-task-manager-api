package main

import "encoding/json"

type Status int

const (
	Open = iota
	InProgress
	Completed
	Canceled
)

func (str Status) String() string {
	switch str {
	case Open:
		return "Not Started"

	case InProgress:
		return "In Progress"

	case Completed:
		return "Completed"

	case Canceled:
		return "Cancelled"

	default:
		return "Unknown"
	}
}

func (stat Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(stat.String())
}
