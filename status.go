package main

type Status int

const (
	Open = iota
	InProgress
	Completed
	Canceled
)

func (s Status) String() string {
	switch s {
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
