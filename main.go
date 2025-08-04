package main

import (
	"log"
	"net/http"
)

func main() {
	sqliteStorage := InitDB()
	defer sqliteStorage.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", getTaskHandler(sqliteStorage))
	mux.HandleFunc("POST /tasks", insertTaskHandler(sqliteStorage))
	mux.HandleFunc("PUT /tasks/{id}", updateTaskHandler(sqliteStorage))
	mux.HandleFunc("DELETE /tasks/{id}", deleteTaskHandler(sqliteStorage))

	http.ListenAndServe(":8080", mux)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
