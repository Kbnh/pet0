package dto

import "github.com/google/uuid"

type CreateTaskOutput struct {
	ID uuid.UUID `json:"id"`
}

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
