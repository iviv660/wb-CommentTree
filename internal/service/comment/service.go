package comment

import "github.com/iviv660/wb-CommentTree.git/internal/repository"

type Service struct {
	repo repository.CommentRepository
}

func New(repo repository.CommentRepository) *Service {
	return &Service{repo: repo}
}
