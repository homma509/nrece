package handler

import (
	"context"

	"github.com/homma509/nrece/domain/repository"
	"github.com/homma509/nrece/usecase"
)

// FileHandler ...
type FileHandler struct {
	u *usecase.Receipt
}

// NewFileHandler ...
func NewFileHandler(repo repository.ReceiptRepository) (*FileHandler, error) {
	return &FileHandler{
		usecase.NewReceipt(repo),
	}, nil
}

// MoveFile ...
func (h *FileHandler) MoveFile(ctx context.Context, s3url string) error {
	if err := h.u.Move(ctx, s3url); err != nil {
		return err
	}
	return nil
}
