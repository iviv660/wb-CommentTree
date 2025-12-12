package comment

import (
	"context"
	"strings"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
)

func (s *Service) Create(ctx context.Context, in model.Comment) (model.Comment, error) {
	if strings.TrimSpace(in.Body) == "" {
		return model.Comment{}, model.ErrEmptyBody
	}

	comment, err := s.repo.Set(ctx, model.Comment{
		ParentID: in.ParentID,
		Body:     in.Body,
	})
	if err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}
