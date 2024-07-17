package handler

import (
	"fmt"
	"log"
	"net/http"
)

// MethodNotAllowedHandler is executed when the HTTP method is incorrect
func MethodNotAllowedHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	rw.WriteHeader(http.StatusNotFound)

	Body := "Method not allowed!\n"
	fmt.Fprintf(rw, "%s", Body)
}
