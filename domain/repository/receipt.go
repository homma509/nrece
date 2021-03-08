package repository

import (
	"context"

	"github.com/homma509/nrece/domain/model"
)

// ReceiptRepository レセプトリポジトリのインターフェース
type ReceiptRepository interface {
	Save(context.Context, model.Receipt) error
}
