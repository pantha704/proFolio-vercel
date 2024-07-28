package api

import (
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/")
    
    switch path {
    case "user":
        GetAllUsersHandler(w, r)
    default:
        http.Error(w, "Not Found", http.StatusNotFound)
    }
}
