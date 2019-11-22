package helper

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// GetParam will find value by {key} from request vars
func GetParam(r *http.Request, key string) (string, error) {
	vars := mux.Vars(r)
	value, ok := vars[key]
	if !ok {
		return "", errors.New("Can't find param: " + key)
	}
	return value, nil
}

// GetBody will fetch data as bytes from request body
func GetBody(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
