package main

import (
	"net/http"
)

func main() {
	InitDB()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", getTaskHandler)
	mux.HandleFunc("POST /tasks", insertTaskHandler)
	mux.HandleFunc("PUT /tasks/{id}", updateTaskHandler)
	mux.HandleFunc("DELETE /tasks/{id}", deleteTaskHandler)

	http.ListenAndServe(":8080", mux)
}
