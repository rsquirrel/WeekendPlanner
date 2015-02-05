package api

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
