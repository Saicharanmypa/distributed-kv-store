package main

import (
	"fmt"
	"net/http"

	"distributed-kv-store/internal/api"
	"distributed-kv-store/internal/store"
)

func main() {

	store := store.NewStore()
	handler := api.NewHandler(store)

	http.HandleFunc("/set", handler.Set)
	http.HandleFunc("/get", handler.Get)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}
