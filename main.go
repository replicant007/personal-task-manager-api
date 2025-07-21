package main

import (
	"net/http"
)

func main() {
	InitDB()
	defer db.Close()

	http.HandleFunc("/tasks", getTaskHandler)
	http.ListenAndServe(":8080", nil)
}
