package adapters

import (
	"GabrielChaves1/task-tracker/internal/domain"
	"GabrielChaves1/task-tracker/internal/ports"
	"GabrielChaves1/task-tracker/internal/utils"
	"fmt"
	"time"
)

type JSONStorage struct{}

func NewJSONStorage() ports.Storage {
	return &JSONStorage{}
}

func (j *JSONStorage) Add(text string) (string, error) {
	tasks, err := utils.ReadFile[domain.Task]("tasks.json")

	if err != nil {
		return "", err
	}

	task := domain.Task{
		ID:       len(tasks) + 1,
		Text:     text,
		Datetime: time.Now().Format(time.DateTime),
		Status:   "pending",
	}

	tasks = append(tasks, task)

	err = utils.WriteFile("tasks.json", tasks)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Task added successfully (ID: %d)", task.ID), nil
}

func (j *JSONStorage) Remove(id int) (string, error) {
	tasks, err := utils.ReadFile[domain.Task]("tasks.json")

	if err != nil {
		return "", err
	}

	if id <= 0 || id > len(tasks) {
		return "", fmt.Errorf("invalid task ID: %d", id)
	}

	tasks = append(tasks[:id-1], tasks[id:]...)

	err = utils.WriteFile("tasks.json", tasks)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Task removed successfully (ID: %d)", id), nil
}

func (j *JSONStorage) Update(id int, new string) (string, error) {
	tasks, err := utils.ReadFile[domain.Task]("tasks.json")

	if err != nil {
		return "", err
	}

	if id <= 0 || id > len(tasks) {
		return "", fmt.Errorf("invalid task ID: %d", id)
	}

	tasks[id-1].Text = new

	err = utils.WriteFile("tasks.json", tasks)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Task updated successfully (ID: %d)", id), nil
}

func (j *JSONStorage) List(filter string) (string, error) {
	tasks, err := utils.ReadFile[domain.Task]("tasks.json")

	if err != nil {
		return "", err
	}

	var result string

	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			result += fmt.Sprintf("[%d] %s (%s) - %s\n", task.ID, task.Text, task.Status, task.Datetime)
		}
	}

	return result, nil
}

func (j *JSONStorage) UpdateStatus(id int, status string) (string, error) {
	tasks, err := utils.ReadFile[domain.Task]("tasks.json")

	if err != nil {
		return "", err
	}

	if id <= 0 || id > len(tasks) {
		return "", fmt.Errorf("invalid task ID: %d", id)
	}

	tasks[id-1].Status = status

	err = utils.WriteFile("tasks.json", tasks)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Task status updated successfully to [%s] (ID: %d)", status, id), nil
}
