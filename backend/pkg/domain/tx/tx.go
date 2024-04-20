package tx

type Tx interface {
	Rollback() error
	Commit() error
}
