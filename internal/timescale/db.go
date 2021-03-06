// Code generated by sqlc. DO NOT EDIT.

package timescale

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
	if q.averageStmt, err = db.PrepareContext(ctx, average); err != nil {
		return nil, fmt.Errorf("error preparing query Average: %w", err)
	}
	if q.countCpuUtilizationStmt, err = db.PrepareContext(ctx, countCpuUtilization); err != nil {
		return nil, fmt.Errorf("error preparing query CountCpuUtilization: %w", err)
	}
	if q.getCpuUtilizationStmt, err = db.PrepareContext(ctx, getCpuUtilization); err != nil {
		return nil, fmt.Errorf("error preparing query GetCpuUtilization: %w", err)
	}
	if q.insertCpuUtilizationStmt, err = db.PrepareContext(ctx, insertCpuUtilization); err != nil {
		return nil, fmt.Errorf("error preparing query InsertCpuUtilization: %w", err)
	}
	if q.upsertCpuUtilizationStmt, err = db.PrepareContext(ctx, upsertCpuUtilization); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertCpuUtilization: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.averageStmt != nil {
		if cerr := q.averageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing averageStmt: %w", cerr)
		}
	}
	if q.countCpuUtilizationStmt != nil {
		if cerr := q.countCpuUtilizationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countCpuUtilizationStmt: %w", cerr)
		}
	}
	if q.getCpuUtilizationStmt != nil {
		if cerr := q.getCpuUtilizationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCpuUtilizationStmt: %w", cerr)
		}
	}
	if q.insertCpuUtilizationStmt != nil {
		if cerr := q.insertCpuUtilizationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertCpuUtilizationStmt: %w", cerr)
		}
	}
	if q.upsertCpuUtilizationStmt != nil {
		if cerr := q.upsertCpuUtilizationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertCpuUtilizationStmt: %w", cerr)
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
	db                       DBTX
	tx                       *sql.Tx
	averageStmt              *sql.Stmt
	countCpuUtilizationStmt  *sql.Stmt
	getCpuUtilizationStmt    *sql.Stmt
	insertCpuUtilizationStmt *sql.Stmt
	upsertCpuUtilizationStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		averageStmt:              q.averageStmt,
		countCpuUtilizationStmt:  q.countCpuUtilizationStmt,
		getCpuUtilizationStmt:    q.getCpuUtilizationStmt,
		insertCpuUtilizationStmt: q.insertCpuUtilizationStmt,
		upsertCpuUtilizationStmt: q.upsertCpuUtilizationStmt,
	}
}
