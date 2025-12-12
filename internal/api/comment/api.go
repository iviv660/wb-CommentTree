// internal/api/comment/api.go
package comment

import (
	"github.com/go-chi/chi/v5"
	"github.com/iviv660/wb-CommentTree.git/internal/service"
)

type API struct {
	mux     *chi.Mux
	service service.CommentService
}

func NewAPI(mux *chi.Mux, service service.CommentService) *API {
	return &API{
		mux:     mux,
		service: service,
	}
}

func (a *API) RegisterHandler() {
	a.mux.Get("/", a.index)
	a.mux.Get("/comments", a.getComment)
	a.mux.Post("/comments", a.createComment)
	a.mux.Delete("/comments/{id}", a.deleteComment)
}
