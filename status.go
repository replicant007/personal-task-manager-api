package main

import "encoding/json"

type Status int

const (
	Open = iota
	InProgress
	Completed
	Canceled
)

func (st Status) String() string {
	switch st {
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

func (st Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.String())
}
