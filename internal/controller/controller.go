package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/avran02/verba-group/internal/dto"
	"github.com/avran02/verba-group/internal/mapper"
	"github.com/avran02/verba-group/internal/service"
	"github.com/go-chi/chi/v5"
)

type Controller interface {
	CreateTaskHandler(w http.ResponseWriter, r *http.Request)
	ListTasksHandler(w http.ResponseWriter, r *http.Request)
	GetTaskHandler(w http.ResponseWriter, r *http.Request)
	UpdateTaskHandler(w http.ResponseWriter, r *http.Request)
	DeleteTaskHandler(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	s service.Service
}

func New(s service.Service) Controller {
	return &controller{
		s: s,
	}
}

func (c *controller) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("controller.CreateTaskHandler - Creating new task")

	var req dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	taskID, err := c.s.CreateTask(req)
	if err != nil {
		slog.Error("Failed to create task", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := dto.CreateTaskResponse{
		ID:          taskID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		CreatedAt:   req.DueDate,
		UpdatedAt:   req.DueDate,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (c *controller) ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("controller.ListTasksHandler - Listing all tasks")

	tasks, err := c.s.ListTasks()
	if err != nil {
		slog.Error("Failed to list tasks", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (c *controller) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("controller.GetTaskHandler - Fetching task")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Invalid task ID", "id", idStr, "error", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := c.s.GetTask(id)
	if err != nil {
		slog.Error("Failed to get task", "id", id, "error", err)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (c *controller) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("controller.UpdateTaskHandler - Updating task")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Invalid task ID", "id", idStr, "error", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task, err := c.s.UpdateTask(id, req)
	if err != nil {
		slog.Error("Failed to update task", "id", id, "error", err)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	resp := mapper.ToUpdateTaskResponse(task)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("controller.DeleteTaskHandler - Deleting task")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("Invalid task ID", "id", idStr, "error", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := c.s.DeleteTask(id); err != nil {
		slog.Error("Failed to delete task", "id", id, "error", err)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
