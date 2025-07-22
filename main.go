package main

import (
	"net/http"
)

func main() {
	InitDB()
	defer db.Close()

	http.HandleFunc("/tasks", taskHandler)
	http.HandleFunc("/tasks/", updateTaskHandler)
	http.ListenAndServe(":8080", nil)
}
