package api

import (
	"encoding/json"
	"net/http"

	"distributed-kv-store/internal/store"
)

type Handler struct {
	store *store.Store
}

func NewHandler(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {

	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)

	key := data["key"]
	value := data["value"]

	h.store.Set(key, value)

	w.Write([]byte("OK"))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")

	value, ok := h.store.Get(key)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte(value))
}