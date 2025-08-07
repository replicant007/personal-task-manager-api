package main

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "gitlab.com/nikolayignatov/personal-task-manager-api/docs" // This is important to register the docs
	"gitlab.com/nikolayignatov/personal-task-manager-api/docsapi"
	"gitlab.com/nikolayignatov/personal-task-manager-api/internal/db"
)

func main() {
	sqliteStorage := db.InitDB()
	defer sqliteStorage.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", docsapi.GetTaskHandler(sqliteStorage))
	mux.HandleFunc("POST /tasks", docsapi.InsertTaskHandler(sqliteStorage))
	mux.HandleFunc("PUT /tasks/{id}", docsapi.UpdateTaskHandler(sqliteStorage))
	mux.HandleFunc("DELETE /tasks/{id}", docsapi.DeleteTaskHandler(sqliteStorage))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", mux)
}
