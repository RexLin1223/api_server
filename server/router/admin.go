package router

import (
	"api/database"
	"api/server/config"
	"api/server/helper"
	"api/server/verification"
	"encoding/json"
	"errors"
	"net/http"
)

// Parse requests
func (h *AdminRequestHandler) Parse(method string, payload string) (string, error) {
	switch method {
	case "GET":
		result, err := h.Query(payload)
		return result, err
	case "POST":
		err := h.Insert(payload)
		return "", err
	case "PUT":
		err := h.Update(payload)
		return "", err
	case "DELETE":
		err := h.Delete(payload)
		return "", err
	default:
		return "", errors.New("Unspported method")
	}
}

// AdminRequestHandler performs requests from admin user
type AdminRequestHandler struct {
	dbInfo   *database.ConnectionInfo
	category string
	index    string // Optional
}

// NewAdminRequestHandler create new AdminRequestHandler
func NewAdminRequestHandler(table string, index string) IRequestHandler {
	serverConfig := config.Load()
	return &AdminRequestHandler{
		dbInfo: &database.ConnectionInfo{
			Domain:     serverConfig.DBInfo.Domain,
			Port:       serverConfig.DBInfo.Port,
			Username:   serverConfig.DBInfo.User,
			Password:   serverConfig.DBInfo.Password,
			TargetName: serverConfig.DBInfo.TargetName,
		},
		category: table,
		index:    index,
	}
}

// Insert into database
func (h *AdminRequestHandler) Insert(inputJSON string) error {
	params := make(map[string]interface{})
	err := json.Unmarshal([]byte(inputJSON), &params)
	if err != nil {
		return err
	}

	conn := database.NewMySQLLConnector(h.dbInfo)
	return conn.Create(h.category, params)
}

// Query result from database
func (h *AdminRequestHandler) Query(inputJSON string) (string, error) {
	conn := database.NewMySQLLConnector(h.dbInfo)
	params, err := conn.Read(h.category, h.index)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Update specific value into database
func (h *AdminRequestHandler) Update(inputJSON string) error {
	params := make(map[string]interface{})
	err := json.Unmarshal([]byte(inputJSON), &params)
	if err != nil {
		return err
	}

	conn := database.NewMySQLLConnector(h.dbInfo)
	return conn.Update(h.category, params, h.index)
}

// Delete data from database
func (h *AdminRequestHandler) Delete(inputJSON string) error {
	conn := database.NewMySQLLConnector(h.dbInfo)
	return conn.Delete(h.category, h.index)
}

// AdminHandler is entry of request with URL as "/admin"
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if _, found := r.Header["Authorization"]; !found {
		http.Error(w, "", 401)
		return
	}

	auth := verification.NewAuthenticator()
	if err := auth.VerifyToken(r.Header.Get("Authorization")); err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	category, err := helper.GetParam(r, "category")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	index, _ := helper.GetParam(r, "index") // Optional

	body, err := helper.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	handler := NewAdminRequestHandler(category, index)
	result, err := handler.Parse(r.Method, body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
