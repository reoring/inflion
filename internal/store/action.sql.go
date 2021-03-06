// Code generated by sqlc. DO NOT EDIT.
// source: action.sql

package store

import (
	"context"
)

const getActions = `-- name: GetActions :many
SELECT id, user_id, project_id, name, body, created_at, updated_at
FROM action
WHERE project_id = $1
`

func (q *Queries) GetActions(ctx context.Context, projectID int64) ([]Action, error) {
	rows, err := q.query(ctx, q.getActionsStmt, getActions, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProjectID,
			&i.Name,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
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
