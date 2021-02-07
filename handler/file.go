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

// CopyFile ...
func (h *FileHandler) CopyFile(ctx context.Context, s3url string) error {
	if err := h.u.Copy(ctx, s3url); err != nil {
		return err
	}
	return nil
}
