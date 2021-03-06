package handler

import (
	"context"

	"github.com/homma509/nrece/domain/repository"
	"github.com/homma509/nrece/usecase"
)

// FileHandler ...
type FileHandler struct {
	u *usecase.File
}

// NewFileHandler ...
func NewFileHandler(repo repository.FileRepository) (*FileHandler, error) {
	return &FileHandler{
		usecase.NewFile(repo),
	}, nil
}

// CopyFile ...
func (h *FileHandler) CopyFile(ctx context.Context, s3URL string) error {
	if err := h.u.Copy(ctx, s3URL); err != nil {
		return err
	}
	return nil
}
