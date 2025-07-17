package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println(tasks)

	http.HandleFunc("/tasks", getTaskHandler)
	http.ListenAndServe(":8080", nil)
}
