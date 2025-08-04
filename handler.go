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

func getTaskHandler(ts TaskStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tasks, _ := ts.GetAllTasks()
		log.Printf("Fetched %d tasks\n", len(tasks))
		json.NewEncoder(w).Encode(tasks)
	}
}

func insertTaskHandler(ts TaskStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var task Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Task sent with the request: %v", task)

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

		if err := ts.CreateTask(task); err != nil {
			http.Error(w, "Failed to insert task", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

func updateTaskHandler(ts TaskStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
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

		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Task ID is missing", http.StatusBadRequest)
			return
		}

		err := ts.UpdateTask(id, newTask)
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
}

func deleteTaskHandler(ts TaskStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "Task ID is missing", http.StatusBadRequest)
			return
		}

		err := ts.DeleteTask(id)
		if err != nil {
			if strings.Contains(err.Error(), "does not exist") {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			}
			return
		}
		log.Printf("Taks with id %s deleted\n", id)
		w.WriteHeader(http.StatusNoContent)
	}
}
