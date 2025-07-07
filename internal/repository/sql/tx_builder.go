package sqlrepo

import (
	"context"
	"database/sql"
)

//
//const (
//	LevelDefault IsolationLevel = iota
//	LevelReadUncommitted
//	LevelReadCommitted
//	LevelWriteCommitted
//	LevelRepeatableRead
//	LevelSnapshot
//	LevelSerializable
//	LevelLinearizable
//)

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

func (s SQLStorage) defaultTx(ctx context.Context) tx {
	return &txStruct{
		options: &sql.TxOptions{
			Isolation: sql.LevelDefault,
			ReadOnly:  false,
		},
		s:   s,
		ctx: ctx,
	}
}

func (t *txStruct) begin(f func(*sql.Tx) error) (err error) {
	t.Tx, err = t.s.BeginTx(t.ctx, t.options)
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

/*
func (s SQLStorage) withTx(ctx context.Context, f func(*sql.Tx) error, ro bool, iso ...sql.IsolationLevel) error {
	tx, err := s.BeginTx(ctx, &sql.TxOptions{
		Isolation: iso,
		ReadOnly:  ro,
	})

	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = f(tx)
	return err
}
*/
