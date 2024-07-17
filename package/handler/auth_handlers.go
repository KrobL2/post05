package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/mactsouk/restdb"
)

// LoginHandler is for updating the LastLogin time of a user
// And changing the Active field to true
func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler Serving:", r.URL.Path, "from", r.Host)

	d, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = restdb.User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("Input user:", user)

	if !restdb.IsUserValid(user) {
		log.Println("User", user.Username, "not valid!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := restdb.FindUserUsername(user.Username)
	log.Println("Logging in:", t)

	t.LastLogin = time.Now().Unix()
	t.Active = 1
	if restdb.UpdateUser(t) {
		log.Println("User updated:", t)
		rw.WriteHeader(http.StatusOK)
	} else {
		log.Println("Update failed:", t)
		rw.WriteHeader(http.StatusBadRequest)
	}
}

// LogoutHandler is for logging out a user
// And changing the Active field to false
func LogoutHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("LogoutHandler Serving:", r.URL.Path, "from", r.Host)

	d, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = restdb.User{}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !restdb.IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	t := restdb.FindUserUsername(user.Username)
	log.Println("Logging out:", t.Username)
	t.Active = 0
	if restdb.UpdateUser(t) {
		log.Println("User updated:", t)
		rw.WriteHeader(http.StatusOK)
	} else {
		log.Println("Update failed:", t)
		rw.WriteHeader(http.StatusBadRequest)
	}
}

// LoggedUsersHandler returns the list of all logged in users
func LoggedUsersHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("LoggedUsersHandler Serving:", r.URL.Path, "from", r.Host)

	var user = restdb.User{}
	err := user.FromJSON(r.Body)

	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !restdb.IsUserValid(user) {
		log.Println("User", user.Username, "exists!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SliceToJSON(restdb.ReturnLoggedUsers(), rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
