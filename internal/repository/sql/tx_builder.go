package sqlrepo

import (
	"context"
	"database/sql"
)

type tx interface {
	begin(func(*sql.Tx) error) error

	readOnly() tx
	lvlReadUncomm() tx
	lvlReadComm() tx
	lvlWriteComm() tx
	lvlRepeatableRead() tx
	lvlSnapshot() tx
	lvlSerializable() tx
	lvlLinearizable() tx
}

type txStruct struct {
	*sql.Tx
	ctx     context.Context
	options *sql.TxOptions

	s SQLStorage
}

func (s *sqlRepository) defaultTx(ctx context.Context) tx {
	return &txStruct{
		options: &sql.TxOptions{
			Isolation: sql.LevelDefault,
			ReadOnly:  false,
		},
		s:   s.storage,
		ctx: ctx,
	}
}

func (t *txStruct) begin(f func(*sql.Tx) error) (err error) {
	t.Tx, err = t.s.SQLDB().BeginTx(t.ctx, t.options)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			_ = t.Commit()
		} else {
			_ = t.Rollback()
		}
	}()

	err = f(t.Tx)
	return err
}

func (tx *txStruct) readOnly() tx {
	tx.options.ReadOnly = true
	return tx
}
func (tx *txStruct) lvlReadUncomm() tx {
	tx.options.Isolation = sql.LevelReadUncommitted
	return tx
}
func (tx *txStruct) lvlReadComm() tx {
	tx.options.Isolation = sql.LevelReadCommitted
	return tx
}
func (tx *txStruct) lvlWriteComm() tx {
	tx.options.Isolation = sql.LevelWriteCommitted
	return tx
}
func (tx *txStruct) lvlRepeatableRead() tx {
	tx.options.Isolation = sql.LevelRepeatableRead
	return tx
}
func (tx *txStruct) lvlSnapshot() tx {
	tx.options.Isolation = sql.LevelSnapshot
	return tx
}
func (tx *txStruct) lvlSerializable() tx {
	tx.options.Isolation = sql.LevelSerializable
	return tx
}
func (tx *txStruct) lvlLinearizable() tx {
	tx.options.Isolation = sql.LevelLinearizable
	return tx
}
