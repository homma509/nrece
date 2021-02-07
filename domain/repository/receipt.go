package repository

import (
	"context"
	"io"
)

// ReceiptRepository レセプトリポジトリのインターフェース
type ReceiptRepository interface {
	Get(ctx context.Context, s3url string) (io.ReadCloser, error)
	Copy(ctx context.Context, fromURL string, toURL string) error
	// Store(context.Context, model.Receipt) error
}
