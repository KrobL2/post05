package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func (h *Handler) InitRoutes() *mux.Router {
	rMux := mux.NewRouter() // Create a new ServeMux using Gorilla
	// Define Handler Functions
	rMux.HandleFunc("/time", TimeHandler).Methods("GET")

	// Register GET
	// getMux := rMux.Methods(http.MethodGet).Subrouter()
	rMux.HandleFunc("/getall", GetAllHandler).Methods("GET")
	rMux.HandleFunc("/getid/{username}", GetIDHandler).Methods("GET")
	rMux.HandleFunc("/logged", LoggedUsersHandler).Methods("GET")
	rMux.HandleFunc("/username/{id:[0-9]+}", GetUserDataHandler).Methods("GET")

	// Register POST
	// Add User + Login + Logout
	// postMux := rMux.Methods(http.MethodPost).Subrouter()
	rMux.HandleFunc("/login", LoginHandler).Methods("POST")
	rMux.HandleFunc("/logout", LogoutHandler).Methods("POST")
	rMux.HandleFunc("/add", AddHandler).Methods("POST")

	rMux.HandleFunc("/username/{id}", GetUserDataHandler).Methods("GET")

	// Register PUT
	// Update User
	// putMux := rMux.Methods(http.MethodPut).Subrouter()
	rMux.HandleFunc("/update", UpdateHandler).Methods("PUT")

	// DELETE
	// deleteMux := rMux.Methods(http.MethodDelete).Subrouter()
	rMux.HandleFunc("/username/{id:[0-9]+}", DeleteHandler).Methods("DELETE")

	// COMMON
	rMux.NotFoundHandler = http.HandlerFunc(DefaultHandler)
	// notAllowed := notAllowedHandler{}
	// rMux.MethodNotAllowedHandler = notAllowed
	rMux.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	return rMux
}

/*

/* package main



type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	MethodNotAllowedHandler(rw, r)
}
*/
