package comment

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/iviv660/wb-CommentTree.git/internal/model"
	"github.com/iviv660/wb-CommentTree.git/internal/service"
)

func (a *API) getComment(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	var parentID *int64
	if v := q.Get("parent_id"); v != "" {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			http.Error(w, "invalid parent_id", http.StatusBadRequest)
			return
		}
		parentID = &id
	} else if v := q.Get("parent"); v != "" {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			http.Error(w, "invalid parent", http.StatusBadRequest)
			return
		}
		parentID = &id
	}

	page := 1
	if p := q.Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	limit := 20
	if l := q.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 && v <= 100 {
			limit = v
		}
	}

	rawSort := strings.ToLower(q.Get("sort"))
	sort := ""
	switch rawSort {
	case "", "newest":
		sort = "created_at_desc"
	case "oldest":
		sort = "created_at_asc"
	default:
		sort = "created_at_desc"
	}

	search := q.Get("q")

	comments, err := a.service.Get(r.Context(), service.GetCommentsInput{
		ParentID: parentID,
		Page:     page,
		Limit:    limit,
		Sort:     sort,
		Query:    search,
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFoundComment) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode([]model.Comment{})
			return
		}

		http.Error(w, "failed to get comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(comments)
}
