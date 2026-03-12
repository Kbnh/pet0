package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/pet0/internal/domain"
	"github.com/Kbnh/pet0/pkg/otel/tracer"
	"github.com/Kbnh/pet0/pkg/transaction"
)

func (p *Postgres) CreateTask(ctx context.Context, task domain.Task) error {
	ctx, span := tracer.Start(ctx, "adapter postgres CreateTask")
	defer span.End()

	const sql = `INSERT INTO tasks (id, title, description, status)
				VALUES ($1, $2, $3, $4)`

	args := []any{
		task.ID,
		task.Title,
		task.Description,
		task.Status,
	}

	txOrPool := transaction.TryExtractTX(ctx)

	_, err := txOrPool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
