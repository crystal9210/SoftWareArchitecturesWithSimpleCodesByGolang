package interface_adapter

import (
	"architectures/clean_architecture/after/usecase"
	"fmt"
	"net/http"
)

// 構造体の定義
type UserController struct {
	service *usecase.UserService
}

func NewUserController(service *usecase.UserService) *UserController {
	return &UserController{service: service}
}

func (controller *UserController) UserHandler(w http.ResponseWriter, r *http.Request) {
	users := controller.service.GetAllUsers()
	for _, user := range users {
		fmt.Fprintf(w, "ID: %d, Name: %s\n", user.ID, user.Name)
	}
}
