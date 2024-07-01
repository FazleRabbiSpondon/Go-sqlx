package handlers

import (
	"Go0sqlx/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	newUser := db.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		fmt.Println("Error found data: ", newUser)
		return
	}
	fmt.Println(newUser)
	if err := db.Insert(newUser); err != nil {

		fmt.Println(err.Error())
		return
	}

}


func GetUser(w http.ResponseWriter, r *http.Request) {

	ex_user, err := db.Select()
	if err != nil {
	json.NewEncoder(w).Encode(&ex_user)
	}
	fmt.Printf("%T", ex_user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	newUser := db.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		fmt.Println("Error found data: ", newUser)
		return
	}
	fmt.Println(newUser)
	if err := db.Insert(newUser); err != nil {

		fmt.Println(err.Error())
		return
	}

}