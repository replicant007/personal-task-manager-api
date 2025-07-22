package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		insertTaskHandler(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := getAllTasks()
	log.Printf("Fetched %d tasks\n", len(tasks))
	json.NewEncoder(w).Encode(tasks)
}

func insertTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("The decoded task from the request is %v", task)

	task.Id = uuid.New().String()
	task.CreatedDate = time.Now()

	if task.CompletedStatus < 0 || task.CompletedStatus > 4 {
		http.Error(w, "Status should be between 0 and 4 inclusive", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(task.Title) == "" {
		http.Error(w, "Task title is required", http.StatusBadRequest)
		return
	}

	if err := createTask(task); err != nil {
		http.Error(w, "Failed to insert task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func modifyTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	if id == "" {
		http.Error(w, "Task ID is missing", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		updateTaskHandler(w, r, id)
	case http.MethodDelete:
		deleteTaskHandler(w, r, id)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request, id string) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(newTask.Title) == "" {
		http.Error(w, "Task title is required", http.StatusBadRequest)
		return
	}

	if newTask.CompletedStatus < 0 || newTask.CompletedStatus > 4 {
		http.Error(w, "Status should be between 0 and 4 inclusive", http.StatusBadRequest)
		return
	}

	newTask.CreatedDate = time.Now()

	err := updateTask(id, newTask)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update task", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request, id string) {
	err := deleteTask(id)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
