package service

import "GabrielChaves1/task-tracker/internal/ports"

type TaskService struct {
	storage ports.Storage
}

func NewTaskService(storage ports.Storage) *TaskService {
	return &TaskService{
		storage: storage,
	}
}

func (a *TaskService) Add(text string) (string, error) {
	if text == "" {
		return "", nil
	}

	return a.storage.Add(text)
}

func (a *TaskService) Remove(id int) (string, error) {
	if id == 0 {
		return "", nil
	}

	return a.storage.Remove(id)
}

func (a *TaskService) Update(id int, text string) (string, error) {
	if id == 0 || text == "" {
		return "", nil
	}

	return a.storage.Update(id, text)
}

func (a *TaskService) List(filter ...string) (string, error) {
	var f string
	if len(filter) > 0 {
		f = filter[0]
	}
	return a.storage.List(f)
}

func (a *TaskService) IsValidStatus(status string) (string, error) {
	if status == "" {
		return "", nil
	}

	if status != "pending" && status != "completed" && status != "in-progress" {
		return "incorrect status", nil
	}

	return status, nil
}

func (a *TaskService) UpdateStatus(id int, status string) (string, error) {
	if id == 0 || status == "" {
		return "", nil
	}

	return a.storage.UpdateStatus(id, status)
}