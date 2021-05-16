// Code generated by sqlc. DO NOT EDIT.

package negativations

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.deleteNegativatedStmt, err = db.PrepareContext(ctx, deleteNegativated); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteNegativated: %w", err)
	}
	if q.getNegativatedByIDStmt, err = db.PrepareContext(ctx, getNegativatedByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetNegativatedByID: %w", err)
	}
	if q.listNegativatedStmt, err = db.PrepareContext(ctx, listNegativated); err != nil {
		return nil, fmt.Errorf("error preparing query ListNegativated: %w", err)
	}
	if q.negativateStmt, err = db.PrepareContext(ctx, negativate); err != nil {
		return nil, fmt.Errorf("error preparing query Negativate: %w", err)
	}
	if q.updateNegativatedStmt, err = db.PrepareContext(ctx, updateNegativated); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateNegativated: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteNegativatedStmt != nil {
		if cerr := q.deleteNegativatedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteNegativatedStmt: %w", cerr)
		}
	}
	if q.getNegativatedByIDStmt != nil {
		if cerr := q.getNegativatedByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNegativatedByIDStmt: %w", cerr)
		}
	}
	if q.listNegativatedStmt != nil {
		if cerr := q.listNegativatedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listNegativatedStmt: %w", cerr)
		}
	}
	if q.negativateStmt != nil {
		if cerr := q.negativateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing negativateStmt: %w", cerr)
		}
	}
	if q.updateNegativatedStmt != nil {
		if cerr := q.updateNegativatedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNegativatedStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	deleteNegativatedStmt  *sql.Stmt
	getNegativatedByIDStmt *sql.Stmt
	listNegativatedStmt    *sql.Stmt
	negativateStmt         *sql.Stmt
	updateNegativatedStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		deleteNegativatedStmt:  q.deleteNegativatedStmt,
		getNegativatedByIDStmt: q.getNegativatedByIDStmt,
		listNegativatedStmt:    q.listNegativatedStmt,
		negativateStmt:         q.negativateStmt,
		updateNegativatedStmt:  q.updateNegativatedStmt,
	}
}
