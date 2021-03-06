// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package timescale

import (
	"context"
	"database/sql"
	"time"
)

const average = `-- name: Average :one
SELECT time_bucket('10 minutes', time) AS period,
       instance_id,
       avg(value)
FROM metrics
WHERE instance_id = $1
GROUP BY period, instance_id
LIMIT 1
`

type AverageRow struct {
	Period     interface{}
	InstanceID string
	Avg        interface{}
}

func (q *Queries) Average(ctx context.Context, instanceID string) (AverageRow, error) {
	row := q.queryRow(ctx, q.averageStmt, average, instanceID)
	var i AverageRow
	err := row.Scan(&i.Period, &i.InstanceID, &i.Avg)
	return i, err
}

const countCpuUtilization = `-- name: CountCpuUtilization :one
SELECT count(time)
FROM metrics
WHERE time = $1
  AND type = $2
  AND instance_id = $3
`

type CountCpuUtilizationParams struct {
	Time       time.Time
	Type       string
	InstanceID string
}

func (q *Queries) CountCpuUtilization(ctx context.Context, arg CountCpuUtilizationParams) (int64, error) {
	row := q.queryRow(ctx, q.countCpuUtilizationStmt, countCpuUtilization, arg.Time, arg.Type, arg.InstanceID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getCpuUtilization = `-- name: GetCpuUtilization :many
SELECT time, instance_id, type, value
FROM metrics
ORDER BY time
LIMIT 100
`

func (q *Queries) GetCpuUtilization(ctx context.Context) ([]Metric, error) {
	rows, err := q.query(ctx, q.getCpuUtilizationStmt, getCpuUtilization)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Metric
	for rows.Next() {
		var i Metric
		if err := rows.Scan(
			&i.Time,
			&i.InstanceID,
			&i.Type,
			&i.Value,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCpuUtilization = `-- name: InsertCpuUtilization :one
INSERT INTO metrics(time, instance_id, type, value)
VALUES ($1, $2, $3, $4)
RETURNING time
`

type InsertCpuUtilizationParams struct {
	Time       time.Time
	InstanceID string
	Type       string
	Value      sql.NullFloat64
}

func (q *Queries) InsertCpuUtilization(ctx context.Context, arg InsertCpuUtilizationParams) (time.Time, error) {
	row := q.queryRow(ctx, q.insertCpuUtilizationStmt, insertCpuUtilization,
		arg.Time,
		arg.InstanceID,
		arg.Type,
		arg.Value,
	)
	var time time.Time
	err := row.Scan(&time)
	return time, err
}

const upsertCpuUtilization = `-- name: UpsertCpuUtilization :one
INSERT INTO metrics(time, instance_id, type, value)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING
RETURNING time
`

type UpsertCpuUtilizationParams struct {
	Time       time.Time
	InstanceID string
	Type       string
	Value      sql.NullFloat64
}

func (q *Queries) UpsertCpuUtilization(ctx context.Context, arg UpsertCpuUtilizationParams) (time.Time, error) {
	row := q.queryRow(ctx, q.upsertCpuUtilizationStmt, upsertCpuUtilization,
		arg.Time,
		arg.InstanceID,
		arg.Type,
		arg.Value,
	)
	var time time.Time
	err := row.Scan(&time)
	return time, err
}
