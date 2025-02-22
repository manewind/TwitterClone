package controller

import (
	"auth-services/internal/models"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type UserController struct {
	userService models.UserInterface
}

func NewUserController(userService models.UserInterface) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) MainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		uc.CreateUserController(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (uc *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {

	user, err := decodeResponseUser(r)
	if err != nil {
		uc.userService.GetLogger().Error("Error decoding response", zap.Error(err))
		http.Error(w, "Error decode response", http.StatusBadRequest)
	}

	uc.userService.GetLogger().Info("Creating new user", zap.String("user", user.Username))

	err = uc.userService.CreateUser(r.Context(), user)
	if err != nil {
		uc.userService.GetLogger().Error("Error creating user", zap.Error(err))
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := encodeResponse(w, user); err != nil {
		uc.userService.GetLogger().Error("Error encoding response", zap.Error(err))
	}
}

func decodeResponseUser(r *http.Request) (models.User, error) {
	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return models.User{}, err
	}
	return req, nil

}

func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
