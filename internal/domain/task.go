package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title" validate:"required,min=3,max=64"`
	Description *string    `json:"description,omitempty"`
	Status      Status     `json:"status" validate:"required"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func NewTask(title, description string) (Task, error) {
	t := Task{
		ID:          uuid.New(),
		Title:       title,
		Description: &description,
		Status:      StatusNew,
	}

	if err := t.Validate(); err != nil {
		return Task{}, fmt.Errorf("t.Validate: %w", err)
	}

	return t, nil
}

func (t Task) Validate() error {
	err := validate.Struct(t)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}
	return nil
}
