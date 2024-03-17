package main

import (
	"architectures/clean_architecture/after/infrastructure"
	"architectures/clean_architecture/after/interface_adapter"
	"architectures/clean_architecture/after/usecase"
	"net/http"
)

func main() {
	repo := infrastructure.NewInMemoryUserRepository()
	service := usecase.NewUserService(repo)
	controller := interface_adapter.NewUserController(service)

	http.HandleFunc("/users", controller.UserHandler)
	http.ListenAndServe(":8080", nil)
}
