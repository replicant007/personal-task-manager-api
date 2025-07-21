package main

import (
	"time"
)

type Task struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	CompletedStatus Status    `json:"completed_status"`
	CreatedDate     time.Time `json:"created_date"`
}
