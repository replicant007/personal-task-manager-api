package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB

func InitDB() {
	createDB()

	checkConnection()
	createTables()

	seedStatuses()
	seedTasks()
}

func createDB() {
	var err error
	db, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
}

func checkConnection() {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Successfully connected to SQLite database.")
}

func createTables() {
	taskTable := `
	CREATE TABLE IF NOT EXISTS Tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		completed_status INTEGER NOT NULL,
		created_date DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(taskTable); err != nil {
		log.Fatal(err)
	}

	statusTable := `
	CREATE TABLE IF NOT EXISTS Status (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);`

	if _, err := db.Exec(statusTable); err != nil {
		log.Fatal(err)
	}
}

func seedStatuses() {

	if countStatusRecords() != 0 {
		return
	}

	statuses := []struct {
		id   int
		name string
	}{
		{0, "Not Started"},
		{1, "In Progress"},
		{2, "Completed"},
		{3, "Cancelled"},
	}

	stmtStatuses, err := db.Prepare("INSERT INTO Status (id, name) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Failed to prepare insert statement:", err)
	}
	defer stmtStatuses.Close()

	for _, s := range statuses {
		_, err := stmtStatuses.Exec(s.id, s.name)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func countStatusRecords() int {
	row := db.QueryRow("SELECT COUNT(*) FROM Status")
	count := 0
	err := row.Scan(&count)
	if err != nil {
		log.Fatal("Counting Status records", err)
	}
	return count
}

func seedTasks() {
	stmtTasks, err := db.Prepare(`
		INSERT INTO Tasks (id, title, description, completed_status, created_date) 
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmtTasks.Close()

	if !taskExists("task-001") {
		_, err = stmtTasks.Exec("task-001", "Learn Go", "Complete the tutorial on structs and interfaces", 2, time.Date(2025, 7, 14, 9, 0, 0, 0, time.UTC))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task seeded: task-001")
	}

	if !taskExists("task-002") {
		_, err = stmtTasks.Exec("task-002", "Prepare project", "Create a new Golang project in Gitlab", 1, time.Date(2025, 7, 13, 9, 0, 0, 0, time.UTC))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task seeded: task-002")
	}
}

func taskExists(id string) bool {
	row := db.QueryRow("SELECT COUNT(*) FROM Tasks WHERE id = ?", id)
	var count int
	if err := row.Scan(&count); err != nil {
		log.Fatal("Checking if task exists: ", err)
	}
	return count != 0
}

func getAllTasks() []Task {
	rows, err := db.Query("SELECT * FROM Tasks")
	if err != nil {
		log.Fatal("Failed to query data:", err)
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		var statusInt int

		err := rows.Scan(&t.Id, &t.Title, &t.Description, &statusInt, &t.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}

		t.CompletedStatus = Status(statusInt)
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return []Task{}
	}

	return tasks
}

func createTask(task Task) error {
	stmt, err := db.Prepare(`
		INSERT INTO Tasks (id, title, description, completed_status, created_date)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Id, task.Title, task.Description, int(task.CompletedStatus), task.CreatedDate)
	return err
}
