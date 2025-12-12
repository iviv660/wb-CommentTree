package comment

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
)

type createCommentRequest struct {
	ParentID *int64 `json:"parent_id"`
	Body     string `json:"body"`
}

func (a *API) createComment(w http.ResponseWriter, r *http.Request) {
	var req createCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Body) == "" {
		http.Error(w, "comment text is required", http.StatusBadRequest)
		return
	}

	in := model.Comment{
		ParentID: req.ParentID,
		Body:     req.Body,
	}

	out, err := a.service.Create(r.Context(), in)
	if err != nil {
		http.Error(w, "failed to create comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(out)
}
