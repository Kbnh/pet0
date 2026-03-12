package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/pet0/internal/domain"
	"github.com/Kbnh/pet0/internal/dto"
	"github.com/Kbnh/pet0/pkg/otel/tracer"
	"github.com/Kbnh/pet0/pkg/transaction"
)

func (u *UseCase) CreateTask(ctx context.Context, input dto.CreateTaskInput) (dto.CreateTaskOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase CreateTask")
	defer span.End()

	var output dto.CreateTaskOutput

	task, err := domain.NewTask(input.Title, input.Description)
	if err != nil {
		return output, fmt.Errorf("domain.NewTask: %w", err)
	}

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		err = u.postgres.CreateTask(ctx, task)
		if err != nil {
			return fmt.Errorf("postgres.CreateTask: %w", err)
		}

		return nil
	})
	if err != nil {
		return output, fmt.Errorf("transaction.Wrap: %w", err)
	}

	return dto.CreateTaskOutput{
		ID: task.ID,
	}, nil

}
