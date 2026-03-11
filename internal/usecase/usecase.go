package usecase

import (
	"context"

	"github.com/Kbnh/pet0/internal/domain"
	"github.com/google/uuid"
)

type Postgres interface {
	CreateTask(ctx context.Context, task domain.Task) error
	GetTask(ctx context.Context, taskID uuid.UUID) (domain.Task, error)
	GetTasks(ctx context.Context) ([]domain.Task, error)
	UpdateTask(ctx context.Context, task domain.Task) error
	DeleteTask(ctx context.Context, taskID uuid.UUID) error
}

// type Redis interface {}

// type Kafka interface {}

type UseCase struct {
	postgres Postgres
	// redis Redis
	// kafka Kafka
}

func New(postgres Postgres) *UseCase {
	return &UseCase{
		postgres: postgres,
		// redis: redis,
		// kafka: kafka,
	}
}
