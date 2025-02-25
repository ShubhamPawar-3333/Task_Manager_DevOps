package db

import (
	"Task-Manager-Backend/internal/models"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitSQLite() error {
	var err error
	DB, err = sql.Open("sqlite", "./tasks.db")
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN,
			created_at DATETIME
		)
	`)
	if err != nil {
		return err
	}
	log.Println("Connected to SQLite")
	return nil
}

func CloseSQLite() {
	DB.Close()
}

func GetAllTasks() ([]models.Task, error) {
	rows, err := DB.Query("SELECT id, title, completed, created_at FROM tasks")
	if err != nil {
		return []models.Task{}, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
		if err != nil {
			return []models.Task{}, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	err := DB.QueryRow("SELECT id, title, completed, created_at FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt)
	if err == sql.ErrNoRows {
		return models.Task{}, fmt.Errorf("task not found")
	}
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func InsertTask(t models.Task) error {
	_, err := DB.Exec(
		"INSERT INTO tasks (id, title, completed, created_at) VALUES (?, ?, ?, ?)",
		t.ID, t.Title, t.Completed, t.CreatedAt,
	)
	return err
}

func UpdateTask(id string, t models.Task) error {
	_, err := DB.Exec(
		"UPDATE tasks SET title = ?, completed = ? WHERE id = ?",
		t.Title, t.Completed, id,
	)
	return err
}

func DeleteTask(id string) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
