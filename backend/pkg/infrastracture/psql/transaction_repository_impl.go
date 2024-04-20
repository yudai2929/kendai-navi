package psql

import (
	"context"

	"github.com/yudai2929/kendai-navi/backend/pkg/domain/repository"
	db "github.com/yudai2929/kendai-navi/backend/pkg/lib/ent"
	"github.com/yudai2929/kendai-navi/backend/pkg/lib/errors"
)

type TransactionRepositoryImpl struct {
	client *db.Client
}

func NewTransactionRepositoryImpl(client *db.Client) repository.TransactionRepository {
	return &TransactionRepositoryImpl{client: client}
}

func (r *TransactionRepositoryImpl) Tx(ctx context.Context, txFn repository.TxFn) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			err := tx.Rollback()
			if err == nil {
				return
			}
			// TODO: log
		}
	}()

	if err := txFn(ctx, tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrap(err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err)
	}
	return nil
}
