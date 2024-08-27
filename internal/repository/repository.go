package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/avran02/verba-group/config"
	"github.com/avran02/verba-group/internal/models"
)

type Repository interface {
	CreateTask(models.Task) (int, error)
	ListTasks() ([]models.Task, error)
	GetTask(id int) (models.Task, error)
	UpdateTask(models.Task) (models.Task, error)
	DeleteTask(id int) error

	Close() error
}

type repository struct {
	db *sql.DB
}

func (r *repository) CreateTask(task models.Task) (int, error) {
	slog.Info("repository.CreateTask", "task", task)

	query := `
		INSERT INTO tasks (title, description, due_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt

	var taskID int
	err := r.db.QueryRow(query, task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt).Scan(&taskID)
	if err != nil {
		slog.Error("failed to create task", "error", err)
		return 0, fmt.Errorf("failed to create task: %w", err)
	}

	return taskID, nil
}

func (r *repository) ListTasks() ([]models.Task, error) {
	slog.Info("repository.ListTasks")

	query := `SELECT id, title, description, due_date, created_at, updated_at FROM tasks`
	rows, err := r.db.Query(query)
	if err != nil {
		slog.Error("failed to list tasks", "error", err)
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			slog.Error("failed to scan task", "error", err)
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *repository) GetTask(id int) (models.Task, error) {
	slog.Info("repository.GetTask", "id", id)

	query := `SELECT id, title, description, due_date, created_at, updated_at FROM tasks WHERE id = $1`
	var task models.Task
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, ErrNotFound
		}
		slog.Error("failed to get task", "error", err)
		return task, err
	}

	return task, nil
}

func (r *repository) UpdateTask(task models.Task) (models.Task, error) {
	slog.Info("repository.UpdateTask", "task", task)

	query := `
		UPDATE tasks
		SET title = $1, description = $2, due_date = $3, updated_at = $4
		WHERE id = $5`

	task.UpdatedAt = time.Now()

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.DueDate,
		task.UpdatedAt,
		task.ID,
	)
	if err != nil {
		slog.Error("failed to update task", "error", err)
		return models.Task{}, fmt.Errorf("failed to update task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("failed to retrieve affected rows", "error", err)
		return models.Task{}, fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return models.Task{}, ErrNotFound
	}

	return task, nil
}

func (r *repository) DeleteTask(id int) error {
	slog.Info("repository.DeleteTask", "id", id)

	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		slog.Error("failed to delete task", "error", err)
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("failed to retrieve affected rows", "error", err)
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) Close() error {
	slog.Info("repository.Close")
	return r.db.Close()
}

func New(conf config.Postgres) Repository {
	slog.Info("repository.New")
	db, err := sql.Open("postgres", getDsn(conf))
	if err != nil {
		slog.Error("can't open db conn", "error", err)
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		slog.Error("can't ping database", "error", err)
		os.Exit(1)
	}

	slog.Info("db connected")
	return &repository{
		db: db,
	}
}

func getDsn(conf config.Postgres) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Database,
	)
}
