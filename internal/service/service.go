package service

import (
	"time"

	"github.com/avran02/verba-group/internal/dto"
	"github.com/avran02/verba-group/internal/mapper"
	"github.com/avran02/verba-group/internal/models"
	"github.com/avran02/verba-group/internal/repository"
)

type Service interface {
	CreateTask(dto.CreateTaskRequest) (int, error)
	ListTasks() ([]models.Task, error)
	GetTask(id int) (models.Task, error)
	UpdateTask(id int, task dto.UpdateTaskRequest) (models.Task, error)
	DeleteTask(id int) error
}

type service struct {
	repository repository.Repository
}

func (s *service) CreateTask(task dto.CreateTaskRequest) (int, error) {
	model := mapper.FromCreateTaskRequest(task)
	return s.repository.CreateTask(model)
}

func (s *service) ListTasks() ([]models.Task, error) {
	return s.repository.ListTasks()
}

func (s *service) GetTask(id int) (models.Task, error) {
	return s.repository.GetTask(id)
}

func (s *service) UpdateTask(id int, task dto.UpdateTaskRequest) (models.Task, error) {
	model := mapper.FromUpdateTaskRequest(task)
	model.ID = id
	model.UpdatedAt = time.Now()
	return s.repository.UpdateTask(model)
}

func (s *service) DeleteTask(id int) error {
	return s.repository.DeleteTask(id)
}

func New(repository repository.Repository) Service {
	return &service{repository: repository}
}
