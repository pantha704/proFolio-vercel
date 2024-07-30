package api

import (
	"fmt"
	"net/http"
	"profolio-vercel/handlers"
	"strings"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	switch {
	case strings.HasPrefix(path, "users"):
		handlers.UserHandler(w, r) // Delegate to UserHandler
	default:
		fmt.Fprintf(w, "Welcome to the main handler!")
	}
}
