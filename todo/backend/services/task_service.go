package services

import (
		"github.com/Matama091/go_exercise/todo/backend/models"
		"github.com/Matama091/go_exercise/todo/backend/repositories"
)

func GetAllTasks() ([]models.Task, error) {
    return repositories.GetAllTasks()
}

func CreateTask(task *models.Task) error {
    return repositories.CreateTask(task)
}

func UpdateTask(id string, task *models.Task) error {
    return repositories.UpdateTask(id, task)
}

func DeleteTask(id string) error {
    return repositories.DeleteTask(id)
}
