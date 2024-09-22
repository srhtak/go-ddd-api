package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/srhtak/go-ddd-api/internal/application"
)

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    log.Println("Received CreateUser request")
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        log.Printf("Error decoding request body: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    log.Printf("Attempting to create user: %s", input.Username)

    err := h.userService.CreateUser(input.Username, input.Password)
    if err != nil {
        log.Printf("Error creating user: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("User created successfully")
    w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if h.userService.AuthenticateUser(input.Username, input.Password) {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}