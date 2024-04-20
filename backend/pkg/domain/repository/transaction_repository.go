package repository

import (
	"context"

	"github.com/yudai2929/kendai-navi/backend/pkg/domain/tx"
)

type TxFn func(ctx context.Context, tx tx.Tx) error

type TransactionRepository interface {
	Tx(ctx context.Context, fn TxFn) error
}
