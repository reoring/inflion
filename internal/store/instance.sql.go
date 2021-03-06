// Code generated by sqlc. DO NOT EDIT.
// source: instance.sql

package store

import (
	"context"
)

const resolveIdByInstanceId = `-- name: ResolveIdByInstanceId :one
SELECT id
FROM instance
WHERE
  instance_id = $1
LIMIT 1
`

func (q *Queries) ResolveIdByInstanceId(ctx context.Context, instanceID string) (int64, error) {
	row := q.queryRow(ctx, q.resolveIdByInstanceIdStmt, resolveIdByInstanceId, instanceID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectInstance = `-- name: SelectInstance :many
SELECT id, name, instance_id, status, created_at, updated_at, project_id
FROM instance
`

func (q *Queries) SelectInstance(ctx context.Context) ([]Instance, error) {
	rows, err := q.query(ctx, q.selectInstanceStmt, selectInstance)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Instance
	for rows.Next() {
		var i Instance
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.InstanceID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProjectID,
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

const upsertInstance = `-- name: UpsertInstance :one
INSERT INTO instance(instance_id, name, project_id, status)
VALUES ($1, $2, $3, $4)
ON CONFLICT ON CONSTRAINT instance_instance_id_key
    DO UPDATE
    SET name = EXCLUDED.name, status = EXCLUDED.status
RETURNING id, name, instance_id, status, created_at, updated_at, project_id
`

type UpsertInstanceParams struct {
	InstanceID string
	Name       string
	ProjectID  int64
	Status     string
}

func (q *Queries) UpsertInstance(ctx context.Context, arg UpsertInstanceParams) (Instance, error) {
	row := q.queryRow(ctx, q.upsertInstanceStmt, upsertInstance,
		arg.InstanceID,
		arg.Name,
		arg.ProjectID,
		arg.Status,
	)
	var i Instance
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.InstanceID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
	)
	return i, err
}
