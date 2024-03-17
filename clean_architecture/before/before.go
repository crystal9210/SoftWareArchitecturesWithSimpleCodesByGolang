package main

import (
	"fmt"
	"net/http"
)

// この設計ではビジネスロジック（ユーザーを取得する）が直接HTTPハンドラに結びつけられ、データソース（ここでは静的な配列）へのアクセスも直接組み込まれている

// User model
type User struct {
	ID   int
	Name string
}

// Database simulation
var users = []User{
	{ID: 1, Name: "John Doe"},
}

// Fetch all users
func GetAllUsers() []User {
	return users
}

// HTTP handler for fetching users
func UserHandler(w http.ResponseWriter, r *http.Request) {
	users := GetAllUsers()
	for _, user := range users {
		fmt.Fprintf(w, "ID: %d, Name: %s\n", user.ID, user.Name)
	}
}

func main() {
	http.HandleFunc("/users", UserHandler)
	http.ListenAndServe(":8080", nil)
}
