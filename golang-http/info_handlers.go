package main

import (
	"fmt"
	"net/http"
)

func UndefinedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Invoked", "true")
	http.NotFound(w, r)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("X-Invoked", "true")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
