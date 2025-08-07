package models

import (
	"time"
)

// Task represents a to-do item.
// swagger:model Task
type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// CompletedStatus is the current status of the task
	// enum: [0,1,2,3]
	// x-enum-varnames: [Not Started, In Progress, Completed, Cancelled]
	CompletedStatus Status `json:"completed_status"`
	// swagger:strfmt date-time
	CreatedDate time.Time `json:"created_date"`
}
