package server

import (
	"api/server/config"
	"api/server/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start server
func Start() {
	// Load api_meta.yaml
	conf := config.Load()

	r := mux.NewRouter()
	// Register route and handler
	r.HandleFunc(conf.Router.Home, router.HomeHandler)
	r.HandleFunc(conf.Router.Authorize, router.AuthHandler).Methods("GET", "POST")

	// Admin
	r.HandleFunc("/admin", router.AdminHandler).Methods("GET", "POST", "PUT", "DELETE")
	r.HandleFunc("/admin/{category}", router.AdminHandler).Methods("GET", "POST", "PUT", "DELETE")

	// Member
	r.HandleFunc("/member", router.MemberHandler).Methods("GET", "POST", "PUT", "DELETE")
	r.HandleFunc("/member/{category}", router.MemberHandler).Methods("GET", "POST", "PUT", "DELETE")

	// Assistant
	r.HandleFunc("/assistant", router.AssistantHandler).Methods("GET", "POST", "PUT", "DELETE")
	r.HandleFunc("/assistant/{category}", router.AssistantHandler).Methods("GET", "POST", "PUT", "DELETE")

	// Starting listen request
	log.Println("Start listening port:" + conf.Port)
	log.Fatal(http.ListenAndServe(conf.Domain+":"+conf.Port, r))
}
