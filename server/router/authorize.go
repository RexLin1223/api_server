package router

import (
	"api/server/verification"
	"log"
	"net/http"
)

// AuthHandler handles all of requests as "/auth"
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	auth := verification.NewAuthenticator()
	token, err := auth.GenerateToken(r.RequestURI, username, password)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 401)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(token)))
}
