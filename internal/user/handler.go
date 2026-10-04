package user

import (
	"go-rest-api-server/internal/handlers"
	"net/http"
)

var _ handlers.Handler = (*handler)(nil)

const (
	usersURL = "/users"
	userURL  = "/users/{uuid}"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *http.ServeMux) {
	router.HandleFunc("GET "+usersURL, h.GetList)
	router.HandleFunc("POST "+usersURL, h.CreateUser)
	router.HandleFunc("GET "+userURL, h.GetUserByUUID)
	router.HandleFunc("PUT "+userURL, h.UpdateUser)
	router.HandleFunc("PATCH "+userURL, h.PartiallyUpdateUser)
	router.HandleFunc("DELETE "+userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Here will be a list of users"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Here will be the CreateUser"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Here will be the GetUserByUUID"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Here will be the UpdateUser"))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Here will be the PartiallyUpdateUser"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Here will be the DeleteUser"))
}
