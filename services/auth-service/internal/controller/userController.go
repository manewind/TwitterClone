package controller

import (
	"auth-services/internal/models"
	"encoding/json"
	"net/http"
)

type UserController struct {
	userService models.UserInterface
}

func NewUserController(userService models.UserInterface) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	err := uc.userService.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
