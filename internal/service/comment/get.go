package comment

import (
	"context"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
	"github.com/iviv660/wb-CommentTree.git/internal/service"
)

func (s *Service) Get(ctx context.Context, in service.GetCommentsInput) ([]model.Comment, error) {
	comments, err := s.repo.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, model.ErrNotFoundComment
	}

	return comments, nil
}
