package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"gitlab.com/nikolayignatov/personal-task-manager-api/internal/models"
)

type TaskStorage interface {
	GetAllTasks() ([]models.Task, error)
	CreateTask(task models.Task) error
	UpdateTask(id string, task models.Task) error
	DeleteTask(id string) error
	CloseDB()
}

type SQLiteTaskStorage struct {
	db *sql.DB
}

func InitDB() TaskStorage {
	sqlTs := createDB()
	sqlTs.checkConnection()
	sqlTs.createTables()

	sqlTs.seedStatuses()
	sqlTs.seedTasks()

	return sqlTs
}

func createDB() *SQLiteTaskStorage {
	db, err := sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	return &SQLiteTaskStorage{db: db}
}

func (sqlTs *SQLiteTaskStorage) checkConnection() {
	err := sqlTs.db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Successfully connected to SQLite database.")
}

func (sqlTs *SQLiteTaskStorage) createTables() {
	taskTable := `
	CREATE TABLE IF NOT EXISTS Tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		completed_status INTEGER NOT NULL,
		created_date DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := sqlTs.db.Exec(taskTable); err != nil {
		log.Fatal(err)
	}

	statusTable := `
	CREATE TABLE IF NOT EXISTS Status (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);`

	if _, err := sqlTs.db.Exec(statusTable); err != nil {
		log.Fatal(err)
	}
}

func (sqlTs *SQLiteTaskStorage) seedStatuses() {

	if sqlTs.countStatusRecords() != 0 {
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

	stmtStatuses, err := sqlTs.db.Prepare("INSERT INTO Status (id, name) VALUES (?, ?)")
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

func (sqlTs *SQLiteTaskStorage) countStatusRecords() int {
	row := sqlTs.db.QueryRow("SELECT COUNT(*) FROM Status")
	count := 0
	err := row.Scan(&count)
	if err != nil {
		log.Fatal("Counting Status records", err)
	}
	return count
}

func (sqlTs *SQLiteTaskStorage) seedTasks() {
	stmtTasks, err := sqlTs.db.Prepare(`
		INSERT INTO Tasks (id, title, description, completed_status, created_date) 
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmtTasks.Close()

	if !sqlTs.taskExists("task-001") {
		_, err = stmtTasks.Exec("task-001", "Learn Go", "Complete the tutorial on structs and interfaces", 2, time.Date(2025, 7, 14, 9, 0, 0, 0, time.UTC))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task seeded: task-001")
	}

	if !sqlTs.taskExists("task-002") {
		_, err = stmtTasks.Exec("task-002", "Prepare project", "Create a new Golang project in Gitlab", 1, time.Date(2025, 7, 13, 9, 0, 0, 0, time.UTC))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task seeded: task-002")
	}
}

func (sqlTs *SQLiteTaskStorage) taskExists(id string) bool {
	row := sqlTs.db.QueryRow("SELECT COUNT(*) FROM Tasks WHERE id = ?", id)
	var count int
	if err := row.Scan(&count); err != nil {
		log.Printf("Checking if task exists: %s", err)
	}
	return count != 0
}

func (sqlTs *SQLiteTaskStorage) GetAllTasks() ([]models.Task, error) {
	rows, err := sqlTs.db.Query("SELECT * FROM Tasks")
	if err != nil {
		log.Printf("Failed to query data: %s", err)
		return []models.Task{}, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		var statusInt int

		err := rows.Scan(&t.Id, &t.Title, &t.Description, &statusInt, &t.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}

		t.CompletedStatus = models.Status(statusInt)
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return []models.Task{}, err
	}

	return tasks, nil
}

func (sqlTs *SQLiteTaskStorage) CreateTask(task models.Task) error {
	stmt, err := sqlTs.db.Prepare(`
		INSERT INTO Tasks (id, title, description, completed_status, created_date)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("create task failed: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Id, task.Title, task.Description, int(task.CompletedStatus), task.CreatedDate)
	if err != nil {
		return fmt.Errorf("create task failed: %w", err)
	}
	return err
}

func (sqlTs *SQLiteTaskStorage) UpdateTask(id string, task models.Task) error {
	if !sqlTs.taskExists(id) {
		return fmt.Errorf("task with ID %q does not exist", id)
	}

	stmt, err := sqlTs.db.Prepare(`
		UPDATE Tasks 
		SET title = ?, description = ?, completed_status = ?, created_date = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Description, int(task.CompletedStatus), task.CreatedDate, id)
	return err
}

func (sqlTs *SQLiteTaskStorage) DeleteTask(id string) error {
	if !sqlTs.taskExists(id) {
		return fmt.Errorf("task with id %q does not exist", id)
	}

	stmt, err := sqlTs.db.Prepare("DELETE FROM Tasks WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (sqlTs *SQLiteTaskStorage) CloseDB() {
	sqlTs.db.Close()
}
