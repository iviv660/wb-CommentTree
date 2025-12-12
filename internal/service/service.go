package service

import (
	"context"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
)

type GetCommentsInput struct {
	ParentID *int64
	Page     int
	Limit    int
	Sort     string
	Query    string
}

type CommentService interface {
	Create(ctx context.Context, comment model.Comment) (model.Comment, error)
	Get(ctx context.Context, in GetCommentsInput) ([]model.Comment, error)
	Delete(ctx context.Context, id int64) error
}
