package comment

import (
	"context"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
)

func (r *Repository) Set(ctx context.Context, in model.Comment) (model.Comment, error) {
	const query = `
INSERT INTO comments (parent_id, body)
VALUES ($1, $2)
RETURNING id, parent_id, body, created_at;
`

	var out model.Comment
	if err := r.pool.QueryRow(ctx, query,
		in.ParentID,
		in.Body,
	).Scan(
		&out.ID,
		&out.ParentID,
		&out.Body,
		&out.CreatedAt,
	); err != nil {
		return model.Comment{}, err
	}

	return out, nil
}
