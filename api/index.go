package api

import (
	"fmt"
	"net/http"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the main handler!")
}