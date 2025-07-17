package main

import (
	"time"
)

type Task struct {
	Id              string
	Title           string
	Description     string
	CompletedStatus Status
	CreatedDate     time.Time
}

var tasks = []Task{
	{"task-001", "Learn Go", "Complete the tutorial on structs and interfaces", InProgress, time.Date(2025, 7, 14, 9, 0, 0, 0, time.UTC)},
	{"task-002", "Prepare project", "Create a new Golang project in Gitlab", Open, time.Date(2025, 7, 13, 9, 0, 0, 0, time.UTC)},
}
