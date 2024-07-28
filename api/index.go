package api

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/")
    
    switch path {
	case "":
		fmt.Fprintf(w, "Hello from the main handler!")
    case "user":
        GetAllUsersHandler(w, r)
    default:
        http.Error(w, "Not Found", http.StatusNotFound)
    }
}
