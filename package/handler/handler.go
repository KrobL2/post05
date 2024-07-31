package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func (h *Handler) InitRoutes() *mux.Router {
	rMux := mux.NewRouter() // Create a new ServeMux using Gorilla

	rMux.HandleFunc("/time", TimeHandler).Methods("GET")

	// Register GET
	getMux := rMux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/getall", GetAllHandler).Methods("GET")
	getMux.HandleFunc("/getid/{username}", GetIDHandler).Methods("GET")
	getMux.HandleFunc("/logged", LoggedUsersHandler).Methods("GET")
	getMux.HandleFunc("/username/{id:[0-9]+}", GetUserDataHandler).Methods("GET")

	rMux.HandleFunc("/username/{id}", GetUserDataHandler).Methods("GET")

	// Register POST
	postMux := rMux.Methods(http.MethodPost).Subrouter()
	// Add User + Login + Logout
	postMux.HandleFunc("/login", LoginHandler).Methods("POST")
	postMux.HandleFunc("/logout", LogoutHandler).Methods("POST")
	postMux.HandleFunc("/add", AddHandler).Methods("POST")

	// Register PUT
	putMux := rMux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", UpdateHandler).Methods("PUT")

	// DELETE
	deleteMux := rMux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username/{id:[0-9]+}", DeleteHandler).Methods("DELETE")

	// COMMON
	rMux.NotFoundHandler = http.HandlerFunc(DefaultHandler)
	rMux.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	return rMux
}

/*

/*



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
