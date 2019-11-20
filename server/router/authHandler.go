package router

import (
	"api/server/authentication"
	"api/server/config"
	"log"
	"net/http"
)

// AuthHandler handles all of requests as "/auth"
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	serverConf := config.Load()

	username := r.FormValue("username")
	password := r.FormValue("password")

	w.Header().Set("Content-Type", "application/json")

	auth := authentication.NewAuthenticator()
	auth.SetHost(serverConf.DBInfo.Domain, serverConf.DBInfo.Port).
		SetAuthentication(serverConf.DBInfo.User, serverConf.DBInfo.Password).
		SetDatabase(serverConf.DBInfo.Name)

	token, err := auth.GenerateToken(r.RequestURI, username, password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(string(token)))
}
