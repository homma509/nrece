package repository

import (
	"context"
	"io"
)

// ReceiptRepository レセプトリポジトリのインターフェース
type ReceiptRepository interface {
	Get(ctx context.Context, name string) (io.ReadCloser, error)
	Copy(ctx context.Context, dst, src string) error
	Publish(ctx context.Context, name string) error
	// Store(context.Context, model.Receipt) error
}
