package repository

import (
	"context"
	"io"
)

// FileRepository ファイルリポジトリのインターフェース
type FileRepository interface {
	Get(ctx context.Context, name string) (io.ReadCloser, error)
	Copy(ctx context.Context, dst, src string) error
	Publish(ctx context.Context, name string) error
}
