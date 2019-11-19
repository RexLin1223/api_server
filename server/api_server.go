package server

import (
	"encoding/json"
	"log"
	"net/http"

	auth "api/server/authentication"

	"github.com/gorilla/mux"
)

type tokenResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) > 0 {
		w.Write([]byte("hello " + vars["name"] + "!"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello!\n"))
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	token, err := auth.GetToken(username, password)
	if err != nil {
		log.Fatal(err)
	}

	resp := tokenResponse{
		Result:  "Successful",
		Message: "Token",
		Token:   token,
	}
	payload, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Server", "golang server")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(payload)))
}

func memberHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello this is member page!\n"))
}
func assistantHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello this is assistant page!\n"))
}
func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello this is admin page!\n"))
}

// Start server
func Start() {
	// Load api_meta.yaml

	r := mux.NewRouter()

	// Register route and handler
	r.HandleFunc("/", homepageHandler)
	r.HandleFunc("/admin", adminHandler)
	r.HandleFunc("/member", adminHandler)
	r.HandleFunc("/assistant", assistantHandler)
	r.HandleFunc("/auth", authHandler)

	port := "8857"
	log.Println("Start listening port:" + port)
	// Blocking here
	log.Fatal(http.ListenAndServe(":"+port, r))
}
