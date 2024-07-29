package api

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello from the main handler!")
}
