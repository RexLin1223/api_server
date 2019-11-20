package router

import (
	"api/server/authentication"
	"api/server/config"
	"log"
	"net/http"
)

// MemberHandler is entry of request with URL as "/member"
func MemberHandler(w http.ResponseWriter, r *http.Request) {
	serverConf := config.Load()

	auth := authentication.NewAuthenticator()
	auth.SetHost(serverConf.DBInfo.Domain, serverConf.DBInfo.Port).
		SetAuthentication(serverConf.DBInfo.User, serverConf.DBInfo.Password).
		SetDatabase(serverConf.DBInfo.Name)

	err := auth.VerifyToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello this is member page!\n"))
	}
}
