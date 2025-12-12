package comment

import (
	"context"
	"fmt"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
	"github.com/iviv660/wb-CommentTree.git/internal/service"
)

func (r *Repository) Get(ctx context.Context, in service.GetCommentsInput) ([]model.Comment, error) {
	const baseQuery = `
SELECT id, parent_id, body, created_at
FROM comments
WHERE 1=1
`

	query := baseQuery
	args := make([]any, 0, 4)
	argPos := 1

	if in.ParentID != nil {
		query += fmt.Sprintf(" AND (id = $%d OR parent_id = $%d)", argPos, argPos)
		args = append(args, *in.ParentID)
		argPos++
	}

	if in.Query != "" {
		query += fmt.Sprintf(" AND search_vector @@ plainto_tsquery('russian', $%d)", argPos)
		args = append(args, in.Query)
		argPos++
	}

	orderBy := "created_at DESC"
	switch in.Sort {
	case "created_at_asc":
		orderBy = "created_at ASC"
	case "created_at_desc":
		orderBy = "created_at DESC"

	}

	query += " ORDER BY " + orderBy

	limit := in.Limit
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	page := in.Page
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.Comment, 0, limit)
	for rows.Next() {
		var c model.Comment
		if err := rows.Scan(&c.ID, &c.ParentID, &c.Body, &c.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
