package docsapi

import (
	"net/http"

	"gitlab.com/nikolayignatov/personal-task-manager-api/internal/db"
	"gitlab.com/nikolayignatov/personal-task-manager-api/internal/handlers"
)

// @Summary Get all tasks
// @Description Returns a list of all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetTaskHandler(ts db.TaskStorage) http.HandlerFunc {
	return handlers.GetTaskHandler(ts)
}

// @Summary Insert a new task
// @Description Adds a new task to the system
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body TaskInput true "Task to insert"
// @Success 201 {object} models.Task
// @Failure 400 {string} string "Invalid input"
// @Router /tasks [post]
func InsertTaskHandler(ts db.TaskStorage) http.HandlerFunc {
	return handlers.InsertTaskHandler(ts)
}

// @Summary Update a task
// @Description Updates a task by ID
// @Tags tasks
// @Accept json
// @Param id path string true "Task ID"
// @Param task body TaskUpdate true "Updated task"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [put]
func UpdateTaskHandler(ts db.TaskStorage) http.HandlerFunc {
	return handlers.UpdateTaskHandler(ts)
}

// @Summary Delete a task
// @Description Deletes a task by ID
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [delete]
func DeleteTaskHandler(ts db.TaskStorage) http.HandlerFunc {
	return handlers.DeleteTaskHandler(ts)
}

type TaskInput struct {
	Title           string `json:"title" example:"Write Swagger annotations"`
	Description     string `json:"description" example:"Add example values to POST request"`
	CompletedStatus int    `json:"completed_status" example:"2"`
}

type TaskUpdate struct {
	Title           string `json:"title" example:"Updated title...."`
	Description     string `json:"description" example:"Updated description..."`
	CompletedStatus int    `json:"completed_status" example:"1"`
}
