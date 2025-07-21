package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := getAllTasks()
	log.Printf("Fetched %d tasks\n", len(tasks))
	json.NewEncoder(w).Encode(tasks)
}
