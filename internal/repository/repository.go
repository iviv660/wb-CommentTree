package repository

import (
	"context"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
	"github.com/iviv660/wb-CommentTree.git/internal/service"
)

type CommentRepository interface {
	Get(ctx context.Context, in service.GetCommentsInput) ([]model.Comment, error)
	Set(ctx context.Context, comment model.Comment) (model.Comment, error)
	Delete(ctx context.Context, id int64) error
}
