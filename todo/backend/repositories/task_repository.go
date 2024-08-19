package repositories

import (
	"github.com/Matama091/go_exercise/todo/backend/config"
	"github.com/Matama091/go_exercise/todo/backend/models"
)

func GetAllTasks() ([]models.Task, error) {
	rows, err := config.DB.Query("SELECT id, title, description, due_date, status, created_at, updated_at FROM Tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTask(task *models.Task) error {
	query := "INSERT INTO Tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, task.Title, task.Description, task.DueDate, task.Status)
	return err
}

func UpdateTask(id string, task *models.Task) error {
	query := "UPDATE Tasks SET title=?, description=?, due_date=?, status=? WHERE id=?"
	_, err := config.DB.Exec(query, task.Title, task.Description, task.DueDate, task.Status, id)
	return err
}

func DeleteTask(id string) error {
	query := "DELETE FROM Tasks WHERE id=?"
	_, err := config.DB.Exec(query, id)
	return err
}
