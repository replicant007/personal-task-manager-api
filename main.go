package main

import (
	"net/http"
)

func main() {
	InitDB()
	defer db.Close()

	http.HandleFunc("/tasks", taskHandler)
	http.HandleFunc("/tasks/", modifyTaskHandler)
	http.ListenAndServe(":8080", nil)
}
