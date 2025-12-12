package comment

import (
	"context"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
)

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM comments WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return model.ErrNotFound
	}
	return nil
}
