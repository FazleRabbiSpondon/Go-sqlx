package main

import (
	"Go0sqlx/db"
	"Go0sqlx/web/handlers"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	db.InitDB()
	http.HandleFunc("/create", handlers.CreateUser)
	http.HandleFunc("/get", handlers.GetUser)


	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
