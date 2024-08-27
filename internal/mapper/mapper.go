package mapper

import (
	"github.com/avran02/verba-group/internal/dto"
	"github.com/avran02/verba-group/internal/models"
)

func FromCreateTaskRequest(req dto.CreateTaskRequest) models.Task {
	return models.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
	}
}

func FromUpdateTaskRequest(req dto.UpdateTaskRequest) models.Task {
	return models.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
	}
}

func ToUpdateTaskResponse(task models.Task) dto.UpdateTaskResponse {
	return dto.UpdateTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		UpdatedAt:   task.UpdatedAt,
		CreatedAt:   task.CreatedAt,
	}
}
