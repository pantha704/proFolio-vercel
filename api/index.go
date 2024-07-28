package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	
	router := mux.NewRouter()
    RegisterUserRoutes(router)
    router.ServeHTTP(w, r)

	switch path {
	case "user":
		GetAllUsersHandler(w, r)
	// case "cover-letter":
	// 	GetCoverLetterHandler(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
