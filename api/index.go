package api

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	switch {
	case strings.HasPrefix(path, "users"):
		UserHandler(w, r)
	default:
		fmt.Fprintf(w, "Welcome to the main handler!")
	}
}
