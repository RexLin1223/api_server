package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler is entry of request with URL as "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if len(vars) > 0 {
		w.Write([]byte("hello " + vars["name"] + "!"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello!\n"))
	}
}
